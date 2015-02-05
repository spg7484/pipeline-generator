package pipeline

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func templates_jenkins_multi_job_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x56,
		0x4d, 0x73, 0xdb, 0x36, 0x10, 0xbd, 0xfb, 0x57, 0x70, 0x74, 0xf1, 0x49,
		0x90, 0x9d, 0x36, 0x9d, 0x1c, 0x68, 0xe5, 0x43, 0x89, 0x67, 0xd2, 0x89,
		0x15, 0x8d, 0x14, 0x8d, 0xcf, 0x10, 0xb9, 0x96, 0x60, 0x81, 0x00, 0x07,
		0x00, 0x5b, 0xa9, 0x99, 0xfc, 0xf7, 0x2e, 0x40, 0x00, 0x04, 0x4d, 0xa5,
		0xb5, 0xc6, 0x39, 0x99, 0xd8, 0x2f, 0x2c, 0xde, 0xbe, 0x7d, 0x72, 0xfe,
		0xf6, 0x50, 0xf1, 0xec, 0x2f, 0x50, 0x9a, 0x49, 0x71, 0x73, 0x79, 0x4d,
		0xae, 0x2e, 0x33, 0x10, 0x85, 0x2c, 0x99, 0xd8, 0xde, 0x5c, 0xae, 0xbf,
		0xdd, 0x8e, 0xdf, 0x5c, 0xbe, 0x9d, 0x5e, 0xe4, 0x85, 0xac, 0x88, 0x61,
		0x7b, 0xca, 0xc9, 0x23, 0x88, 0x3d, 0x13, 0x9a, 0xd4, 0xbc, 0xd9, 0xda,
		0xbf, 0x55, 0xc3, 0x0d, 0x7b, 0x94, 0x1b, 0x72, 0x67, 0x3f, 0xfe, 0x94,
		0x9b, 0x85, 0x92, 0x8f, 0x50, 0x98, 0xac, 0x0d, 0xb8, 0x19, 0xf9, 0x84,
		0x71, 0x08, 0x1c, 0xb7, 0x8e, 0x77, 0xd7, 0xe4, 0xfa, 0xf5, 0x68, 0x7a,
		0x91, 0x65, 0x39, 0x2d, 0x0c, 0xde, 0xae, 0x27, 0xee, 0x50, 0x82, 0x2e,
		0x14, 0xab, 0xad, 0x65, 0x9a, 0x4f, 0xd2, 0x93, 0xf5, 0xee, 0x01, 0xea,
		0x8f, 0x50, 0x83, 0x28, 0xb1, 0x4b, 0x06, 0x7a, 0xfa, 0x40, 0xb9, 0x86,
		0x7c, 0x32, 0xb0, 0xdb, 0xe0, 0x5a, 0xc9, 0x1a, 0x94, 0xf1, 0x47, 0x34,
		0x68, 0x20, 0x25, 0xa3, 0x1b, 0x99, 0xbc, 0x82, 0xd5, 0xc0, 0x99, 0x00,
		0xb2, 0xf0, 0x1f, 0x8b, 0x36, 0xe7, 0x18, 0xdb, 0x2f, 0xd1, 0x8c, 0xf8,
		0x1c, 0xc7, 0x21, 0x34, 0xf4, 0x7f, 0x45, 0xde, 0x90, 0x3f, 0x46, 0x6d,
		0x65, 0xac, 0x6d, 0xa8, 0xde, 0xcf, 0x69, 0x05, 0xd3, 0x9a, 0x2a, 0xca,
		0x39, 0xf0, 0x0c, 0x0e, 0x50, 0x34, 0xb6, 0xf3, 0x7c, 0x12, 0x9d, 0x21,
		0x5a, 0x1b, 0xba, 0x05, 0x67, 0xf9, 0xfe, 0x3d, 0x23, 0xab, 0x70, 0xca,
		0x7e, 0xfc, 0xc8, 0x27, 0x9d, 0xaf, 0xed, 0x7a, 0x72, 0x46, 0xdb, 0xee,
		0xdd, 0x93, 0xfe, 0xc3, 0x73, 0x5d, 0x54, 0x59, 0xc1, 0xa9, 0xd6, 0x37,
		0xa3, 0x5d, 0x53, 0x6a, 0x29, 0x08, 0x5a, 0xc8, 0xbc, 0xe1, 0x7c, 0x35,
		0xbb, 0x1b, 0xb5, 0xb0, 0x17, 0x54, 0x2c, 0x25, 0xad, 0xa6, 0x46, 0x35,
		0x08, 0x67, 0x38, 0xb9, 0x81, 0x30, 0x4d, 0x37, 0x1c, 0xca, 0x00, 0x75,
		0x3c, 0x5b, 0xe7, 0x86, 0xcb, 0x62, 0xff, 0xa1, 0x61, 0xbc, 0xbc, 0xdf,
		0x81, 0xf8, 0x28, 0xff, 0x16, 0xda, 0x28, 0xa0, 0x95, 0x33, 0x21, 0x87,
		0x42, 0xd2, 0xff, 0xc6, 0x0d, 0x8b, 0xad, 0xeb, 0xe7, 0x94, 0x1a, 0x44,
		0xd9, 0x42, 0x46, 0xb1, 0xed, 0x16, 0x39, 0xed, 0x9f, 0x26, 0x45, 0xd1,
		0x28, 0x05, 0xc2, 0xb8, 0xa0, 0x50, 0xe7, 0xa9, 0xd9, 0x75, 0x60, 0xbf,
		0x30, 0xd1, 0x23, 0x7f, 0x06, 0xeb, 0x3f, 0xb4, 0x99, 0x71, 0xc0, 0xf5,
		0x8e, 0xea, 0x76, 0x88, 0xee, 0xeb, 0x0a, 0x87, 0x12, 0x2d, 0xbd, 0x18,
		0xcc, 0xd5, 0xc1, 0x92, 0x65, 0xc8, 0x06, 0x45, 0xc5, 0x16, 0x90, 0x13,
		0xcd, 0xc6, 0xba, 0x90, 0x11, 0xd1, 0xf9, 0x9c, 0x7e, 0x16, 0xa1, 0xe6,
		0x4c, 0x8a, 0x07, 0xb6, 0xed, 0x2a, 0x63, 0x3a, 0xfa, 0x23, 0xe5, 0x1c,
		0xd3, 0x82, 0x21, 0x0d, 0xb2, 0x98, 0x2c, 0x90, 0xc2, 0x95, 0x0e, 0x5c,
		0xe8, 0x0c, 0x69, 0x1c, 0x1c, 0x6a, 0xa9, 0xa1, 0x44, 0x06, 0x05, 0x40,
		0x13, 0x4b, 0x1a, 0xe8, 0xf9, 0x82, 0x3d, 0x3d, 0x61, 0x90, 0xb5, 0xa4,
		0x81, 0xb8, 0x38, 0x1a, 0x27, 0xb8, 0x6c, 0x38, 0xe8, 0x05, 0x35, 0x3b,
		0x5c, 0xff, 0x81, 0x29, 0x8d, 0xaf, 0xe8, 0x61, 0x09, 0x38, 0x6a, 0x64,
		0x39, 0xc2, 0x9b, 0x9c, 0x7a, 0x6d, 0x0a, 0x7b, 0x95, 0xf5, 0x1c, 0x57,
		0x46, 0x51, 0x03, 0xdb, 0x63, 0xec, 0xf7, 0x84, 0x6b, 0x98, 0x8a, 0x38,
		0x96, 0xcc, 0xe9, 0x4f, 0x2f, 0xad, 0x33, 0xa7, 0x29, 0xb8, 0xa4, 0xca,
		0xbc, 0xe7, 0xdc, 0x3e, 0xad, 0x45, 0x2f, 0xb5, 0xf4, 0x60, 0x8e, 0xf9,
		0x8e, 0x89, 0x27, 0x6a, 0xed, 0x19, 0xe7, 0x6e, 0x9a, 0x5f, 0x05, 0x26,
		0x2f, 0x41, 0xe3, 0x88, 0xbb, 0x5b, 0x6f, 0xdf, 0x7f, 0xfe, 0xb2, 0x5e,
		0x7e, 0x42, 0xe5, 0xfb, 0xcf, 0xb0, 0xb4, 0xa0, 0x23, 0xf7, 0x57, 0xc1,
		0x8f, 0x9f, 0x1f, 0x70, 0x46, 0xb3, 0x9d, 0xe5, 0x59, 0x14, 0xd0, 0xd3,
		0xce, 0x8e, 0x7a, 0x93, 0x97, 0x70, 0x0f, 0x09, 0x87, 0xca, 0xdc, 0x51,
		0xd9, 0x2f, 0x43, 0x4a, 0x7d, 0x0b, 0x88, 0x61, 0xa2, 0xa1, 0xb6, 0xed,
		0xae, 0xff, 0xd5, 0x7a, 0x36, 0xfb, 0xb4, 0x5a, 0xdd, 0xae, 0xbf, 0x38,
		0x98, 0x4e, 0x44, 0x5c, 0x3c, 0xbb, 0xbd, 0x13, 0xab, 0xea, 0xdf, 0xed,
		0x37, 0x3e, 0xaf, 0x9b, 0x0d, 0x67, 0x7a, 0x17, 0x05, 0x00, 0xfb, 0x66,
		0x0f, 0x19, 0x99, 0xc3, 0xc1, 0xdc, 0x51, 0xbc, 0x99, 0xa7, 0x1b, 0x99,
		0xd3, 0x86, 0xd8, 0x5b, 0x0b, 0x54, 0x10, 0xd5, 0x54, 0xfa, 0xa8, 0x0d,
		0x54, 0x9a, 0x78, 0x85, 0x6d, 0x6f, 0x27, 0xae, 0x7c, 0x54, 0x6b, 0xaf,
		0x4b, 0xc4, 0x75, 0x10, 0xa4, 0xfb, 0x5b, 0x6b, 0x8c, 0x3f, 0x38, 0x2e,
		0x65, 0xf0, 0x6b, 0x73, 0x4d, 0x7e, 0x27, 0xbf, 0x8d, 0x52, 0xb4, 0x10,
		0xe1, 0x56, 0xe0, 0xda, 0x2d, 0x8b, 0x92, 0xea, 0x7f, 0x7f, 0xed, 0x62,
		0x6b, 0xb7, 0xea, 0x83, 0xf6, 0x71, 0x03, 0x4f, 0x47, 0x7b, 0x2c, 0x7f,
		0xed, 0xc3, 0x22, 0x94, 0x09, 0x05, 0x3a, 0x99, 0xb3, 0xcd, 0xf5, 0x50,
		0xed, 0x5d, 0x83, 0x53, 0xb4, 0xca, 0x03, 0x06, 0x14, 0xfb, 0x07, 0xca,
		0xde, 0x35, 0x4f, 0x71, 0xeb, 0x45, 0x8e, 0x7d, 0xe8, 0xbb, 0x57, 0xe4,
		0xd5, 0xeb, 0x01, 0x6c, 0x09, 0xaf, 0xcf, 0xbd, 0xee, 0x84, 0xac, 0x0e,
		0x8a, 0x3e, 0xbb, 0xf0, 0x5c, 0x96, 0xb0, 0x08, 0x8e, 0x6e, 0x98, 0x81,
		0xd1, 0x83, 0xb2, 0xf6, 0x3f, 0x1a, 0x3b, 0x2d, 0xdd, 0x49, 0x78, 0xb4,
		0x9c, 0x16, 0x17, 0xbf, 0x3f, 0x3f, 0xd3, 0x18, 0xdf, 0xc8, 0x3d, 0x33,
		0xbb, 0xb9, 0xec, 0x3a, 0x09, 0xa2, 0xf0, 0x33, 0x77, 0x22, 0x0b, 0x2f,
		0xc3, 0xef, 0xc9, 0x2b, 0xcf, 0x2f, 0x37, 0x24, 0x17, 0x42, 0xd2, 0xdb,
		0xe2, 0x56, 0xf6, 0xee, 0x15, 0xad, 0xeb, 0x16, 0xe3, 0xb3, 0xb4, 0xc2,
		0xaf, 0xc7, 0xf4, 0xe2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x99,
		0x6b, 0x44, 0x29, 0x0b, 0x00, 0x00,
	},
		"templates/jenkins/multi-job.xml",
	)
}

func templates_jenkins_normal_job_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xec, 0x18,
		0x5d, 0x6f, 0xdb, 0x38, 0xf2, 0xdd, 0xbf, 0x42, 0xc8, 0x15, 0xc8, 0xdd,
		0xe2, 0x2c, 0x37, 0xd9, 0x3b, 0xec, 0x1e, 0xe0, 0x78, 0xdb, 0x38, 0x6e,
		0xd7, 0x87, 0x24, 0x0d, 0x62, 0xa7, 0x7d, 0x5c, 0xd0, 0xd2, 0xd8, 0xe6,
		0x9a, 0x12, 0x05, 0x92, 0x4a, 0xeb, 0xcb, 0xf5, 0xbf, 0xef, 0xf0, 0x53,
		0x94, 0xa5, 0xb4, 0x09, 0xb0, 0x2f, 0x0b, 0x04, 0x7e, 0x11, 0x67, 0x86,
		0xc3, 0xf9, 0xfe, 0xf0, 0xf8, 0x97, 0x2f, 0x05, 0x4b, 0xee, 0x41, 0x48,
		0xca, 0xcb, 0xb3, 0xe3, 0x93, 0xf4, 0xf5, 0x71, 0x02, 0x65, 0xc6, 0x73,
		0x5a, 0x6e, 0xce, 0x8e, 0xef, 0x96, 0xef, 0x86, 0x3f, 0x1f, 0xff, 0x32,
		0x19, 0x8c, 0x2b, 0xc1, 0x7f, 0x87, 0x4c, 0x4d, 0x06, 0x49, 0x32, 0x26,
		0x99, 0x42, 0x62, 0x39, 0x32, 0x87, 0x1c, 0x64, 0x26, 0x68, 0xa5, 0x21,
		0x16, 0xb0, 0x03, 0xa8, 0x2e, 0xa0, 0x82, 0x32, 0x47, 0x3e, 0x14, 0xe4,
		0x64, 0x4d, 0x98, 0x84, 0xf1, 0xa8, 0x03, 0xd7, 0xc4, 0xc8, 0xb6, 0x02,
		0xa1, 0xdc, 0x11, 0x01, 0x12, 0xd2, 0x9c, 0x92, 0x15, 0x67, 0xe9, 0xef,
		0x50, 0xee, 0x68, 0x29, 0xd3, 0x8a, 0x56, 0xc0, 0x68, 0x09, 0xe9, 0x8d,
		0xfb, 0xb8, 0xb1, 0x77, 0xf6, 0x49, 0xc5, 0xea, 0x0d, 0x2d, 0xcf, 0x8e,
		0x72, 0x04, 0xa3, 0x06, 0xfb, 0xa1, 0x27, 0x1d, 0x5a, 0xcc, 0x9b, 0xd7,
		0xe9, 0xcf, 0xe9, 0x4f, 0x47, 0x96, 0x33, 0xf2, 0x56, 0x44, 0xee, 0xae,
		0x49, 0x01, 0x93, 0x87, 0x87, 0x24, 0x5d, 0xba, 0x43, 0xf2, 0xf5, 0xeb,
		0x78, 0x14, 0x30, 0x9e, 0x54, 0x2a, 0xb2, 0x81, 0x40, 0xbb, 0xf0, 0x27,
		0x43, 0xdc, 0xe0, 0xac, 0xc8, 0xa3, 0x67, 0xc8, 0x6c, 0x94, 0x1e, 0xb5,
		0xb5, 0xc6, 0x17, 0xe8, 0x3a, 0x49, 0xe7, 0x72, 0x5e, 0x52, 0x45, 0x09,
		0xfb, 0x2f, 0x5f, 0xe1, 0x3b, 0x9a, 0x50, 0x66, 0x45, 0x92, 0x31, 0x22,
		0xe5, 0xd9, 0xd1, 0xb6, 0xce, 0x25, 0x2f, 0x53, 0xab, 0x98, 0x4c, 0x37,
		0x54, 0xa5, 0xef, 0xa9, 0x5a, 0x4c, 0xaf, 0x8e, 0x82, 0x19, 0x10, 0xf6,
		0xe6, 0x34, 0x3d, 0x4d, 0xff, 0xe3, 0x34, 0x1e, 0x67, 0xbc, 0x5c, 0xd3,
		0xcd, 0x47, 0xeb, 0xdb, 0xc9, 0xe9, 0x78, 0xd4, 0x06, 0x58, 0xa2, 0x5a,
		0x82, 0xb8, 0x85, 0x82, 0x2b, 0x98, 0x1a, 0xac, 0x0c, 0x36, 0xe8, 0x79,
		0xf2, 0xee, 0x80, 0xd8, 0xd3, 0x6a, 0x3e, 0x82, 0x19, 0x5b, 0xa1, 0x54,
		0x77, 0xb7, 0x97, 0xc6, 0x50, 0x1a, 0xe4, 0x99, 0x8d, 0x9e, 0xcc, 0x0d,
		0xef, 0xf5, 0x8b, 0x34, 0x5e, 0x09, 0x52, 0x66, 0x5b, 0xf8, 0xa6, 0x84,
		0xe7, 0x86, 0x66, 0x51, 0x41, 0x16, 0xc9, 0x56, 0x6a, 0x67, 0xfd, 0x30,
		0xd2, 0xe2, 0x35, 0x78, 0xba, 0xa6, 0x20, 0x8c, 0x9c, 0x65, 0xec, 0xf9,
		0x3e, 0x41, 0x0f, 0x99, 0x8e, 0x47, 0x6d, 0x51, 0xc6, 0x39, 0x95, 0x64,
		0xc5, 0x60, 0x51, 0xaf, 0x0a, 0x9e, 0xd7, 0xac, 0x09, 0xfa, 0x2e, 0xc2,
		0x5e, 0x10, 0x90, 0xd5, 0xe8, 0x85, 0xfb, 0x18, 0xa3, 0x44, 0x8d, 0x37,
		0xfa, 0x30, 0xee, 0x11, 0xfe, 0x1e, 0x4a, 0x10, 0x44, 0x35, 0x28, 0x6b,
		0x9f, 0x1a, 0x61, 0x3a, 0x21, 0xc3, 0xa3, 0xdf, 0x25, 0xb4, 0x0c, 0x49,
		0xad, 0xb6, 0x5c, 0x7c, 0x10, 0x53, 0x5e, 0x14, 0x54, 0x29, 0x10, 0x9e,
		0x41, 0x17, 0xe1, 0x02, 0x8a, 0x01, 0x29, 0x8d, 0x97, 0xa7, 0xfa, 0xeb,
		0x13, 0x17, 0x3b, 0x59, 0x91, 0xcc, 0xa6, 0x85, 0x45, 0x5a, 0xc2, 0xcf,
		0x18, 0xf6, 0x1f, 0x6a, 0x15, 0x08, 0x3c, 0xe3, 0x0e, 0xdc, 0x92, 0x57,
		0xa2, 0x2e, 0xe1, 0xdc, 0x9b, 0xd4, 0xd1, 0xb6, 0x81, 0xde, 0x6c, 0x3a,
		0x2a, 0x6e, 0x38, 0x63, 0x9e, 0x2a, 0x82, 0x58, 0x12, 0xba, 0x29, 0xb9,
		0x80, 0x6b, 0xae, 0xe8, 0x7a, 0x6f, 0xe5, 0xf7, 0xa4, 0x3d, 0x98, 0x90,
		0x02, 0x8b, 0x2d, 0x61, 0x8c, 0x7f, 0x9e, 0x32, 0x5e, 0x82, 0x73, 0xc4,
		0x21, 0xd4, 0xc5, 0x60, 0x4d, 0x59, 0x3e, 0xdd, 0x72, 0x8e, 0x31, 0xfa,
		0x8d, 0xdc, 0xac, 0x15, 0x65, 0xe9, 0x05, 0xac, 0x49, 0xcd, 0xd4, 0x79,
		0x74, 0xe5, 0x68, 0xe4, 0xf8, 0x20, 0xd1, 0x92, 0x73, 0x36, 0x71, 0x34,
		0xe3, 0x91, 0x07, 0xb8, 0x32, 0x18, 0x1c, 0xb7, 0xde, 0xf8, 0x67, 0x18,
		0x95, 0x2a, 0xdc, 0x17, 0xc0, 0xd0, 0x97, 0xf7, 0xb0, 0x24, 0x62, 0x03,
		0xea, 0x82, 0x8a, 0x06, 0xb1, 0x06, 0x81, 0x15, 0x16, 0x3c, 0x00, 0xbe,
		0x64, 0xac, 0xce, 0x21, 0xbf, 0x85, 0x4d, 0xa8, 0xda, 0x11, 0x58, 0x27,
		0xa1, 0x8c, 0xa4, 0xb2, 0x91, 0xa2, 0xcb, 0x5b, 0x07, 0x38, 0x2b, 0x08,
		0x65, 0x1e, 0x2a, 0x77, 0xb4, 0x5a, 0x92, 0x8d, 0x33, 0x96, 0x3f, 0x39,
		0x1f, 0x94, 0xad, 0x27, 0x27, 0xae, 0xc2, 0x69, 0xb7, 0x63, 0x5b, 0x41,
		0x61, 0x31, 0x60, 0x74, 0x14, 0x1d, 0x02, 0xb0, 0x37, 0x98, 0x50, 0x3a,
		0xbc, 0xef, 0x5e, 0xcc, 0x8a, 0x58, 0xac, 0x3f, 0x3d, 0x1d, 0xbe, 0x6b,
		0x73, 0xf8, 0xa2, 0xa0, 0x94, 0x4d, 0xeb, 0x1b, 0xa1, 0x44, 0xe6, 0x43,
		0x09, 0xba, 0xd9, 0xa0, 0x19, 0x1d, 0x9d, 0x8b, 0x07, 0x0f, 0x4d, 0xb1,
		0x4a, 0x2f, 0xed, 0x77, 0xd3, 0x5f, 0x74, 0x2d, 0xf9, 0x21, 0x71, 0x3f,
		0xe4, 0x14, 0x15, 0x2c, 0x17, 0xc2, 0x37, 0x5c, 0x2a, 0x1b, 0xa6, 0xbf,
		0x72, 0xbe, 0x93, 0xed, 0x28, 0x3e, 0x44, 0x0e, 0x5a, 0xb5, 0xeb, 0x91,
		0x97, 0xb1, 0xcf, 0x45, 0x82, 0xda, 0x50, 0xfe, 0x24, 0x48, 0x55, 0x75,
		0x44, 0xf7, 0xa1, 0x4c, 0x50, 0xdd, 0x8c, 0x33, 0x2e, 0xd2, 0xb7, 0xf8,
		0x35, 0xd5, 0x5f, 0xe7, 0xd1, 0xad, 0xd0, 0x7a, 0x02, 0x1d, 0xb6, 0xdc,
		0x7f, 0xa5, 0xaf, 0x9b, 0x96, 0x6b, 0x80, 0x57, 0xa4, 0x32, 0xed, 0x12,
		0xcd, 0x27, 0x0a, 0xdd, 0x85, 0x22, 0xd8, 0xa0, 0xb7, 0xe8, 0x7e, 0xe7,
		0xdd, 0xa7, 0xcf, 0x0a, 0xae, 0xd7, 0xa1, 0xbf, 0x51, 0xf5, 0x55, 0xad,
		0xb8, 0x78, 0xf6, 0xd4, 0xe0, 0x66, 0xa3, 0x25, 0x14, 0x15, 0xa6, 0x1c,
		0x4c, 0x5e, 0x3d, 0x9c, 0xdf, 0xcd, 0x2f, 0x2f, 0x7e, 0xbb, 0xbe, 0xbb,
		0x3a, 0x9f, 0xdd, 0x7e, 0x1d, 0xbe, 0x7a, 0x78, 0x3f, 0x5f, 0xfe, 0x76,
		0x3b, 0xfb, 0x38, 0x5f, 0xcc, 0x3f, 0x5c, 0xff, 0x93, 0x41, 0xb9, 0x51,
		0xdb, 0xb3, 0x9f, 0x30, 0x8c, 0x0f, 0x6f, 0x7a, 0x8e, 0x75, 0x95, 0xe3,
		0xf1, 0x82, 0x4a, 0x04, 0xef, 0x8d, 0x15, 0x5c, 0xc5, 0xe9, 0xc0, 0x9f,
		0x3e, 0x60, 0x74, 0x15, 0xb5, 0x1e, 0xef, 0x38, 0x59, 0xa7, 0x19, 0x06,
		0xd2, 0x5f, 0x70, 0xc2, 0x30, 0x6d, 0xba, 0xdd, 0xad, 0x6d, 0xbd, 0x33,
		0xc9, 0xa3, 0x5b, 0xc1, 0x5a, 0xb6, 0xfb, 0xfe, 0xcb, 0x4c, 0xf2, 0x32,
		0x93, 0xbc, 0xcc, 0x24, 0x2f, 0x33, 0x49, 0xd2, 0x9d, 0x49, 0x46, 0x3d,
		0x43, 0x45, 0xd3, 0xd0, 0xff, 0xda, 0xad, 0xf1, 0x91, 0xaa, 0x6f, 0x86,
		0xab, 0x41, 0xb3, 0x70, 0x2e, 0x18, 0xb9, 0x87, 0x4b, 0xb2, 0x02, 0xe6,
		0x9a, 0x01, 0x7a, 0x14, 0xe3, 0x11, 0xf2, 0x6b, 0x9e, 0xbb, 0xbd, 0x37,
		0xa6, 0xc0, 0x14, 0x8c, 0xf1, 0x2d, 0xa6, 0x5a, 0x1d, 0x52, 0xde, 0x72,
		0x52, 0xf8, 0xc0, 0xf6, 0xc7, 0x41, 0x53, 0x89, 0xf2, 0x83, 0x02, 0x94,
		0x5b, 0x53, 0x33, 0x9e, 0xed, 0xac, 0x06, 0x5b, 0x28, 0x2f, 0xf8, 0xe7,
		0x52, 0x2a, 0x01, 0xa4, 0x30, 0x20, 0x1c, 0x0e, 0xfd, 0xa5, 0xef, 0xd2,
		0x75, 0x99, 0xdd, 0x55, 0x4f, 0x61, 0xd5, 0xa1, 0x1a, 0xd8, 0xc6, 0x86,
		0x25, 0x0f, 0x63, 0xd6, 0x66, 0x4a, 0xd0, 0xea, 0x00, 0x1c, 0x62, 0x25,
		0x84, 0x09, 0xda, 0x04, 0xcb, 0xc3, 0x06, 0x92, 0xf4, 0x2d, 0x2e, 0xf8,
		0x6b, 0x92, 0xa9, 0x0b, 0xa8, 0xac, 0x89, 0xba, 0x31, 0x94, 0xf1, 0x6a,
		0x4f, 0x1c, 0x59, 0x3a, 0xc5, 0x83, 0xbf, 0x13, 0xa2, 0x27, 0xa6, 0x78,
		0x73, 0x92, 0xfe, 0x78, 0x9a, 0x9e, 0x34, 0x11, 0xe4, 0xff, 0x93, 0xd1,
		0xae, 0xba, 0xb1, 0xdf, 0x3a, 0x70, 0xcc, 0xbc, 0x61, 0x1c, 0x16, 0xfd,
		0x69, 0x63, 0x2e, 0xac, 0x29, 0xd3, 0x85, 0x53, 0xd3, 0x87, 0x97, 0x34,
		0x9d, 0x83, 0x37, 0xff, 0x93, 0xe8, 0xfc, 0x1d, 0x85, 0xb3, 0x4b, 0x4b,
		0xd9, 0x40, 0x24, 0x30, 0xe4, 0xcb, 0x1f, 0xab, 0x36, 0x2d, 0xbd, 0xdc,
		0x00, 0x0a, 0xb9, 0x31, 0xd9, 0xc2, 0xdd, 0x3c, 0x8a, 0x5a, 0x21, 0x1a,
		0x97, 0xad, 0x48, 0xb6, 0x5b, 0xf2, 0x4b, 0x22, 0xd5, 0xa2, 0xce, 0x32,
		0x90, 0x72, 0x5d, 0x33, 0x97, 0xc7, 0x8f, 0xa2, 0xa3, 0x46, 0xef, 0x7c,
		0xf8, 0xce, 0x28, 0xb2, 0x50, 0xba, 0xcf, 0x6c, 0xf6, 0x13, 0xac, 0x23,
		0xef, 0x19, 0x5f, 0x11, 0xb6, 0x00, 0xa5, 0xd0, 0xb5, 0x7a, 0xc4, 0xea,
		0x25, 0x0c, 0x8d, 0xd6, 0x2b, 0x16, 0x20, 0x39, 0xc7, 0xe2, 0xfc, 0x0e,
		0xef, 0x82, 0xa8, 0x04, 0x2d, 0x95, 0xb7, 0x5b, 0xd4, 0xd8, 0x1e, 0x23,
		0xe8, 0xcf, 0xdf, 0x47, 0x7d, 0x1e, 0xe2, 0x27, 0xe4, 0x54, 0xb4, 0x4a,
		0x10, 0xb9, 0xc3, 0x69, 0x7e, 0x0b, 0x8c, 0x45, 0xf5, 0xa3, 0x28, 0x48,
		0x89, 0x31, 0xf8, 0xb7, 0xc4, 0x4f, 0x83, 0x49, 0x86, 0xba, 0x69, 0xb7,
		0x14, 0x64, 0x43, 0xb3, 0x01, 0x7c, 0xa9, 0xb8, 0x50, 0xc9, 0xcd, 0xfc,
		0x66, 0x76, 0x39, 0xbf, 0x9e, 0xb9, 0x89, 0xf5, 0xec, 0xd5, 0xdf, 0x21,
		0xdb, 0xf2, 0xe4, 0xe8, 0xd5, 0x43, 0xc0, 0x7c, 0x9c, 0xdd, 0xea, 0xe1,
		0xf5, 0xeb, 0x51, 0xf2, 0xff, 0x24, 0xab, 0x55, 0x32, 0x5c, 0x9f, 0x24,
		0xc3, 0xfc, 0x78, 0x78, 0xfc, 0x0f, 0xcf, 0x44, 0xcf, 0xb8, 0x8b, 0x5f,
		0xdf, 0x3e, 0xed, 0xf2, 0xa9, 0xbb, 0xdc, 0xff, 0xcf, 0x56, 0x34, 0x81,
		0x62, 0x19, 0x4f, 0x04, 0x48, 0xc0, 0x4b, 0xc3, 0x2d, 0x11, 0x79, 0x62,
		0x87, 0x69, 0x7c, 0xa8, 0x59, 0x07, 0x07, 0xa6, 0xb7, 0x5b, 0x65, 0xbd,
		0x5d, 0xb4, 0x5d, 0x83, 0xfe, 0x2d, 0x33, 0x1f, 0x58, 0xca, 0x95, 0x42,
		0xbf, 0xf2, 0x54, 0xf5, 0x0a, 0x3b, 0xd6, 0xb6, 0xbb, 0xaa, 0x99, 0x5b,
		0x57, 0xd8, 0x4a, 0xa2, 0xca, 0x5d, 0x98, 0x23, 0x66, 0xdd, 0x49, 0x54,
		0xb5, 0x71, 0x14, 0xa2, 0x15, 0xc5, 0x02, 0x10, 0x65, 0x43, 0x8e, 0x43,
		0xb7, 0x6d, 0xe1, 0x33, 0xbd, 0x50, 0xdc, 0x61, 0x65, 0xd2, 0xf5, 0xcd,
		0x96, 0x08, 0x1b, 0xc3, 0xdf, 0x24, 0x69, 0xb2, 0xaa, 0xcc, 0x97, 0x7c,
		0x5e, 0xe6, 0xf4, 0x9e, 0xe6, 0x35, 0x86, 0x97, 0x6f, 0x64, 0x1d, 0x78,
		0x9f, 0xce, 0x56, 0xfa, 0x10, 0x45, 0xda, 0xf0, 0x51, 0x92, 0xf7, 0xa8,
		0xeb, 0xb1, 0x6f, 0x45, 0xb6, 0xd5, 0x9b, 0x50, 0x90, 0x83, 0x84, 0x10,
		0xee, 0x54, 0x0a, 0xd2, 0x8e, 0x6e, 0x24, 0xe6, 0x25, 0xdb, 0xcf, 0xd7,
		0x9d, 0xa4, 0xed, 0x80, 0x3d, 0xbd, 0xde, 0x84, 0xa4, 0xfa, 0x80, 0x68,
		0x47, 0x19, 0x01, 0x82, 0x00, 0x7a, 0xc2, 0x99, 0x15, 0x95, 0xda, 0x3b,
		0xd9, 0xc2, 0x10, 0xd8, 0x41, 0xf4, 0x59, 0xa2, 0x5f, 0xb1, 0x76, 0x66,
		0x39, 0x0b, 0x5d, 0xe3, 0x4e, 0x7f, 0x45, 0xca, 0xda, 0xc4, 0xa6, 0x0c,
		0x76, 0x22, 0x35, 0xe6, 0x69, 0x91, 0x66, 0xe8, 0x67, 0x51, 0x17, 0x72,
		0x2f, 0x15, 0x14, 0x32, 0x6d, 0x65, 0x72, 0x6a, 0x02, 0x2b, 0xec, 0x61,
		0x6e, 0xb3, 0x4e, 0x8d, 0x47, 0xfd, 0x52, 0xe6, 0x4a, 0x5f, 0x88, 0x28,
		0x73, 0xa5, 0xb3, 0x6f, 0x9e, 0xe0, 0x5c, 0xf0, 0x63, 0x3c, 0x17, 0x98,
		0xfd, 0x22, 0x0e, 0x2f, 0xdf, 0xeb, 0xa2, 0x22, 0x6f, 0x7d, 0xd3, 0x11,
		0x5f, 0x47, 0x5a, 0x2f, 0xb5, 0xb3, 0xd3, 0x9f, 0xab, 0xd8, 0x23, 0x86,
		0x75, 0x0d, 0x50, 0x0b, 0xd7, 0xb2, 0xea, 0x41, 0x25, 0xac, 0x88, 0x40,
		0xd1, 0xb0, 0x0c, 0xd3, 0xff, 0x41, 0xde, 0x7a, 0xe6, 0xd0, 0x6e, 0x2d,
		0xca, 0xa1, 0x23, 0xd5, 0x9b, 0xe8, 0xbf, 0x3b, 0x66, 0x8b, 0x9a, 0xc2,
		0x73, 0x9f, 0x3b, 0x5c, 0x38, 0xfb, 0x98, 0x3e, 0x99, 0xb1, 0x9e, 0x91,
		0x6e, 0x3c, 0xa2, 0x71, 0xa6, 0x2f, 0x5d, 0x1d, 0xb6, 0xbe, 0x99, 0x5b,
		0xbf, 0xc6, 0xdd, 0x5b, 0x1e, 0x4a, 0x94, 0x53, 0xbd, 0x30, 0x4d, 0x16,
		0x77, 0xd3, 0xe9, 0x6c, 0xb1, 0x30, 0xdc, 0x1c, 0x28, 0x26, 0x74, 0x82,
		0x7c, 0xa2, 0x6a, 0x7b, 0xcd, 0x1b, 0x49, 0x7c, 0x26, 0x3d, 0x86, 0x6e,
		0xcc, 0x77, 0xd8, 0xb8, 0x9e, 0x69, 0xbf, 0x03, 0x2d, 0x9f, 0xcf, 0xae,
		0x1b, 0x5c, 0x68, 0x92, 0xa8, 0x7e, 0x47, 0xe3, 0xcd, 0x1f, 0x01, 0x00,
		0x00, 0xff, 0xff, 0x9f, 0x80, 0x76, 0xc2, 0xc5, 0x1a, 0x00, 0x00,
	},
		"templates/jenkins/normal-job.xml",
	)
}

func templates_jenkins_pipeline_xml() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xa4, 0x53,
		0xdb, 0x6e, 0x13, 0x31, 0x10, 0x7d, 0xcf, 0x57, 0x58, 0x2b, 0x5e, 0xe3,
		0x6d, 0x41, 0x42, 0x7d, 0x70, 0x5c, 0x50, 0xa0, 0x12, 0x88, 0x4b, 0x11,
		0x85, 0x57, 0xe4, 0xee, 0x4e, 0x1c, 0x83, 0x77, 0x6c, 0xf9, 0x92, 0xa4,
		0xaa, 0xfa, 0xef, 0xcc, 0x5e, 0xdc, 0x76, 0xcb, 0x45, 0xaa, 0xfa, 0x94,
		0xf8, 0x9c, 0x33, 0x67, 0xce, 0x78, 0xbc, 0xe2, 0xf4, 0xd0, 0x59, 0xb6,
		0x83, 0x10, 0x8d, 0xc3, 0x55, 0x75, 0xcc, 0x8f, 0x2a, 0x06, 0xd8, 0xb8,
		0xd6, 0xa0, 0x5e, 0x55, 0xdf, 0x2e, 0xce, 0x96, 0x27, 0xd5, 0xa9, 0x5c,
		0x88, 0x08, 0xbc, 0x35, 0xea, 0xd2, 0x59, 0xfe, 0x13, 0xf0, 0x97, 0xc1,
		0xc8, 0xbd, 0xf1, 0x60, 0x0d, 0x02, 0x7f, 0x43, 0x3f, 0x64, 0x70, 0x75,
		0x3e, 0x01, 0xdf, 0x0d, 0xec, 0x99, 0xb7, 0x59, 0x1b, 0x32, 0x6c, 0x27,
		0x72, 0x59, 0xe4, 0xcb, 0x91, 0x79, 0x75, 0xc4, 0x4f, 0xf8, 0xcb, 0x4a,
		0x2e, 0x18, 0x13, 0xa8, 0x3a, 0x90, 0xd7, 0xd7, 0x8c, 0x7f, 0xa2, 0x3f,
		0xec, 0xe6, 0x46, 0xd4, 0x03, 0xd2, 0x53, 0x1b, 0x63, 0x13, 0x84, 0xb7,
		0x07, 0x68, 0x72, 0x72, 0x21, 0xca, 0x8d, 0xb2, 0x11, 0x44, 0xfd, 0x10,
		0xbe, 0x93, 0x7e, 0xc9, 0x90, 0x61, 0x2e, 0x1b, 0xa1, 0x5e, 0xe2, 0x83,
		0xf3, 0x10, 0x92, 0x81, 0xc8, 0x1a, 0xab, 0x62, 0x5c, 0x55, 0xdb, 0xdc,
		0x46, 0x87, 0xbc, 0x73, 0x94, 0x93, 0xf7, 0xc1, 0x9f, 0x9d, 0x8f, 0x9a,
		0xab, 0x0f, 0x26, 0xa6, 0xaa, 0x1e, 0xca, 0x1a, 0xd7, 0x79, 0x87, 0x80,
		0xe9, 0xab, 0x87, 0x66, 0x68, 0x46, 0xe0, 0x23, 0x2f, 0xe4, 0xc7, 0x72,
		0x7d, 0xdf, 0x65, 0x34, 0xf9, 0xef, 0xec, 0x03, 0xbd, 0x31, 0x21, 0xa6,
		0xf7, 0xee, 0x72, 0x2e, 0xb9, 0x45, 0xc7, 0x28, 0xf5, 0xd3, 0xb3, 0x88,
		0xfa, 0xcf, 0x21, 0x05, 0xba, 0xcf, 0x9b, 0x52, 0x17, 0xe5, 0x0b, 0x8a,
		0x36, 0x03, 0x7a, 0x49, 0xdc, 0xba, 0xfd, 0x6b, 0xad, 0x03, 0x68, 0x95,
		0xa0, 0x2d, 0x5c, 0xb9, 0xff, 0x7f, 0xb0, 0xc5, 0x7b, 0xed, 0x6c, 0xee,
		0x30, 0xca, 0xe3, 0xd1, 0xb9, 0x1c, 0x07, 0x5f, 0x47, 0x6b, 0x42, 0x2d,
		0x91, 0x22, 0x91, 0xcf, 0x74, 0xba, 0xed, 0xb8, 0x53, 0x49, 0xdd, 0xbd,
		0x86, 0xfb, 0x50, 0x2f, 0xc9, 0xbe, 0xa5, 0x7e, 0xef, 0x90, 0x76, 0xbf,
		0x53, 0x56, 0x3e, 0x17, 0xf5, 0x03, 0xa4, 0xf8, 0xac, 0xb7, 0x0a, 0x35,
		0xcc, 0x7c, 0x0a, 0xd4, 0x4b, 0x94, 0xb5, 0x6e, 0xff, 0x51, 0x61, 0x56,
		0xf6, 0x22, 0x18, 0xad, 0xe9, 0x23, 0x91, 0x29, 0x64, 0x52, 0xfe, 0x8d,
		0xe9, 0x2b, 0x68, 0x52, 0x38, 0xf8, 0xb3, 0x69, 0x3d, 0x91, 0x9e, 0xcf,
		0xa3, 0x97, 0x23, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x06, 0x92,
		0x4d, 0x61, 0x93, 0x03, 0x00, 0x00,
	},
		"templates/jenkins/pipeline.xml",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func assetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"templates/jenkins/multi-job.xml":  templates_jenkins_multi_job_xml,
	"templates/jenkins/normal-job.xml": templates_jenkins_normal_job_xml,
	"templates/jenkins/pipeline.xml":   templates_jenkins_pipeline_xml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func assetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"jenkins": &_bintree_t{nil, map[string]*_bintree_t{
			"multi-job.xml":  &_bintree_t{templates_jenkins_multi_job_xml, map[string]*_bintree_t{}},
			"normal-job.xml": &_bintree_t{templates_jenkins_normal_job_xml, map[string]*_bintree_t{}},
			"pipeline.xml":   &_bintree_t{templates_jenkins_pipeline_xml, map[string]*_bintree_t{}},
		}},
	}},
}}
