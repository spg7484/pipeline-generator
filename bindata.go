// Code generated by go-bindata.
// sources:
// templates/jenkins/multi-job.xml
// templates/jenkins/normal-job.xml
// templates/jenkins/pipeline.xml
// DO NOT EDIT!

package pipeline

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesJenkinsMultiJobXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x59\xe9\x6f\xdb\x38\x16\xff\x5c\xff\x15\x86\x31\x40\x81\xc1\x46\x4a\x33\x3b\x33\x5d\x40\xc9\x34\x71\xd2\x6e\x16\x39\xbc\x3e\xa6\x1f\x0b\x5a\xa2\x6d\xd6\x94\x28\x90\x54\x1a\x6f\xd1\xff\x7d\x1f\x4f\x51\x87\x1b\x27\x5b\x2c\x26\xf9\x10\xe9\x5d\x7c\x7c\xc7\x8f\x8f\x4a\xf2\xc7\x63\x4e\x87\x0f\x98\x0b\xc2\x8a\xd3\xd7\x6f\xa2\xe3\xd7\x43\x5c\xa4\x2c\x23\xc5\xfa\xf4\xf5\x62\xfe\xfe\xe8\xed\xeb\x3f\xce\x06\x49\xca\xf2\x48\x92\x2d\xa2\xd1\x67\x5c\x6c\x49\x21\xa2\x92\x56\x6b\xf5\x37\xaf\xa8\x24\x9f\xd9\x32\xba\x55\x0f\xff\x62\xcb\x09\x67\x9f\x71\x2a\x87\x46\xe0\x74\x64\x15\x8e\x9c\xe0\x91\x61\xbc\x7b\x13\xbd\xf9\x75\x74\x36\x18\x0e\x13\x94\x4a\x58\x5d\xc4\xfa\x25\xc3\x22\xe5\xa4\x54\x94\xb3\x24\x0e\xdf\x14\x77\x8b\x71\x79\x89\x4b\x5c\x64\xe0\x25\xc1\xe2\x6c\x85\xa8\xc0\x49\xdc\xa1\x2b\xe1\x92\xb3\x12\x73\x69\x5f\x81\x20\x70\x94\x11\xb4\x64\xc1\x2e\x48\x89\x29\x29\x70\x34\xb1\x0f\x13\xa3\xb3\xf3\xee\x67\x40\x86\xf8\xec\x8e\x9c\xa8\xf3\xff\x38\x7a\x1b\xfd\x36\x32\x96\xc1\xb6\x44\x62\x7b\x87\x72\x7c\x56\x22\x8e\x28\xc5\x74\x88\x1f\x71\x5a\x29\xcf\x93\xd8\x33\x9d\xb4\x90\x68\x8d\x35\xe5\xeb\xd7\x61\x34\x73\x6f\xc3\x6f\xdf\x92\xb8\xe6\x19\xaf\xe3\x67\xb8\x6d\x55\x54\xba\x52\xc6\xd1\xc3\x2e\xda\x54\x99\x60\x85\x4f\xd7\x9a\xc8\x4d\xb5\x8c\x3e\xe8\x3f\x36\x55\x9d\x3d\x1b\x21\x95\xa1\x93\xe8\xb8\xde\x62\x69\xc4\x17\x9c\x6a\xaf\x8d\x0d\x78\xd3\x5e\x07\x4c\xeb\xf6\xcb\x9c\xd0\x99\x8b\x9b\xa9\x4b\x44\x9a\x0f\x53\x8a\x84\x38\x1d\x75\x4d\x29\x3b\xb3\xf1\xed\x28\x74\xff\xdd\x49\xf4\x77\xed\xfa\xab\x57\x10\x8c\x62\x45\xd6\x7f\x9a\x1a\x3f\x3b\x51\x8e\x85\x04\x25\x52\x09\xcc\xa7\x38\x67\x12\x8f\x35\x4f\x28\xea\xab\xa4\x67\xad\x45\x4b\x52\x0b\x82\x81\x3a\x26\x8b\xe9\x8d\x0e\x88\x22\x69\x2b\xf1\x81\x66\x40\xa5\xcf\x8d\x64\xc9\x51\x91\x6e\xf0\x7e\x9f\x2e\xb4\xc0\xac\xc4\xa9\xf5\xa6\x50\xc5\xf3\x73\xac\x1c\xaa\x79\x64\x45\x30\xd7\x9e\x69\xf6\x60\x8f\x67\x4d\x63\x49\x1c\xae\x9e\x64\x44\xa0\x25\xc5\xb3\x6a\x99\xb3\xac\xa2\x75\xff\x75\x19\xa6\x08\x38\xb4\x00\x84\xf9\x21\xe4\x48\x5e\x81\x46\x1f\xc7\xe8\x64\xec\x03\x2e\x30\x47\xb2\x66\x99\x70\x54\x40\x53\x40\xe1\x17\x7d\x52\xd0\x18\x44\x95\xdc\x30\x7e\xcf\xc7\x2c\xcf\x89\x94\x98\x3b\x03\x5d\x86\x6d\x1f\x8a\x51\xa1\xf3\x39\x56\x4f\x1f\x19\xdf\x8a\x12\xa5\xa6\x3d\x0d\xd3\x08\x7e\x81\xf6\xbb\xaf\xa4\x17\x70\x86\x3b\xf4\x81\x6d\xa0\xaa\xc0\x17\x2e\xa0\x56\xb6\x49\x74\x61\x53\x45\x30\x61\x94\x3a\xa9\x80\x62\x44\xc8\xba\x60\x1c\xdf\x31\x49\x56\x3b\xe3\xbf\x13\xed\xe1\x18\x15\x28\xaf\xd9\x06\xb0\x89\x7d\x19\x53\x56\x78\x67\xdb\x64\x23\xbc\xac\x08\xcd\xc6\x1b\xc6\xa0\x26\xbf\xd3\x7b\x00\x70\x34\xba\xc4\x2b\x04\xe8\x7e\x11\xa8\x8c\x62\x6b\x07\x84\xe6\x8c\xd1\x33\x2b\x93\xc4\x8e\x60\x21\xd9\x67\x6e\xb5\x76\xcb\x50\x22\xa4\xd7\xe7\x98\x42\x32\x1f\xf0\x1c\xf1\x35\x96\x97\x84\xd7\x8c\x15\xe6\x80\xf6\xd8\x11\xf0\x63\x4a\xab\x0c\x67\x53\xbc\xf6\xc7\x49\x40\x56\x2d\x27\x02\xaf\x4c\xa9\x28\x9c\xed\x10\xaf\x72\x44\xa8\xa3\x8a\x2d\x29\xe7\x68\x6d\xcb\xd6\xbd\x59\x5e\x9a\x87\x06\x7e\x78\xe5\x3e\x19\x1d\xfc\x28\x71\x21\x6a\x8d\xe1\x10\xaa\x96\xac\x86\xd1\x0d\x4b\x11\x35\x55\x05\x55\xeb\x10\xbc\x27\x81\xb5\x85\x88\xe4\x25\x0d\x15\x9d\x49\x50\xa4\x01\x75\x1f\xac\xd0\xae\x66\x2f\xc2\x1c\xb0\x22\x2c\x01\x47\xb9\x73\x3c\x89\xc3\x6d\x0e\x00\x94\x20\xee\xf0\xa0\x3a\x15\x15\x53\x86\x72\x9b\x1c\xf7\xa6\xe7\x08\x83\x48\x59\x0b\xa1\x32\xcd\x5c\x82\xaf\x5b\x5d\xae\x1f\x37\xb8\xb8\x64\x5f\x0a\x21\x39\x46\xb9\x26\xc1\xe8\xe3\x94\x9e\x94\xeb\x1a\x5b\x94\x87\x98\xea\x48\x29\x43\x92\x93\xf5\x1a\x8a\xd4\x44\xc1\x26\xf2\x5a\x5c\x17\x44\x12\x44\x61\xb4\xf2\x01\xb1\x51\x75\x0a\x11\x1c\x7f\x73\xf3\x5c\x4f\x18\x0a\xc1\x7f\x1e\xda\x5f\x08\x99\x46\x74\xcb\x34\x10\x31\x61\x42\x1a\x80\xf8\x27\x63\x5b\xd1\xc4\x8f\x36\x73\xd0\xc8\xe7\xde\x95\xc3\xcc\xc1\xdc\x13\xec\x48\x9d\xc3\x00\xf9\xd0\xb2\x06\x28\xdc\x72\x6d\xf2\xc0\x81\x8f\x8f\xc4\x73\xa6\xcf\x0b\xa3\x59\xcf\x2c\x1b\x24\xcc\x30\xa5\x9f\x8e\x01\x6f\x3d\xa5\x21\x03\xba\xa2\x2e\x78\xd8\x06\x14\xe4\x1a\xc3\x6c\x56\x2d\x15\xab\x6e\xa2\xc3\xfc\x99\x38\x9b\xee\x7c\x1f\xfa\x9f\x04\xf8\x7e\xf4\xd3\x9d\xe3\x08\xa1\x90\x8a\xc9\x04\x46\xc9\xdc\x1d\x98\x01\x21\x94\xc3\x8f\x25\xe0\x6d\x06\x79\x70\x01\x0d\x28\xa1\xa0\x6d\x00\xf0\xa9\xd5\x12\x8a\x12\x0a\xc2\x00\x2b\xa0\x24\xa7\xea\x48\x9e\x20\xb9\x81\x31\xbc\x43\x0a\xe5\x73\xf4\x38\xc5\x90\x69\x38\xbd\x20\xbc\xc1\x5b\xc3\xcd\x42\x2d\xa5\x38\xbb\x99\x54\xc0\xb7\xde\x79\x7f\x7b\x58\x5d\x55\x88\x63\x46\xf4\x3d\xa0\xa1\x56\x93\x43\x15\x18\x96\xb9\x3c\xa7\xaa\x67\x6c\xf4\x42\x4a\x23\xcc\x5e\x5f\x57\x62\x8f\xad\x2d\xa1\x54\x67\xf3\xbe\x00\xe5\x29\x16\x90\xe2\x7a\xd5\xf7\xe7\xd7\x37\x8b\xe9\x15\xdc\x40\xbe\x2b\x16\x1a\xd4\xc5\x7d\x5f\xd0\xdd\xf5\x0a\x72\x34\xde\xa8\x3a\xf3\xcd\xd7\xcf\x6c\x79\x6c\x46\xc3\x61\xf0\xd3\x87\xec\x30\x8b\x4e\xf1\x03\x51\xb0\xa9\xbb\x42\x17\x0f\x86\x29\x47\xf4\x4e\xca\x0d\x73\x50\xe1\x4b\xb8\x56\xfc\xbb\xc2\x15\xce\x0c\x0c\x88\xba\x61\x7b\x78\x4d\xf5\x3e\xd8\xdf\xef\x4e\x63\x7b\x71\x67\x7f\xe6\x26\xf1\xd2\x7e\x6b\x1e\x25\xfa\x6e\xd1\x6e\x77\x15\x52\x49\x8a\x4a\x1f\xbd\x75\xce\x66\x8b\xf1\xf8\x6a\x36\x7b\xbf\xb8\xd1\x4e\xf5\x48\x0c\x0e\x76\xaf\x17\x9e\xf6\x82\x4c\x3b\x9b\x29\x2b\x77\x08\x6e\x43\x2b\xb8\x29\x47\x63\x78\x39\xb7\x2f\x3e\x91\xa1\x04\xdc\xdb\x7e\xf9\x35\x3a\xe9\xdc\xdb\x6a\xbc\x71\x04\x27\xb0\x22\x14\xf2\x10\x07\x77\x59\x35\x6b\xd5\xef\x76\x84\x12\x35\x45\x60\x0a\x06\x98\x9f\x0c\x9f\x1b\x82\x99\xd5\x1f\x1d\xf2\x99\xe0\xed\xa8\x5e\x98\xe9\x6f\x01\x88\xda\xb6\xf6\xaf\x8e\x9f\x31\x18\x7c\xdf\x03\x56\x61\x5e\x72\x52\x48\x17\xa9\x60\xf4\xda\x27\x30\xe8\x2d\xdd\xbd\xb1\xef\x3b\xef\xc2\x73\x2b\x29\xab\x25\xcc\x6a\x9b\xf6\x81\xee\x17\xec\x9c\xe6\x48\x6c\x85\x67\x9f\xf3\x74\xa3\x3e\x3c\xf8\x9d\x21\xef\xa8\xca\x63\x68\x05\xd0\xad\xb9\x09\x25\xad\xe6\xf9\xab\xbc\x94\x3b\x6b\xc8\xe1\x60\x87\xee\x23\x6b\x40\xa7\x4a\x53\x2c\xc4\xaa\xf2\x11\x6e\x93\xeb\xa2\xf1\x31\xb4\xa2\x21\xc5\xe7\xc3\x8c\xfd\x57\xb6\x82\xac\x64\x9b\xda\x1a\x2e\xbe\x13\x88\x66\x3b\xdb\x90\xde\xc1\x7c\x78\x8b\xa0\x3d\x69\xa3\x8b\x50\x15\xe9\x6f\x10\x30\x5a\xf0\x2a\x17\x3b\x21\x71\x2e\x9a\xdf\x22\x22\x9d\x31\xff\x39\xc5\xce\x2b\x91\x01\x28\x4b\xb5\xf3\x8d\xaf\x54\xad\xd2\xf9\x1c\xf4\x06\x60\xf4\x97\x51\x08\x29\x2d\x14\xfb\x91\x08\xfd\x3f\xe0\xf3\x0b\xb1\xb9\x83\xcb\xd0\x69\x6e\x1a\xb6\x5f\x70\xd4\x08\x63\x8a\xb3\x93\x0f\xd5\x76\xbd\xd2\x36\xf1\x3f\x36\x53\x4f\x54\xca\xf7\x90\xb6\x74\x1b\x27\xff\xc1\x59\x63\x8d\x76\x15\x34\x24\x8f\xac\x28\xa4\xea\xe4\x1f\xcf\x28\x82\x27\x97\xeb\x99\x1e\xff\x1f\xe7\xff\x5f\x6c\x02\xa8\xcf\xb0\xba\xbe\x7c\x65\x79\x4e\xff\x58\x67\x4f\xf1\x7d\xd3\x9d\x0d\xfa\x47\x22\x37\x77\x2c\x70\xc8\xee\x74\x1f\x7b\x7f\x3b\x3d\x33\xa5\xad\xdd\x3e\xdf\x5c\xdf\x31\xd4\x3c\x79\xcc\xc0\xf9\x91\xa3\xb2\xac\xaf\x54\xad\x65\x10\xdc\xae\x53\x46\x19\x8f\xce\xe1\x69\xac\x9e\x2e\x02\x2d\x5f\x30\x5e\xee\xdd\x71\xa3\x6c\x12\x4d\xbc\x45\xa5\xbe\xc7\xc0\x6d\x9d\xe7\x6a\x67\x01\xad\x7f\x7b\x4f\xac\xfb\xf4\x4d\xf8\x80\x6f\xe3\xf6\x43\x2f\xc4\x1d\xe2\xb7\xac\xd4\xd4\x72\xd0\xc7\xfd\xdf\xeb\xdd\xd9\x7f\x90\xcc\x71\x5e\x52\xb8\xa0\x9c\xfd\xf4\xf5\x62\x71\x7d\x73\xf9\xe9\x6e\x71\x7b\x71\x35\xfd\x76\xf4\xd3\xd7\x0f\xd7\xf3\x4f\xd3\xab\x3f\xaf\x67\xd7\xf7\x77\x7f\xa3\xb8\x58\xcb\xcd\xe9\xef\x50\x9d\x6d\x4d\x67\xb1\x2a\x33\x78\xbd\x24\x02\xc8\x3b\x1d\x20\x73\x32\x76\xe9\x36\x70\x2f\xda\x68\x1f\x12\xb6\x53\x2f\x09\x00\xb1\x44\x39\x84\x3b\x9a\xd7\xcf\xbd\xe9\x0f\x64\xe1\xc0\xfb\x0d\x86\x4c\xf3\x4f\x9b\xb8\x55\x62\xcf\x9a\x89\x27\x6e\x1a\xfd\x6f\x00\x00\x00\xff\xff\x4a\x55\x24\x92\x8d\x1a\x00\x00")

func templatesJenkinsMultiJobXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsMultiJobXml,
		"templates/jenkins/multi-job.xml",
	)
}

func templatesJenkinsMultiJobXml() (*asset, error) {
	bytes, err := templatesJenkinsMultiJobXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/multi-job.xml", size: 6797, mode: os.FileMode(420), modTime: time.Unix(1444928474, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJenkinsNormalJobXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\xff\x6f\xdb\x36\x16\xff\x3d\x7f\x85\x91\x2b\xd0\xbb\x61\x96\x9b\xb4\xb7\x6e\x80\xeb\x25\xb1\x9d\x36\x9b\xe3\xf8\x6c\xa7\xbb\xc3\xe1\x10\xd0\x12\x6d\x6b\xa1\x44\x41\x94\x92\xf8\x7a\xfd\xdf\xef\xf1\xab\x48\x4a\xb6\x93\x75\x18\x50\x2c\x1b\xb0\x45\xef\x1b\xbf\xbd\xf7\xe1\x7b\x24\xdd\xfd\xf1\x21\x21\xad\x3b\x9c\xb3\x98\xa6\xef\x5e\x1e\x05\xaf\x5e\xb6\x70\x1a\xd2\x28\x4e\x57\xef\x5e\x5e\xcf\xcf\xdb\xdf\xbf\xfc\xb1\x77\xd0\xcd\x72\xfa\x2b\x0e\x8b\xde\x41\xab\xd5\x45\x61\x01\xc2\xac\x23\x3e\x22\xcc\xc2\x3c\xce\x38\x45\x12\x08\x5d\x4d\x69\x81\x0a\x9a\xb7\x42\x82\x18\x7b\x77\xb8\x2e\x23\x46\xd3\xa0\x40\xec\x96\x05\x23\xc3\x3e\xe4\xe2\xdc\x02\xda\xb0\x39\xfd\x19\xe3\xac\xd7\x3e\xea\x76\xac\x4f\xc9\x4f\xcb\xc4\x62\x57\x5f\x92\x8b\xf2\x22\x5e\x42\x8f\x06\xae\x95\x06\xb2\x2b\x3f\x36\x76\x2c\xe9\xb1\x6d\xbc\xdb\xa9\x86\x22\xbe\x6f\x81\x31\xc0\x19\x4e\x23\x98\xa1\x18\xb3\xde\x12\x11\x86\xbb\x9d\x1a\x9d\x0b\xc3\x84\x65\x18\xcc\xca\x4f\x20\x30\x1c\x44\x31\x5a\x50\x12\xfc\x8a\xd3\xdb\x38\x65\x41\x16\x67\x98\xc4\x29\x0e\x26\xea\x8f\x89\xd4\xd9\xb4\x32\x52\xae\xe2\xf4\xdd\x61\x04\x64\x58\x9b\x4d\x5b\x8b\xb6\x25\xe7\xe4\x55\xf0\x7d\xf0\x56\x4d\x20\xd8\xe6\x53\x3b\x46\x09\xee\x7d\xfa\xd4\x0a\xe6\xea\xa3\xf5\xf9\x73\xb7\x63\x38\x5a\x94\x15\x68\x85\x8d\xec\x4c\x7f\x09\xe1\x8a\x27\xbb\xdc\x79\x42\x9f\x95\x4a\x48\x93\x20\xa4\x39\xba\xdb\x04\x6a\xd9\x65\x8f\x59\xb0\x8a\x8b\x75\xb9\x08\xde\x8b\xff\x4d\xa4\x3b\xd5\x06\x2c\x85\x4e\x8e\x82\xa3\xe3\xe0\x55\x35\x3e\xe5\x7d\xd7\x39\x11\xbd\x96\x36\xe0\x4b\xf4\xda\x62\xaa\x6e\xff\xb6\x4e\xc8\x35\x77\xd7\x0d\x5a\x8b\x97\xad\xe0\x82\x5d\xa4\x71\x11\x23\xf2\x13\x5d\x40\x9b\x5c\x90\x85\x89\xe7\xdf\x56\x1b\xbc\x81\x59\xff\xf2\xd0\x1e\xd7\xc9\x71\x70\x1c\xfc\x70\x68\xe6\x29\x5d\xc6\xab\x8f\x32\xee\x7a\xc7\xbc\xcf\x36\x41\x0a\x95\x0c\xe7\x53\x9c\xd0\x02\xf7\x05\x97\x99\x09\x69\x68\xf2\xda\x13\xd6\xb2\xdc\x4e\x35\x6f\xd7\xd3\x91\x98\xb4\x52\xcf\x16\x1f\xf3\xa3\xad\x81\x5e\x73\x97\xba\x8b\x1c\xa5\xe1\x1a\xef\xec\xe1\x99\x90\x99\x65\x38\xb4\xfa\x96\x72\x77\xfb\xa6\xc3\xbb\x57\xf1\xe3\x65\x8c\x73\xd1\xcf\xd4\xf6\xdd\xa6\x8e\xfa\x46\xbb\x1d\xb7\x2b\xdd\x28\x66\x68\x41\xf0\xac\x5c\x24\x34\x2a\x49\x15\xb6\x75\x86\x54\xc8\x71\x58\xc2\x2a\xdc\xd9\x9c\x22\x2f\x41\xa3\x89\xa3\x1a\xa1\xef\x71\x8a\x73\x54\x54\x2c\x39\x3f\x25\xd0\x38\x58\x9a\x46\xf7\x0a\x2a\x98\x2a\x8b\x35\xcd\xaf\xf2\x3e\x4d\x92\xb8\x28\x70\xae\x0d\xd4\x19\xca\xa1\x08\x46\xa9\x58\xe5\x3e\xff\xeb\x17\x9a\xdf\xb2\x0c\x85\x32\xb0\x25\x53\x0a\xde\x43\xe0\x5e\x95\x85\x11\xd0\x86\x6b\xf4\x03\x15\x7a\x65\x8a\xcf\xf4\x94\x2a\x59\x97\xa8\xa7\x8d\x7b\xc5\x84\x12\xa2\xa5\x2c\x8a\x14\x89\x57\x29\xcd\xf1\x98\x02\xd8\x6e\x64\xff\xb5\x68\x03\xc7\x84\xc0\x6c\x8d\x08\xa1\xf7\x7d\x42\x53\xd3\x59\x9f\xac\x9c\xb0\x8c\x49\xd4\x5f\x53\x0a\x4e\xba\x23\x38\xcb\x22\x26\xc1\x00\x2f\x51\x49\x8a\x33\x4b\xe5\xb0\xa3\xec\x80\xd0\x9c\x52\xd2\x53\x32\xdd\x8e\x26\x28\x24\x37\x2b\xb7\x5c\xe9\x66\x48\xcc\x0a\xa3\x9f\x63\x02\x8b\x79\x87\xe7\x28\x5f\xe1\x62\x10\xe7\x15\x63\x89\x73\xd8\x24\xb0\x26\xe0\x87\x90\x94\x11\x8e\xa6\x78\x65\xb6\x54\x8b\xcc\xa3\x90\x59\xbd\x92\xae\xc2\x11\xba\x46\x1c\x26\x28\x26\x9a\xca\x6e\xe3\x6c\x8e\x56\xca\x6d\xf5\x97\x5a\x84\xd4\x69\xb2\xa7\x20\x8e\xaf\x3b\xec\xf9\xd0\x59\xf0\x18\xee\x46\x3e\x01\xb6\x37\xe1\x4b\xbe\xbe\x6a\x31\x4c\xec\x6e\xfd\xee\xf1\xb0\x77\xce\xf1\x43\x81\x53\x56\x69\x18\xec\x1e\xd1\x10\x11\xe9\xab\x12\xba\xb7\xc1\x53\x65\x21\x88\x93\x8c\xd8\x8a\x16\x5e\x11\x8b\xba\x0d\xb4\x48\x5d\xb3\x11\xbb\x1e\xd1\xa2\x99\x78\x05\x6e\xee\x30\x61\x6d\xc3\x44\xfc\x51\xe4\xf1\x6a\x05\xce\x22\xf5\xd4\xd0\xaf\x33\x56\xe4\x18\x25\xb0\x6b\x31\x63\x42\xef\xe4\x5a\x23\x98\x62\x9e\xfe\x61\x11\x09\x73\x49\xac\xb2\x05\x18\x57\xc7\x7c\x95\xca\x9e\xda\x38\x85\xef\xf8\x8d\x40\x6c\xfa\x52\x26\x4b\x59\xe7\x98\xad\x29\x89\x7c\xf8\x9f\x5d\xf7\xfb\xc3\xd9\xcc\x45\x7b\x60\xd2\x1c\xd2\x50\x44\x7a\xaf\xba\x1d\xfd\x67\xc5\x0c\x29\x81\xbc\xec\x6c\x74\x3d\xe4\x9b\x27\xa1\xb9\xc3\x83\xe9\xc4\x85\x1c\x93\x8a\x02\x97\x66\x96\xc5\xeb\x54\xb7\xf3\xf8\xf9\x71\xd7\x46\xcd\xb9\x92\xc0\xd1\x25\x4a\x4b\x40\xa8\x8d\x8a\x1e\xf0\x76\xb3\x04\x3a\x21\xd6\x2d\x40\xaa\xd0\x34\xf1\xbd\x6f\x5a\xea\x5f\x58\x67\x6b\xd7\x54\x38\x3a\xa1\xac\x90\x58\xf9\x81\xd2\x5b\xe6\x42\xa9\xcf\x3c\x70\x9c\x70\x6b\xcb\xf6\x90\x60\x6e\x2c\xa7\x92\xe0\xfa\x4b\x8e\xb2\xcc\xb8\x99\x1f\x45\x08\x3c\x53\x2c\x45\x70\x0a\x7f\xf5\xf9\x5f\x67\x96\x96\xc9\x86\x8c\x1c\xe4\xb1\x6f\xec\x3c\x4f\x10\x2f\x51\x26\x72\x50\xf0\xf4\x3c\x51\x6b\xab\x69\x07\x8d\xb1\xb4\xa7\xdd\xc7\x27\xe0\x2a\xfd\x02\x04\x82\xa1\x2f\x4a\x5e\xc3\x3c\x35\x15\x57\xa5\xd4\x1c\x83\xb7\x01\xaa\xf5\x5e\x7c\x3a\xbb\xbe\x18\x0d\x6e\xc6\xd7\x97\x67\xc3\xe9\xe7\xf6\x8b\x4f\xef\x2f\xe6\x37\xd3\xe1\xc7\x8b\xd9\xc5\xd5\xf8\x5b\x82\xd3\x55\xb1\x7e\xf7\x16\xe2\xc6\xd7\xac\xa2\x2e\x82\xcf\x41\xcc\x80\xbc\x11\xb3\x20\x1d\xba\x4e\x7f\x7c\xd6\x5e\x1f\x68\xf3\x8a\x16\x71\x82\xa1\x2a\x48\x60\x16\x83\x79\xf5\x77\xe3\xaa\x5a\xb2\x90\xc0\x7f\x17\x1c\x4b\x6c\x86\x84\xcc\xf7\x1c\x27\x1e\xbe\xaa\x4c\x5a\x00\x54\x0d\xa7\x60\x5b\x17\xf1\xc9\x53\x9e\x25\x73\xf3\xdb\xe7\xdc\xfb\x39\xf7\x7e\xce\xbd\x9f\x73\x6f\xb9\x08\x69\x63\x93\x5e\xf2\xfc\x27\x4f\x65\xbf\xee\x44\x63\xcb\x76\xa7\x07\xae\x56\x72\x06\xc1\xcb\xbd\x58\x6d\x81\x6e\xf6\x5e\xcb\x0f\xf9\xce\x9b\x37\x66\x88\xe2\x0c\xaf\xb2\x65\xe7\x89\xf5\x64\xcf\x37\xe3\x26\x78\x4d\xbd\x24\xe8\x0e\x8f\xd0\x02\x13\xd5\x4f\x08\x37\x40\x0b\x1c\x8d\x69\xa4\xce\x0f\x6d\x09\x00\x48\x9b\xef\xd9\xec\x86\x28\x9d\x52\x94\x68\x18\xd1\x9f\x07\xd5\x3e\x11\x79\xdb\x43\x24\xfd\x01\x7c\xf0\x56\x4e\xf3\x1a\xa7\x03\x7a\x9f\xca\xfa\x42\x90\xa0\x42\xd5\x4a\x7b\xe5\xea\xc6\x74\xed\xb2\xdb\x54\x4d\xea\x40\xa6\x1d\xb0\x21\x01\xa0\x48\x18\x33\xa3\xf2\xc8\x07\x3b\x8a\xb1\x27\x94\x62\x76\x21\xf6\x45\x65\x58\xbd\x08\xdb\x51\x82\xed\x28\xc0\xb6\x97\x5f\x8f\x2d\xbe\xbc\xd2\xeb\xf1\x85\x97\xe3\x54\x22\xd8\xec\xba\x17\x20\x67\x85\x5b\xc1\xa9\x3e\xff\xc7\x99\x5f\x76\xe9\x60\x0e\x69\xb6\xd1\x07\xff\x41\x1f\x3e\xb4\x8e\x81\x0f\x5b\x02\x52\xda\xd7\xc7\xc1\x51\xed\x4c\x5a\x4c\xbe\x9a\x63\x8e\x1c\x22\x7d\xb7\xcf\xa5\x8d\xc2\x32\x26\x3c\x65\xe0\xf2\xa6\x25\x2e\xa7\xe8\xa6\x4a\x16\x1b\x57\x55\x74\xab\xfd\x88\x55\x14\x86\x09\xd8\xad\xdd\xaf\x34\x8e\xcb\x94\xa3\x62\x16\x67\x4a\xf3\xd0\xda\x05\xc0\x71\xc9\x02\x85\xb7\x73\x3a\x42\xac\x98\x95\x61\x88\x19\x5b\x96\x44\xad\xdc\x56\xb6\x95\xe2\x2a\x4f\x3b\x17\x03\x99\x15\x3c\xc3\x5a\x6d\x7a\xb0\x81\xbe\x27\x74\x81\xc8\x0c\x17\x05\x84\x4d\xe5\x92\x9e\xa0\xd9\x57\xf4\xc0\x0c\x25\xa2\x90\x96\x9c\x83\x2e\xce\xb3\x3c\x4e\x0b\x3d\x6f\x56\x4a\xb7\x4d\xa0\x19\xc0\xb7\xae\x79\x53\x45\xdf\x75\x2e\xae\x66\x6b\x4c\x6c\xf7\x4f\x12\x94\x82\xe7\xfe\xa5\xa5\x8b\xab\x56\x08\x63\xe3\xcb\x92\xa0\x55\x1c\x1e\xe0\x87\x8c\xe6\x45\x6b\x72\x31\x19\x8e\x2e\xc6\x43\x55\x00\xbe\x7b\xf1\x57\x1c\xae\x69\xeb\xf0\xc5\x27\xc3\xf9\x38\x9c\xf2\x5a\xf0\xf3\x61\xeb\x7f\xad\xb0\x2c\x5a\xed\xe5\x51\xab\x1d\xbd\x6c\xbf\xfc\x9b\x36\xc2\x4b\xc6\xd9\x87\xd3\xc7\x29\x1f\x2b\xe5\xe6\xbb\x0b\xab\xf6\x82\x7d\xbb\x05\x11\x88\x41\xa9\xbd\x46\x79\xd4\x92\xb5\x29\x34\x54\x9d\xf7\x1d\x88\xac\x56\x0e\xd6\xca\x39\x3a\x66\xfc\xee\x56\xe3\xce\x94\xda\x0b\xf5\x09\x42\x56\x2e\x20\x55\x5b\xfb\xa7\x54\x1f\x8a\x84\x4c\x31\x1f\x67\x75\x48\xb5\x06\x9a\x11\x17\x12\x13\xfd\x65\xa2\xd3\x91\x81\xf0\xfc\x7b\x15\x9b\xb9\x30\x27\x53\x40\x56\x79\x6a\x05\x0f\x0d\x6d\xee\x6b\x57\x5a\xab\x8c\x99\x66\xcc\x65\xda\xd4\x7c\x8a\xb8\xb6\xb8\x75\x25\x48\x4c\x2d\x1d\x79\xc8\xaa\x55\x38\xaf\xae\x01\x41\x83\x99\xa5\x23\xbe\x2d\x2d\xc9\xb7\xf5\x10\xb9\x47\x1b\x36\x8a\x53\x15\xbc\xce\x46\xd5\xcc\xb4\xd5\xf9\xb5\xe6\x29\xd1\x38\xa0\xbf\xdc\x06\xa0\x08\xb8\x8c\x61\xd3\x4f\x75\xbe\xeb\x90\xaa\x79\xed\x3c\x61\x62\xdd\x20\xe4\xca\x0d\xcb\xb9\xd3\xe2\x8e\xe3\xb9\xd3\x34\xca\x69\x1c\xc1\xb8\x8b\x60\x82\x78\x0d\x97\x1a\xa7\xa3\xf9\x4a\x9f\x98\x84\xb1\x95\xf3\x09\x8d\x9b\x1b\xc2\x75\xb8\x62\xdd\x15\x95\x4c\x9b\x8b\xf0\xd3\x89\xca\x13\xd7\x18\x91\x62\xbd\xb1\xce\x50\xd3\x0f\x3e\xc9\xec\x83\xa3\x98\x17\x5f\x30\x81\xd6\xde\x28\x69\x66\xd3\x11\x0d\x0a\x9f\xfa\xb7\xdd\xe8\x7f\x60\xbb\xa9\x58\x06\x3e\x65\x0d\x35\x54\xf7\xf9\x55\x8b\x3c\xf7\x2a\xd3\xab\xf4\x1c\xea\x96\x2a\xeb\xf2\xa8\xa6\xc7\x0c\x4f\x72\x7c\x17\xd3\x92\x09\x27\x39\x65\x53\x5d\x4a\x59\x25\xe1\x56\x11\xcb\xcc\xac\xe0\x59\xdd\x0e\x23\x5b\x04\x2c\x13\x03\x4c\x0a\xf4\x11\x91\xb2\x2a\x88\x3d\x6a\x6d\x56\x99\xb5\x4c\x88\x6c\x58\xcc\xda\x21\x14\xbc\x80\x18\x6f\xad\x85\x2a\x21\x5b\xe4\x8d\xcf\x69\x81\x08\x38\x7a\xa7\x99\xf3\x21\x5e\xad\xb7\xb0\xc6\x34\x4f\xd0\x36\xbd\x11\xbd\xaf\x38\x4b\x31\xc1\xf5\x96\x2c\x3a\x6f\x47\xa6\x0a\x96\xc3\x9e\xbb\x7c\x99\x3d\x78\x3a\x0d\xb6\x64\xc7\x76\x59\x93\x12\xbe\x3d\xa5\xd7\x60\xd1\x19\x4d\xe5\xab\xd5\xe4\xc3\x67\x49\xa2\x01\x64\x7c\x61\x71\xe9\x1e\xfd\x34\xb1\xaa\xfd\x3e\xe5\xa7\xd6\x59\x59\xe0\x31\xbe\x57\x90\xe2\x11\x9d\xdc\x60\x8a\x19\x25\x77\x78\xaa\x8a\x7e\x88\xe8\xb5\x9b\x1b\x34\x0a\x98\x68\x92\x00\x50\x9b\x99\x0a\x18\x20\xac\x94\x8c\x42\x9d\x27\x63\xc4\x0e\x2c\x82\xbc\x25\x5a\x94\x2b\xa6\xdb\xdb\x92\xa7\x2e\xb5\x18\x97\x3f\x83\x3f\xea\xf8\xa3\x45\x4e\xde\x04\xdf\x1d\xfd\x21\xe0\x73\x7e\x31\x1e\x9c\x5d\xbf\x9f\x3d\x03\xcf\x17\x02\xcf\x2e\xe8\xd9\x09\x3e\x7b\xe0\x67\x17\x00\x6d\x87\xa0\x3a\x08\x6d\x61\xd5\x5b\xfb\x6a\xe0\x41\xab\x3d\x05\x1d\xea\xa1\xea\x21\x03\x3f\xdc\x63\x53\x94\xde\x9e\x86\x60\x09\xea\x1a\xe3\xdc\x3e\xd9\xab\xeb\x94\xc1\x6a\xbe\xd4\x19\xa1\x4b\xaf\x15\x32\xdb\x41\x61\x07\xe0\x4c\x92\xa8\x96\xf4\x78\x76\x33\x2e\x92\x44\x75\x88\x01\xc6\xc9\xeb\xe0\xcd\xab\x3f\x04\x5d\x26\x97\x83\x67\x60\x79\x06\x96\x3f\x01\xb0\xb8\x21\x59\x4b\x37\xf6\x44\xe7\x8e\x58\xe7\xaf\x40\x8d\xdc\xbe\xa8\x97\x75\x3b\x57\x69\x48\x2e\x04\x13\x32\x8b\x37\x6f\xfe\x90\xd8\x9f\x9f\xce\x7e\x7e\x4e\x2b\x9e\xa3\xff\x6b\x8b\xfe\x47\x57\x1d\x6b\x5e\xa6\x9d\x5f\xfc\xf3\x72\x08\xf1\x6d\x97\x6c\xa9\xac\xb7\xe6\x57\x83\xab\x6e\x27\x75\x8b\x2f\x62\x8f\x5d\xde\xd8\xf6\x11\xd3\x4f\x41\x2c\x82\x96\x41\xe0\x6f\x2b\xfc\x90\x99\xf3\x1e\xfd\xdd\x84\x42\xdb\xc0\xa2\x9e\xe4\xb8\x59\x4b\x83\xf2\xd0\x11\x10\x36\x3c\x9d\x66\x68\x6b\x82\xa0\x1d\xf0\xf6\x31\xa6\x44\x5e\xda\x07\xfd\x35\x0e\x6f\x59\xb1\x21\x78\x1f\xcc\xdd\x55\x4a\x95\x7e\x1d\xf0\x2a\xb1\x93\x57\xc1\xdb\xe0\xe8\xc8\xbe\x37\xf4\x1e\x84\xb0\x32\xcb\xc0\x0f\xc5\x35\xa6\x3e\x8f\x67\x34\x87\x44\xaf\xcd\xb0\xb9\x95\x96\x01\xbc\xc9\xfc\xd7\x27\x72\x4a\xd3\x22\xdf\xd8\x14\x6e\xb6\xc8\xf9\x01\x5a\x68\x86\xc6\x1f\xe7\xe7\xce\x99\xda\x9e\x01\xce\x4d\x6b\xae\x8a\xea\x88\x63\x5a\x10\x7c\xa9\x24\x4e\x7b\x47\xaf\xf8\x3f\xdd\x0e\xff\xbb\xc6\x47\x0f\x92\x7f\x04\x7c\xf8\xdb\xe7\x6b\x98\x30\x46\x0c\xa1\x26\xc9\x8c\x6f\x58\xf8\xeb\x3a\x9e\x11\xb6\xdd\x76\x9f\x13\xd4\xdc\x57\xd9\xf0\x7d\x6f\xef\xb4\x81\x0f\xbb\x8b\x24\xe7\xac\xb6\x9a\x5d\x22\xb6\xb5\x23\x3e\x5c\x62\xef\x70\x7c\x49\x69\x99\x87\x02\x0a\xfc\x94\x5f\xa0\x5c\xf9\xa0\x7f\x8e\x00\x12\x36\x4b\xff\x1a\xa7\x17\xe9\x17\x11\x86\x62\xb0\x30\x74\xde\x09\x3d\xc9\xeb\x77\xa5\x10\x98\x15\xb5\x83\x79\xfb\xa0\xff\xd7\x32\x8d\x8b\xe0\xa7\x6b\xf8\x2f\x20\x1e\xf4\xed\x34\x0f\xd7\xfc\x95\x9c\x09\x25\x21\x01\x3b\x93\x55\x36\x14\xc2\x2a\x97\x96\xa7\xd9\x6e\x2b\x30\xb1\x16\x5f\xeb\xf0\xc3\xe7\x11\x4d\x57\xb3\x22\x8a\xa9\xfd\xb3\x9b\x8a\x68\x9b\x1f\xa0\x02\x99\x01\x5a\x17\x67\x32\x6b\x99\x85\x88\xe0\x73\x24\x2e\x9a\xa0\x67\x30\x5f\x35\xb2\x3b\x91\x7b\x06\xbb\x63\x06\xe5\xa3\x9a\x6d\x70\x84\xf9\xb3\x11\xfc\x50\x00\x66\x16\xfc\xc7\x43\x91\x78\x47\x52\xc7\x23\x21\xd7\x06\xc1\x93\xe3\xe0\xf5\x0f\xf6\x2d\x47\x18\x67\x31\x78\xe6\x28\x66\xf2\x1e\xf2\xea\x3e\xc5\xb9\xb0\xa2\x6f\x06\x6c\x09\x17\xc4\xca\x1c\xeb\x5b\x55\xdb\x83\xb7\xf5\xd1\xe0\xb4\x54\x11\xa7\x87\x60\xc2\x7b\x95\x20\x3d\x96\xeb\x78\x21\xe7\xf4\xa4\xe3\xa3\x5d\xb9\x10\x37\xa5\x2f\x26\xd3\xab\x9f\x86\xfd\xf9\xcd\x60\x78\x7e\x7a\x3d\x9a\xdf\xcc\xae\xcf\xf8\x37\x80\x9f\x92\x70\xf5\x16\x34\xda\xd4\x95\xfa\x57\xe3\xf9\x70\x0c\x4a\x82\xbd\xa5\x1b\x10\x6c\x77\x71\xe4\x0c\xfd\x91\x13\x60\x4c\xb0\x80\x8f\x65\xea\x5b\xf4\x07\xd7\xd9\xd7\x66\x17\xe0\x00\x85\xeb\x84\x9b\xac\x43\x83\x25\x21\xf2\xc8\x11\x35\xaf\x14\x3c\xaa\xab\xc2\x6f\xdc\xf9\xc6\xe4\x2b\xd5\xe8\xfe\x04\x65\x64\x33\xa7\xf5\x59\x9d\x0e\x27\xa3\x7f\xcd\xaf\xc4\x65\x8c\x90\xf0\x9b\x4b\xc1\x87\x0b\x0e\x9f\x3d\x75\xf1\x2d\x60\xc9\x10\x5d\x34\x75\x3d\xa4\x06\x59\x4f\xf4\x3a\x0d\x80\x4d\x2e\xed\xf4\xac\x00\x9b\xe2\xe2\xa8\xb1\x6f\xba\xbe\x98\x69\x77\xac\xb9\xa1\x27\xa0\xf4\x20\xf2\x44\xb4\x43\x18\x8b\xd8\x0b\xfa\xd2\xb6\x75\x85\xa5\x14\x15\x43\xc4\x6a\x83\xb4\x69\x40\xcb\x55\x0d\xd8\xef\xc7\xeb\xf6\x5e\x7c\x9a\xf5\xa7\x17\x93\xf9\xb7\xad\x42\xbd\x1e\x7e\x77\xb8\xca\x29\xbd\xdb\xb4\xf9\x68\x03\x4d\x3d\xdc\xd9\x84\x73\xe9\xb6\xc3\x29\xbb\xdc\x81\x40\x7a\x26\x7e\xf8\x59\x4d\xd3\x64\x3a\x9c\x0d\xc7\x83\x1b\xd9\x17\xfe\x00\xc2\x16\x3b\xf8\x0d\xbe\xfc\x64\x2f\xae\xfc\x77\x9f\xdf\x76\x19\xba\xe3\x6f\x2e\x21\xc1\x37\x35\x42\x45\x31\xf3\xbc\xe3\x61\xd2\x2e\xbf\x6d\x46\xf4\x1d\x3b\x85\xf5\x22\xa4\x61\xa3\xd5\x5c\x77\xcb\x69\x55\xbf\x66\x65\xf5\x67\x25\xc8\x7d\x0a\x01\xc2\x34\x25\x9b\x8b\x65\xed\x85\x47\x8d\x6c\x6a\x0e\xc4\x37\xd3\x2b\x60\x2b\x49\x8b\x60\x3a\xc0\xef\x7b\x87\x49\x56\x6c\x54\xdf\xaa\xfb\x65\x9f\xd1\xb4\xa9\x36\x0f\xac\x79\x2f\x85\x79\x95\xbf\xa9\x70\x7e\xce\x82\xca\x40\xfc\xdc\x93\x27\x68\x65\xc2\x36\x0c\x7c\x9d\xb9\x3f\xfb\x0c\xc4\x2b\x04\xf3\x06\x5e\x63\x89\xf0\x1d\xfd\x20\x5e\xc1\x86\xd9\x73\x85\x4a\xed\xad\xff\x51\xf0\x26\x78\xed\x57\x03\x3b\x76\x4f\xf5\x80\x7d\x8a\xef\x62\x5e\x21\xc8\x16\x51\x8e\x12\xa8\x43\x73\xe6\xbd\x69\xb7\x1f\x28\x6a\xe7\x5f\x40\xe3\xff\x28\x71\x89\x23\xf9\x32\x97\x59\x01\x50\xe7\x6d\x07\xd4\xdd\x1d\xf1\x71\xd4\x2e\x8d\xf5\x2b\x3a\xeb\x89\x93\x74\xb6\xda\x7a\xf0\xda\xb7\x51\x5a\x2d\xfc\xef\xbb\x52\x7b\x3c\xc5\xf1\x11\xff\x14\x4d\x0f\x3c\xfe\x2f\x8e\x9c\x36\x7c\x2f\x70\x24\xdb\x4a\x94\x5f\xf0\xff\xf0\x04\x27\xd8\xdb\x5c\x43\xbd\x51\x33\xfa\xbb\x7b\xd7\x17\x7a\xd8\x17\x78\x59\xa3\xa7\x09\x62\x66\x3f\x68\xb4\x56\xd1\xbc\xa5\x63\xfe\x24\x45\x31\xaf\x66\xaa\xd7\x8b\x15\xc9\x16\x54\x93\xfe\x4b\x5c\xac\xc7\xd4\xea\x90\x1a\xe9\x36\xf6\xf6\x70\x7a\xe2\x92\x7a\xa3\x7d\xba\xb9\xba\xb3\xc3\x94\x58\xaf\xa9\xac\xc7\x86\xff\x0f\x00\x00\xff\xff\x2a\xe9\x35\xf5\xd1\x41\x00\x00")

func templatesJenkinsNormalJobXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsNormalJobXml,
		"templates/jenkins/normal-job.xml",
	)
}

func templatesJenkinsNormalJobXml() (*asset, error) {
	bytes, err := templatesJenkinsNormalJobXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/normal-job.xml", size: 16849, mode: os.FileMode(420), modTime: time.Unix(1444898786, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJenkinsPipelineXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x93\xcb\x6f\x13\x31\x10\xc6\xef\xf9\x2b\xac\x15\xd7\x78\x53\x0e\xa8\x07\xc7\x05\x05\x2a\x81\x0a\x14\x51\xb8\x22\x77\x77\xe2\x18\xbc\x63\xcb\x8f\x24\x55\xd5\xff\x9d\xd9\x87\x93\xb6\x3c\xa4\xaa\xa7\xc4\xdf\xf7\xf3\xe7\x19\x8f\x57\x9c\xed\x3b\xcb\xb6\x10\xa2\x71\xb8\xac\x4e\xf8\xa2\x62\x80\x8d\x6b\x0d\xea\x65\xf5\xed\xea\x7c\x7e\x5a\x9d\xc9\x99\x88\xc0\x5b\xa3\xae\x9d\xe5\x3f\x01\x7f\x19\x8c\xdc\x1b\x0f\xd6\x20\xf0\xb7\xf4\x43\x01\x37\x97\x93\xf0\xdd\xc0\x8e\x79\x9b\xb5\xa1\xc0\x76\x32\xe7\x05\x9f\x8f\xce\xeb\x05\x3f\xe5\xaf\x2a\x39\x63\x4c\xa0\xea\x40\xde\xde\x32\xfe\x89\xfe\xb0\xbb\x3b\x51\x0f\x4a\x6f\xad\x8d\x4d\x10\xde\xed\xa1\xc9\xc9\x85\x28\xd7\xca\x46\x10\xf5\x63\xf9\x88\x7e\xc9\x90\xe1\x21\x36\x4a\x3d\xe2\x83\xf3\x10\x92\x81\xc8\x1a\xab\x62\x5c\x56\x9b\xdc\x46\x87\xbc\x73\x54\x27\xef\x0b\x7f\x71\x39\x32\x37\x17\x26\xa6\xaa\x1e\xb6\x35\xae\xf3\x0e\x01\xd3\x57\x0f\xcd\x70\x18\x89\x4f\xbc\x90\x1f\xf3\xd5\xfd\x94\x31\xe4\xbf\xbd\x0f\xf6\xda\x84\x98\x3e\xb8\xeb\x01\x39\x9f\x16\x03\x76\x70\x0a\x4a\x0d\x1d\xc8\x0b\x75\x04\x8b\x3e\x96\x5d\x3f\xbf\x6e\x51\xff\x79\x21\x02\xdd\xe7\x75\xd9\x17\xe5\xc9\x82\xfa\x78\xa0\xf4\x4c\xdc\xb8\xdd\x1b\xad\x03\x68\x95\xa0\x2d\x5e\x19\xd6\x3f\xdc\x12\xbe\x72\x36\x77\x48\xd1\x63\x72\x59\x0e\xb9\x8e\x66\x8a\x5a\x22\xd5\x44\x39\xd3\xea\x70\xe2\x56\x25\x75\x7c\x3a\xf7\xa5\x1e\xc9\xbe\xa5\xf3\xde\x23\x3d\x94\xad\xb2\xf2\xa5\xa8\x1f\x29\x25\x67\xb5\x51\xa8\xa9\x93\x14\xf2\x14\x53\x94\x9e\x50\xd6\xba\xdd\x47\x85\x59\xd9\xab\x60\xb4\xa6\x0f\x6a\x22\xff\xe6\xf4\x3b\xa8\x51\xd8\xfb\x32\xd3\x48\x4f\xed\xc9\xc3\x91\xb3\xdf\x01\x00\x00\xff\xff\x30\xbf\x6e\xf1\xbf\x03\x00\x00")

func templatesJenkinsPipelineXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsPipelineXml,
		"templates/jenkins/pipeline.xml",
	)
}

func templatesJenkinsPipelineXml() (*asset, error) {
	bytes, err := templatesJenkinsPipelineXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/pipeline.xml", size: 959, mode: os.FileMode(420), modTime: time.Unix(1444996430, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/jenkins/multi-job.xml":  templatesJenkinsMultiJobXml,
	"templates/jenkins/normal-job.xml": templatesJenkinsNormalJobXml,
	"templates/jenkins/pipeline.xml":   templatesJenkinsPipelineXml,
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
func AssetDir(name string) ([]string, error) {
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"jenkins": &bintree{nil, map[string]*bintree{
			"multi-job.xml":  &bintree{templatesJenkinsMultiJobXml, map[string]*bintree{}},
			"normal-job.xml": &bintree{templatesJenkinsNormalJobXml, map[string]*bintree{}},
			"pipeline.xml":   &bintree{templatesJenkinsPipelineXml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
