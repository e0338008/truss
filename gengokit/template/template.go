// Code generated by go-bindata.
// sources:
// NAME-service/NAME-client/client_main.gotemplate
// NAME-service/NAME-server/server_main.gotemplate
// NAME-service/generated/client/grpc/client.gotemplate
// NAME-service/generated/client/http/client.gotemplate
// NAME-service/generated/endpoints.gotemplate
// NAME-service/generated/transport_grpc.gotemplate
// NAME-service/generated/transport_http.gotemplate
// NAME-service/handlers/client/client_handler.gotemplate
// NAME-service/handlers/server/server_handler.gotemplate
// DO NOT EDIT!

package template

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

var _nameServiceNameClientClient_mainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdf\x6f\xdb\x36\x10\x7e\x16\xff\x8a\xab\x90\x02\x72\xe1\xd2\x1d\xb0\xbd\x18\xf0\x43\x9a\xb5\x6b\x80\xb5\x30\x92\x00\x7d\x29\x30\xd0\xd4\x49\xe2\x22\x91\x2a\x49\x2b\x09\x08\xfd\xef\xc3\x51\xb2\x6c\x37\x36\xb2\x5f\x01\x12\x89\xc7\xbb\xef\x8e\x77\xdf\x47\xa5\x15\xf2\x5e\x94\x08\x8d\x50\x9a\x31\xd5\xb4\xc6\x7a\xc8\x58\x92\xa2\x96\x26\x57\xba\x5c\xfc\xe9\x8c\x4e\x59\x92\x16\xb5\x28\xe3\xb3\xf1\xf4\x30\x8e\xfe\x3a\x6f\xa5\xd1\xdd\xf8\xaa\x74\x19\xad\x5e\x35\x98\x32\x96\xa4\xa5\xa9\x85\x2e\xb9\xb1\xe5\xe2\x71\xa1\xd1\x2f\xa4\xd1\x1e\x1f\x23\x40\x69\x4c\x59\x23\x3f\x70\x29\x6d\x2b\x87\x30\xe5\xab\xed\x86\x4b\xd3\x2c\xda\xfb\x72\x81\xd6\x1a\xeb\x68\x67\xb1\x80\xbb\x4a\x39\xb8\x45\xdb\x29\x89\x2c\x91\xb5\x42\xed\x3f\x09\x9d\xd7\x68\x21\x0d\x81\x5f\xc7\x23\xac\x85\xaf\xe0\x6d\xdf\xc3\xa2\x1a\xf6\xdc\x62\x70\x4d\x59\x42\x69\x86\xc5\xc9\x80\x12\x35\x5a\xe1\x31\x1f\x23\xc6\xb2\x92\xca\xfb\xf6\x9f\x84\x91\x7f\xca\x92\x76\x13\xdd\xd7\xef\x8f\x03\x52\x36\x63\xac\x13\x96\x7a\xfd\x07\xac\x60\x6c\x24\x5f\x0b\xeb\xf0\x5a\xfb\xc9\x4a\x3d\xe5\xb7\x6d\xad\x46\x13\x8d\x83\x5f\x99\xa6\x15\x72\xb4\x0c\xed\xe1\x5f\xad\x68\x8b\xc1\xd2\x6e\xf8\x0d\x96\xca\x79\xb4\x21\xf0\xb1\x59\xfc\x8b\x68\xb0\xef\x69\x85\x96\xb2\xb3\x62\xab\x65\x1c\x7c\x36\x83\x30\x36\x17\x41\xe4\xb9\xac\x15\xb4\x16\xdd\xb6\x41\x07\xda\x80\x1b\x10\x20\x57\x4e\x9a\x0e\xed\x13\xb8\x27\xe7\xb1\x99\x83\xd0\x39\xe0\x63\x8b\xd2\x3b\xd8\x3a\xb4\x0e\xbc\x89\x48\xad\x35\x9d\xca\x11\x7c\x45\x61\x16\xa5\x27\x60\x8b\xce\x81\x29\x40\xe8\x1d\x26\x1f\x06\x3a\x64\x6b\xbd\x32\x1a\x94\x03\x8b\x45\x8d\xd2\x63\x0e\x4a\x47\x38\x82\xa1\xaa\x36\x4a\x0b\xfb\x14\xd3\x92\x69\x34\xd3\x4c\x46\x22\xbb\x65\x34\xbe\xf5\x56\x68\x47\x0d\xe7\x94\x16\x88\xbc\x2e\x22\x51\x68\x27\xac\x32\x5b\xb7\x0b\x95\x46\x3b\x6f\xb7\xd2\x1b\xeb\x60\x63\x7c\x35\x1e\x09\x2a\xe3\xfc\x32\x2a\x62\x37\x08\xc6\x92\x71\x68\x91\x0f\x97\x84\x3d\xfc\xac\x62\x0e\x7e\x1b\x1d\xb3\x94\x76\x63\xea\x74\x0e\x29\xfd\x7e\xba\xbb\x5b\x1f\xb5\x20\xcf\x5d\x27\xd3\x19\x4b\x22\x21\xcf\x23\xd1\xee\x84\xb4\xfc\xe5\xdd\xcf\xef\xe8\xa5\xbc\x59\x5f\x41\x46\xa0\xb3\x33\xa8\x0d\xfa\xca\xe4\x00\xa7\x51\x87\x5d\x42\x0a\xc1\x0a\x5d\x22\x5c\x28\x9d\xe3\xe3\x1c\x2e\x14\x2c\x57\x30\xb1\xe6\x73\x74\x74\x7d\x1f\x82\x2a\x46\x27\x5a\x60\xed\x90\x9e\x77\xe6\x77\xf3\x80\x16\x2e\xd4\x48\xb0\x10\x50\xe7\xd3\x23\x9d\x27\xff\x2e\xc3\x7c\xc2\x39\x97\x81\x4e\x39\x3b\x18\x48\x08\xfc\x2a\x0e\xf4\xd2\x96\x8e\x5f\xd6\xf5\x47\x1a\x7a\xdf\x93\x57\x12\x0f\x1f\xc5\x95\x11\xf5\xa7\xa0\x1d\xb5\xdb\x0d\x0f\xe1\x37\x43\x09\xe0\xb4\x62\x92\x04\xed\x30\xa2\x28\xb8\x88\xaa\x0a\x78\x33\xf1\xe0\xd5\x0a\xd2\x94\x84\xb4\x03\x9d\x93\x27\xac\x60\x7f\x73\xf0\x2f\xf8\x90\x4d\x11\x33\x96\xf4\x40\x7d\x04\xc2\x99\x58\xb0\xc7\x91\x46\xeb\x01\x64\xb9\x82\xc8\x83\x5f\x95\xa8\xb3\xc9\x75\x3e\x18\xbf\x2a\x5f\x5d\x6b\x87\x72\x6b\x31\x9b\x1d\x18\xef\x54\x83\x66\xeb\x33\xba\x8f\xf9\x2d\x4a\xa3\xf3\x19\x51\x43\x15\x11\xf4\xd5\x0a\xb4\xaa\x63\xa6\xa4\x68\x3c\xff\xd8\x5a\xa5\x7d\x91\x19\xc7\x6f\x7d\x8e\xd6\xce\x21\xfd\x40\x67\x85\x87\x4a\xd5\xa4\x63\x51\x2b\x5d\x46\x7c\x12\x8d\x46\x49\x82\x5d\xc2\xeb\x2e\x8d\x65\x12\x76\x62\x1c\xff\xf0\xa8\x7c\xf6\x13\xad\x7a\x96\x24\x39\x16\x68\xa3\x3f\xbf\xaa\x4d\x1c\xc0\xb3\x16\xed\xef\xe4\xd8\x22\x72\xde\x77\x87\x0a\x3c\x57\x5f\x9c\xc5\x92\x2e\x29\x8b\x8d\xf1\x38\x89\xc1\xb5\x28\x55\xa1\x30\xff\xa6\xa3\x1c\x0e\xcb\xea\xd9\x89\x16\xbc\x90\xe1\x75\xf7\x4d\xef\x4f\x79\x8c\xc6\x12\xf7\xa0\xbc\xac\xe0\xcd\x28\xba\xc0\x42\x80\x07\xe5\x2b\xb8\xf0\x18\xc9\x4e\x34\xdc\xeb\x80\x4c\x17\x1e\x9f\x4b\x80\x86\x2e\x1c\x92\x2a\x9f\xd1\x3e\x5d\x52\x7b\x43\x78\x3b\x20\x47\x9d\x44\x94\x03\xd6\x13\x10\xbd\xec\xa2\xe8\x53\x43\x51\x91\xef\x74\xe0\x91\xbb\x11\x88\x7f\x16\xd6\x55\x62\xaf\x93\x01\x3d\x8a\x8b\x16\x16\xbf\x6f\xd1\xf9\x89\x81\x47\x1f\x5b\x1e\xc2\x54\x59\x16\xc2\xdf\x2f\x89\xf4\xcb\xaf\x44\x5d\x93\x71\x12\x73\xa4\xce\x09\x5e\xbe\x40\x4c\x29\xea\x48\xc9\xb3\xa5\xfd\x38\xb8\x1f\xf8\x19\x87\x97\x24\xdd\x74\xc6\xdd\x67\xe9\xf0\x74\xe3\x3f\x2d\xfc\xbd\x90\xf7\xa5\x35\x5b\x9d\x93\xca\xc6\xee\xfc\xc7\xc2\x4f\xe4\x7b\xb9\xe4\x11\x7d\x4d\xe0\xb5\xce\xd2\xa1\xdb\x70\x33\x54\x84\x79\x24\xc8\x32\xd2\xfe\xc8\xf1\x7f\x1a\xd3\x51\xee\xe1\x76\x84\x1b\x74\xad\xd1\xf9\xf9\xdc\xdd\x8c\x1d\xb0\xeb\x80\x67\x39\x16\x62\x5b\xfb\xe5\xcb\x12\x54\xba\x13\xb5\xca\x61\x14\xd9\xeb\xef\xb1\x4b\xc3\xea\xb9\x28\x7b\xc6\xfe\x0a\x00\x00\xff\xff\x23\x40\xd2\xd4\xdd\x0a\x00\x00")

func nameServiceNameClientClient_mainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceNameClientClient_mainGotemplate,
		"NAME-service/NAME-client/client_main.gotemplate",
	)
}

func nameServiceNameClientClient_mainGotemplate() (*asset, error) {
	bytes, err := nameServiceNameClientClient_mainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/NAME-client/client_main.gotemplate", size: 2781, mode: os.FileMode(436), modTime: time.Unix(1478036043, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceNameServerServer_mainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdf\x6f\xdb\x36\x10\x7e\x16\xff\x0a\x56\xe8\x00\x69\x50\xa8\x0e\x5b\xf7\x10\x34\x0f\xad\xdb\xa5\xc1\x92\xc0\x88\x3d\xec\x99\x96\xce\x34\x11\x89\x14\x48\xca\x75\x20\xf8\x7f\x1f\x8e\xa4\x1c\xb9\x75\x9c\x3d\xb4\x40\x2d\x89\xf7\xf1\xe3\x7d\xc7\xfb\x91\x8e\x57\x8f\x5c\x00\x6d\xb9\x54\x84\xc8\xb6\xd3\xc6\xd1\x8c\x24\xe9\xba\xe1\x22\xc5\x67\xeb\xf0\xa1\x60\x7c\x94\x1b\xe7\xba\xe9\x7b\xd9\x75\x46\xaf\x71\x45\xdb\xf0\x5b\x5a\x29\x14\x6f\xf0\xc3\x3e\xd9\x8a\x37\x4d\x4a\x48\x52\x96\xf4\xf7\x9a\xce\xb9\x71\x4f\x24\x49\x85\x6e\xb8\x12\x4c\x1b\x51\xee\x4a\xa4\xaa\xb4\x72\xb0\xf3\xa7\x08\xad\x45\x03\x6c\x02\x11\xa6\xab\x22\xc7\xb5\xa6\x7f\x4b\x87\x28\xe9\x36\xfd\x8a\x55\xba\x2d\x85\xbe\x78\x94\xae\xc4\xff\xa0\xea\x4e\x4b\x15\x78\x4e\x22\x1a\x2d\x22\xd5\x72\x23\x2d\x5d\x80\xd9\xca\x0a\x48\xb2\xe1\xaa\x6e\xc0\xd0\x74\x18\xd8\x8d\x0f\xc4\x9c\xbb\x0d\xbd\xd8\xef\x69\x19\x6d\xb6\xb4\x60\xb6\x60\x52\x92\xd8\x6d\x75\x12\x29\x40\x81\xe1\x0e\xea\x94\x24\xdd\xca\x43\xe6\x9f\x8e\x41\x29\x21\x39\x21\xeb\x5e\x55\x3e\xec\x59\x4e\x07\x92\x6c\xb9\xc1\xb8\x27\x35\xac\x7a\xf1\xb1\xae\x0d\xf5\xff\xae\x28\x5e\x04\x5b\x38\x23\x95\xc8\x52\x6f\x65\xbc\xae\x4d\x5a\xd0\xf4\xf2\xfd\xbb\x3f\xdf\xe1\xcb\x67\x5c\xa6\x5c\xd5\xb4\x05\x67\x64\x65\x69\x23\xad\x03\x45\x11\x09\xd6\xa6\x39\x49\x12\xbc\xab\x67\xe2\xef\x99\xd1\x3a\x25\x7e\xef\x89\xbf\x2e\x97\xf3\x53\x5c\x78\x1f\x2f\x73\xa1\x75\xca\xf5\x87\xe7\x12\x0f\xf3\x19\xcd\x90\x31\x3f\x41\x99\x93\xc4\x73\xcc\xb9\xb1\x90\xe5\xe1\x82\x6e\xb5\x10\x52\x09\x5a\x6b\x0c\x14\x0b\x51\x6a\xb4\x10\xe0\x1f\xec\xd6\xbf\x92\x64\x20\x49\x12\x97\xaf\xbc\xe1\x1e\xbe\xdd\x6a\xb1\x6e\x5d\x40\x64\xda\xb2\x85\xab\x75\xef\xf2\x13\xc8\x59\xc8\xbc\x2c\xac\xe7\xec\x5f\xe9\x36\x59\xea\x6c\x5a\x78\xc4\x67\x58\xf3\xbe\x71\x4b\xd9\x82\x75\xbc\xed\xfe\x59\xce\xfe\x3f\x0b\x26\x3f\x98\x63\xa6\x99\x5f\xcb\x49\xb2\x27\x91\x05\x85\x64\x69\x6b\x05\xc6\x69\x03\x4d\xa3\x31\x24\x35\xac\x61\x94\x7b\x84\x10\x5a\xd7\xab\x27\x48\x63\x94\x3e\xf5\x56\x2a\xb0\xf6\x38\x4c\x36\xe4\x35\xed\x56\x6c\x18\xae\xf5\x3d\x6f\x81\xb2\x98\xec\x0c\xbf\xf6\xfb\x85\x4f\xe6\x10\xbe\x11\x7e\x45\x63\xae\xa3\xa6\x08\xcf\x50\x6f\x59\xe2\x75\x1d\x68\x9b\x78\x35\xd3\xa4\xdb\x80\x01\x54\xe5\xbd\xfa\x12\x0b\xf1\xd9\xab\x61\x30\x5c\x09\xa0\x6f\x25\xbd\xbc\x7a\xf6\xe5\x0e\xdc\x46\xd7\x76\xbf\x0f\x7e\x0f\xc3\x52\xdf\xea\x6f\x60\xe8\x5b\x19\xfd\x3c\x50\x8d\xc5\xcd\xc6\x95\xe0\xfb\xb9\x2d\x57\xd4\x6e\x2b\x76\xc7\x1f\x61\x18\x7e\xb0\x66\x51\x4d\xd4\xf7\xb1\xae\x0f\x47\x50\x67\x78\x25\x95\x28\xa8\x54\xd6\x99\xbe\x05\xe5\xb8\x93\x5a\x79\xc5\xa3\xfa\x51\x71\x32\x0c\xa0\x6a\x94\x30\xee\xb7\xa8\x11\x8f\x1e\xcf\xb2\xc3\x6b\x11\xc0\xde\xe0\xd5\xfc\xe0\xe7\x25\x16\xd9\x19\x95\xc5\xc4\x81\x18\xfe\x3b\xa8\x36\x5c\xc9\x8a\x37\xcf\x17\x00\xc6\x54\x78\x70\xcb\x1f\x21\x43\x33\x05\x63\x34\x66\x62\xe5\x76\x68\x88\x3d\x98\x7d\xe2\xd5\xa3\x30\xba\x57\xf5\x58\x89\x37\xca\x81\x31\x7d\xe7\x0e\xe9\x41\x12\xa1\x29\x36\xb1\xd0\xbf\x92\xef\x98\xb1\xe2\xfc\x14\x28\xe8\x6f\x18\xde\x30\x12\xd8\xbd\x76\x72\xfd\x94\x55\x05\x8d\x93\x81\x2d\x6e\xae\x6f\xee\x97\x47\xdf\xcb\x2f\x0f\x77\xb8\xc7\xfb\xfb\xe1\x82\xae\x5b\xc7\xbe\xa0\xa7\xeb\x2c\xfd\x05\xcb\xf2\xc3\x45\x85\xe5\x33\x3a\x17\xda\x5f\xe8\x29\x27\x3c\x8b\x95\x7a\xf9\x5a\xc1\x1b\xae\x2c\xf6\x69\x2c\x31\xdf\x68\x7d\x81\x25\x2d\xee\xf4\xed\x31\x56\x04\xdc\xf5\x3b\x5f\x12\x2d\xfb\xea\x83\x91\xa5\xa5\xc7\x87\x49\x58\xa6\x45\x80\x07\xa3\xf9\x0b\x3d\xf1\x16\x76\xa3\x6a\xd8\xe5\x67\xb6\x56\x6d\xdd\x48\x05\x2f\x33\xcc\x02\xe0\x1c\x07\xfe\xc8\xe6\x0c\xc7\x3c\x00\xce\x71\xd8\xa7\x76\xa5\x9b\x97\x29\x16\xde\x7e\x8e\x01\xcb\xe7\x8c\x0f\x4b\x34\xe7\x3e\xbe\xd3\x06\x17\x87\xc6\xaf\x87\x29\x38\x4d\x03\x4f\x75\xeb\x6f\xf9\xa3\xaa\xfd\x4d\x64\xcf\xc8\x82\xb6\xd3\x9c\xf0\x93\xeb\x70\xa5\x3f\x25\x27\x90\x32\x4c\xd2\xb1\xb6\xb1\xad\xe0\x6a\xd4\x97\x55\x6e\x57\x1c\x3a\x88\x2d\x62\xef\xce\x5f\x10\x39\x0e\xe4\x57\x35\x8e\xc0\x82\x6e\xa6\x12\xfd\x40\xfd\xb9\x12\x91\x32\x64\x7d\xa3\x0a\xec\x0e\xb8\x5d\x81\x8b\x2e\x65\xa9\xab\x3a\x74\x7d\x9c\xff\xe8\xba\x5c\x7b\xe0\x9b\x2b\xaa\x64\xe3\x4f\x3e\xa8\x01\x63\xf0\xd3\x80\xeb\x8d\x22\x89\xef\x4d\x89\x35\xdb\x69\xfc\xae\x1f\xe6\xb3\x30\x88\xbe\x0b\x9f\xef\x1a\x88\xf4\x7f\x4e\x8c\xb5\x67\x7c\xe5\x75\x2b\xf6\x00\x02\x7d\x32\xc3\x70\x72\xaa\x65\xb6\xa0\xd6\x6c\x5f\xcc\xb0\xa9\x82\xd1\x5d\xcb\x42\xc0\x1b\x35\x0d\xf3\x43\xaf\xde\x1c\x0f\x6a\xd8\x49\xe7\xbb\x10\x6e\xcc\xc9\x9e\xfc\x17\x00\x00\xff\xff\xf4\x3e\x62\x2e\x46\x0b\x00\x00")

func nameServiceNameServerServer_mainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceNameServerServer_mainGotemplate,
		"NAME-service/NAME-server/server_main.gotemplate",
	)
}

func nameServiceNameServerServer_mainGotemplate() (*asset, error) {
	bytes, err := nameServiceNameServerServer_mainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/NAME-server/server_main.gotemplate", size: 2886, mode: os.FileMode(436), modTime: time.Unix(1478035743, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedClientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x55\xcd\x6e\xe3\x36\x10\x3e\x8b\x4f\x31\x15\x82\x42\x0a\x14\xea\x9e\xc2\x97\xf5\x6e\x17\x5b\x34\xa9\x91\x0d\xda\xc3\x62\x51\xd0\xd4\x58\x26\x6c\x93\x2a\x49\x3b\x36\x04\xbd\x7b\x31\x14\xe9\xc8\xde\x64\x0f\x41\xc4\x99\x6f\xfe\xbe\xf9\x71\x5d\xc3\x42\xc8\x8d\x68\x11\x5a\xdb\x49\xe8\xac\x39\xa8\x06\x1d\x08\x68\x9f\x16\x73\x90\x5b\x85\xda\xc3\xca\x58\xf0\x6b\x84\xbe\xe7\x5f\xd1\x1e\x94\x44\xfe\x28\x76\x38\x0c\xe0\xe2\x93\x75\x13\x37\x8c\xa9\x5d\x67\xac\x87\x82\x65\x79\x6b\xb6\x42\xb7\xdc\xd8\xb6\x3e\xd6\x1a\x7d\x2d\x8d\xf6\x78\xf4\x79\xd0\x99\x76\x8b\x7c\x02\x21\xf3\xf7\x35\xf5\x0e\xbd\x68\x84\x17\x01\xa2\xfc\x7a\xbf\xe4\xd2\xec\xea\x6e\xd3\xd6\x68\xad\xb1\x2e\x67\x97\x9a\xd6\xdc\x6d\x94\xaf\xe9\x0f\x75\xd3\x19\xa5\x29\x30\xf9\xf2\x56\x68\x17\xb2\x7c\x07\x7f\x06\xc4\xa4\x58\x56\xd7\xf0\xbc\x56\x0e\x22\x07\x2c\x73\x07\x09\x79\xdf\xf3\x2f\xa1\xdc\x85\xf0\x6b\xb8\x1b\x06\xa8\x5b\xd4\x68\x85\xc7\x26\x67\x59\xb7\x0c\x90\xc5\x87\x4b\x50\xce\x4a\xc6\xea\x1a\x1e\xf1\x05\x2c\xfa\xbd\xd5\x0e\x84\x4e\x7c\xc2\x52\xc8\x0d\x36\xb0\x3c\x5d\x75\x42\x1a\xad\x51\x7a\x65\x34\x87\x2f\x1e\x94\xa3\xbe\xb0\xd5\x5e\x4b\xf2\x54\x90\x1a\x6e\x29\x5f\x3e\x0f\x06\x73\xa3\x75\x05\xa6\x23\x0b\x07\x9c\x47\xf1\x5f\x41\x50\x42\xd1\x2d\x79\xdf\x7f\x36\xd4\x4d\xb8\xea\x2d\xbd\xd0\x56\x10\x88\x2d\xa1\x67\xd9\x41\x58\x90\x32\xa6\x32\x37\x7a\xa5\x5a\xc6\x32\x1a\x8e\x7f\x2b\x58\xc1\xfd\x0c\xac\xd0\x2d\x9e\xc3\xf5\x2c\xcb\xd0\x5a\x52\xac\x8a\x5f\xa5\x2c\x59\x96\xa9\x15\x39\x84\x5f\x66\xa0\xd5\x36\x20\xb2\xb1\x7c\x7a\xc7\x60\x8e\xff\x63\x45\x57\xa0\xb5\x15\xe4\x52\x68\x6d\x3c\x88\xae\xdb\x9e\xa2\xe7\x9c\x1c\x0d\x2c\x1b\x18\xcb\xe4\xa4\x1e\x47\x91\xbe\x7d\xbf\xe8\xee\x45\xc1\x14\xee\x2d\xed\x07\x5c\x19\x8b\x05\x25\x13\xa7\xf3\x6f\xb1\xdd\xa3\x7b\x36\x9f\x9f\x16\xf3\x87\x38\x74\x85\x94\x7c\x8d\xa2\x41\xeb\xca\xb2\xa2\xf0\x59\xdf\xdf\xc1\x8b\xf2\x6b\xb8\xf1\x48\xc1\xf9\x30\xb0\x6c\x22\xed\x36\x6d\xa0\xf6\x7e\x46\x08\x1e\x77\x6d\xe4\x97\xa2\x11\x72\xe4\xec\x46\x25\x50\xea\xc2\x03\xfa\xb5\x69\xdc\x08\x0c\xdc\xf7\xfd\xb3\xf9\xd3\xbc\xa0\x85\x1b\x15\x9b\xf4\x29\x0e\x35\xa4\xe9\xe6\x49\x12\xac\x02\xbf\x14\xe6\x7d\xc3\x19\x5c\x32\xf2\x88\x2f\x23\x29\xc5\x68\x4b\x8c\xe8\x2a\x7e\xe7\x7d\x9f\x6a\x1a\x06\xde\xf7\xd3\x7c\x47\x61\x3e\x85\xaa\x6b\xa1\x3b\x48\xfe\x49\x4b\xd3\x20\x11\x3b\x41\x3c\xe1\x7f\x7b\x74\x7e\x8a\xfb\x88\x6f\xe2\x5c\x67\xb4\xc3\x04\x9c\xce\xef\x8d\xe2\x49\xfd\x7c\xea\x52\x42\xfd\x90\xb0\x17\xa3\xc2\x39\x8f\xf2\xf2\x4c\x59\x51\x06\x49\xec\x0c\xea\x26\x76\x33\x7e\xa5\x0f\x96\x26\x76\xac\x66\xb4\x75\x3d\x01\xa6\xbd\xbc\x6e\x24\x6d\x7d\x70\xf7\x43\x0f\xee\x01\xe0\x67\xcd\xad\x5e\x63\x67\x43\x45\x8b\xc2\x06\xc6\xfc\xa9\xc3\x8b\x5d\x04\xe7\xed\x5e\x7a\x5a\xaa\x38\xa6\xf0\xed\xbb\xf3\x56\xe9\x96\xf0\x75\x0d\xd3\x5d\xa0\xdb\x21\x80\x2e\x47\x78\xf9\xb5\xf0\xb0\x33\x8d\x5a\x29\x0c\x47\x65\x72\x71\x68\xcf\x43\xb4\x0b\x7b\x32\x2d\x6e\xa7\x09\x94\xe3\xfa\xb2\xf1\x1e\xcd\xfd\x31\x6d\xd1\x57\xd4\x4d\xb1\xc1\x53\xb8\x40\x63\x46\xe5\xa5\xb3\xfe\x4c\x6a\x70\x6b\xe0\x2d\xc7\xe1\x5c\x98\xb4\x83\x30\x03\x72\xc9\xa6\x07\x84\x96\x72\x88\xf1\x7f\xb6\xc9\x21\x97\x44\x4e\x79\xb5\x01\x71\x16\x7f\x27\x27\x57\x79\x49\x7f\x4c\x7e\xf9\x7c\xfc\x5f\xc1\xae\x81\xdb\xf4\xc3\xc4\x1f\x3e\x96\xd7\x88\x90\x36\xed\x6f\x27\xd4\xb4\x27\x59\x3a\x9d\x9b\xd7\xd3\x19\x12\x0b\x5b\xab\x56\x70\xa8\xc0\x04\x9d\xf4\x47\x1e\xea\x28\x36\x25\x2f\x62\xd6\xbf\x91\x72\x5c\xf0\xd1\xf1\x8c\x8e\x24\x31\x1d\x9e\x15\x6c\x2a\x38\x84\x89\x1e\xc2\xb9\x1c\x8f\xef\x08\x9d\x9e\xdf\xdb\x5d\x03\x33\x38\x17\xf0\x87\x51\xba\xb8\xdd\x35\xd5\xab\x68\x41\x36\xa3\x57\xce\x79\x59\x26\x77\x91\x19\xe9\x8f\x23\xef\xff\x07\x00\x00\xff\xff\x3a\x03\xa3\xb4\x4e\x08\x00\x00")

func nameServiceGeneratedClientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedClientGrpcClientGotemplate,
		"NAME-service/generated/client/grpc/client.gotemplate",
	)
}

func nameServiceGeneratedClientGrpcClientGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedClientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/client/grpc/client.gotemplate", size: 2126, mode: os.FileMode(436), modTime: time.Unix(1478035718, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedClientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func nameServiceGeneratedClientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedClientHttpClientGotemplate,
		"NAME-service/generated/client/http/client.gotemplate",
	)
}

func nameServiceGeneratedClientHttpClientGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedClientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/client/http/client.gotemplate", size: 105, mode: os.FileMode(436), modTime: time.Unix(1477950707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\xdd\x6e\xdc\x38\x0f\xbd\x1e\x3f\x05\xbf\x20\x40\x67\x3e\x4c\x34\xf7\x0b\xf4\x66\x8b\xfd\xc9\x45\x8b\x62\xdb\x17\xd0\xc8\xf4\x98\x88\x2c\xa9\x12\x3d\xc9\xac\xe1\x77\x5f\x50\xb2\x1d\x67\x63\x2c\x7a\x11\x38\x43\x53\x87\x87\xe7\x90\x72\xd0\xe6\x49\x5f\x10\xd2\xd5\x54\xd5\xe9\x04\xdf\x5b\x4a\xd0\x90\x45\x30\xde\xb1\x26\x97\xa0\x43\x6e\x7d\x9d\x80\x3d\x74\xfa\x09\x81\x5c\x4d\x57\xaa\x7b\x6d\x01\x5d\x1d\x3c\x39\x4e\xd0\x44\xdf\x41\xc2\x78\x25\x83\xe9\x28\x48\x11\x7f\xf4\x98\x18\xb4\xab\x21\x62\x0a\xde\x25\x04\xbe\x05\xcc\x48\x92\x8a\xc0\xad\x4f\xf8\x8a\x72\x04\x9d\xe0\x19\xad\x95\x27\x3a\xe3\x6b\x8c\x49\x00\x04\xaf\xc6\xe9\x77\xe3\xe3\x74\x30\xa3\x1d\x73\x40\x5b\x0b\xbe\x01\xdf\x47\x48\x7d\x08\x3e\x32\xd6\xc0\x51\xbb\x24\xff\x4b\x39\xd2\x96\xfe\xd6\x4c\xde\x09\x5a\xe3\x63\xa7\x39\x29\x78\x64\xd0\x36\x79\x20\x67\x6c\x5f\x63\x5a\xd8\x40\x47\x75\x6d\xf1\x59\x47\x4c\xaa\xaa\xa8\xcb\x40\xfb\x6a\x77\x77\xf1\x56\xbb\x8b\xf2\xf1\x72\x7a\x39\x39\xe4\x93\x48\x85\x2f\x7c\x57\xc9\x4b\xe2\xb6\x3f\x2b\xe3\xbb\xd3\xc5\x3f\x3c\x11\x9f\xe4\x6f\x06\x95\x94\x70\x86\xbb\x61\x50\x5f\x7f\x7d\xcc\x90\x5f\x35\xb7\xf0\x30\x8e\x77\xd5\x21\x3b\xf0\xdb\xa2\xa9\xf1\xd6\xa2\xe1\x34\x37\xc7\xed\x4a\x2b\xe0\x56\x33\x18\xdf\x05\x51\x42\x3b\xd0\x75\x3d\x1b\x20\x5d\x7d\x48\x02\xd6\xa1\x76\x2c\x7a\x9f\x11\xfa\x84\xb5\x08\xab\xa1\x45\x1b\x30\x42\xe2\xd8\x1b\x3e\xca\xeb\xa9\xd4\x76\x25\x72\xec\x41\x0b\x5c\x22\x77\xb1\x08\x41\x47\xdd\x21\x63\x54\xd5\xe9\x24\xf1\x47\x07\xba\x58\x1a\x8f\x40\xfc\x21\x49\xb1\xa6\xb7\xd9\x9a\xa6\x77\x46\x64\x9f\x28\x3b\x14\x67\x3c\xf8\x80\x51\x33\x82\x97\xb3\x01\xe3\xc3\x5c\x50\x00\xcf\x3a\x51\x52\xf0\xbb\x8f\x80\x2f\xba\x0b\x16\x8f\x70\xf3\x3d\x74\x74\x69\x19\x82\x4e\x32\x16\x2b\xa9\x84\xe0\x52\xa8\xd4\x09\xd1\xd7\xbd\xc1\x2c\x83\x76\xd0\x32\x07\xf5\xa7\x76\xb5\x15\x8e\xcf\xc4\x2d\xa0\x36\xed\x34\xdd\xb0\x9f\xab\x1f\xe0\x99\x22\xd6\xd0\x87\x02\x9a\x02\x1a\x6a\xc8\x40\xd0\xdc\x2a\xd8\x3f\x66\x7e\x94\x04\xff\xac\xcf\xf6\x06\x1a\x3a\x4a\x5c\x36\x03\x6a\x4c\x74\x71\x72\x94\xdc\xd5\x3f\x61\x96\xf2\x5b\xb1\x65\xd9\xa4\x4c\x11\xdf\x9a\x5d\xcc\x10\x88\x59\x49\x75\x58\xab\x6b\x2c\xa1\xe3\xb7\xea\xae\x8c\x7b\x5d\x4a\x7b\x93\xd5\x2d\x70\x58\xff\x97\x8d\xb2\x3e\x45\x2b\x12\x85\x3b\x2c\x63\xf5\xca\x97\x1c\x63\x6c\xb4\x0c\xd4\xb6\x13\x02\xb6\x14\xdb\xbe\x18\x7a\x29\xf6\xba\x89\xa7\xec\xc3\x17\x7c\xfe\x34\xf5\x63\x7c\x77\x26\x97\x75\xea\xb2\xb2\x99\xe5\xca\xdb\xe3\x74\x83\x70\x1f\x1d\x50\x1e\x66\xe1\x68\xb4\xb5\x18\xcb\x3c\x4f\x7c\x55\x95\x3b\x7a\xa7\xe9\x50\x0d\x43\xd4\xee\x82\x70\x4f\xf0\xcb\x47\x50\x73\xfe\xe7\xe2\xc7\x38\x56\xbb\x61\xb8\x27\xf5\x45\x77\x38\x8e\xf3\x79\x00\x58\xfa\x50\x73\xb0\x1a\x86\x07\x89\x8e\x63\x35\xbe\x5d\xd7\x9f\x28\x22\x03\x0a\xfb\x15\xc3\x03\xac\xea\xee\x0d\xbf\xc0\x74\x95\xa8\x4f\xe5\x79\x94\x81\xf8\x7f\x38\xab\x61\xf8\xc3\x4b\x1a\xdc\x93\xfa\xab\xdc\xac\xdf\x6f\x01\xa7\xa3\x07\xd8\xbf\x4f\x2a\x57\xee\x2a\xeb\x08\x18\xa3\x8f\x07\x18\xaa\xdd\x6e\xbe\x92\x73\x50\x08\xa3\xda\xd0\x40\x38\x09\x87\x43\xb5\xdb\x51\x93\x53\xff\xf7\x11\x1c\xd9\x8c\xb1\x9b\x5c\x71\x64\x33\x4c\xb5\xdb\x8d\xd5\x12\x9d\x2b\xa8\x9f\xe1\x76\x38\x0a\x4a\xb5\x1b\xab\x61\x28\xf2\x8a\xb8\x9f\x65\xab\xd6\x0a\xe7\xbd\xbd\x67\xcc\x0a\x17\xdf\xd6\xa2\xdf\x33\x6e\xe9\x5e\x84\x17\xb0\xad\x16\x13\x64\x7a\xeb\xb3\x25\xe3\x5b\x5e\xc3\xc3\xfb\x21\x78\xd3\xbc\x60\x6f\x5b\x37\x7f\x01\x97\x35\x1a\xc4\xa8\xe5\x5b\xb8\x0a\x17\x13\x56\xee\x08\xfa\x0f\xe9\x68\xc2\xd8\xd2\xf0\xdd\x10\xe4\x73\xd7\xc5\xd0\xa4\xfe\x35\x5c\x99\x51\xc9\xda\xf0\x72\xcb\xcd\xe2\xe7\xf2\xe6\x3a\x99\x54\xc2\x59\xfd\xe2\xd5\xfc\xfc\x27\x00\x00\xff\xff\x5a\xa4\x38\xcd\x4e\x08\x00\x00")

func nameServiceGeneratedEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedEndpointsGotemplate,
		"NAME-service/generated/endpoints.gotemplate",
	)
}

func nameServiceGeneratedEndpointsGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/endpoints.gotemplate", size: 2126, mode: os.FileMode(436), modTime: time.Unix(1478035559, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x57\x4f\x6f\xe3\xb6\x13\x3d\x8b\x9f\x62\x7e\x46\xf0\x83\x1d\x38\x52\xcf\x01\x72\xd9\xec\x76\xb3\x68\xb3\x0d\xd2\xa0\x3d\x2c\x16\x0b\x5a\x1a\x4b\x84\x29\x52\x21\x69\x27\xae\xa0\xef\x5e\x0c\x29\xc9\x4a\xec\xc8\xde\xe6\x10\x38\x16\xe7\xcf\x9b\xf7\x66\x46\x74\xc5\xd3\x15\xcf\x11\xec\x26\x65\x2c\x49\xe0\xa1\x10\x16\x96\x42\x22\x54\x46\x6f\x44\x86\x16\x2c\x9a\x0d\x9a\x0b\x2b\x32\x84\x85\x50\x99\x50\xb9\x85\xa5\x36\xe0\x0a\x84\xfc\xfe\xee\x1a\x9c\xe1\xca\x56\xda\xb8\x98\x42\x7c\x71\xb0\x76\x42\x8a\x7f\xd0\x7a\x93\xfe\x34\xc9\x4d\x95\xc6\x7f\xfa\x70\x31\x63\xa2\xa4\x87\x30\x65\xd1\x44\xa1\x4b\x0a\xe7\xaa\x09\x63\xd1\x24\xd7\x92\xab\x3c\xd6\x26\x4f\x9e\x13\x3a\x49\xb5\x72\xf8\xec\x26\xfe\x4c\xe7\x12\xe3\x81\x09\xc5\x4c\x4a\x74\x3c\xe3\x8e\x93\x3f\x3d\xe8\x53\xc2\x24\x17\xae\x58\x2f\xe2\x54\x97\x49\xae\x2f\x56\xc2\x25\xf4\xf7\x12\x13\xb9\x75\xb5\x13\x3c\x91\x22\x8b\xaa\x05\x4c\xea\x3a\xbe\xfb\xf0\xc5\xe3\xbc\xe3\xae\x80\x8b\xa6\x99\xb0\x19\xf3\x4c\xdd\xf2\x15\x7e\xbe\xbf\xbb\x0e\xf5\x40\xc9\x57\x68\x81\x83\x45\x07\x7a\x09\xa8\xb2\x4a\x0b\xe5\x2c\xf0\x0d\x17\x92\x2f\x24\x02\xa7\x73\x4f\x58\x5d\x7f\xd6\x5f\x79\x89\x10\xb7\xe9\x62\xfa\xd6\x34\x1d\x37\xcb\xb5\x4a\x5f\x25\x98\xa6\xee\x19\x5a\x26\xe2\xeb\xf0\x39\x1f\xa4\xf9\xd4\xfd\x37\x83\x6a\x11\x8f\x27\x80\x9a\x45\x41\xd5\x3f\x2a\x27\xb4\xb2\x70\x79\x05\xdf\xbe\xbf\x60\xae\xd5\x29\x18\xd4\x2c\x8a\xe0\xd0\xf1\x07\x5c\x6a\x83\xd3\x8e\xff\x07\xdd\x22\x9b\xcd\x59\xd4\xb0\xc8\xa0\x5b\x1b\x05\xff\x27\xd7\xe0\x50\x7b\xa6\xeb\x1a\x1e\xf4\xef\xfa\x09\xcd\x4b\x80\xd0\x34\x2c\xaa\x6b\xc3\x55\x8e\x70\x26\x08\x56\x7f\x7e\x8b\xae\xd0\x99\x25\x8b\xa8\xae\x3b\xf7\x33\xd1\x56\x76\xf9\x0a\xdf\x57\x7c\x6a\x89\x63\x51\x14\xa5\xee\x79\x4e\x9f\x3d\x5f\x71\x5d\xf7\xae\x1d\x75\xde\xe2\x23\xa6\x3a\xf3\xbc\x0f\x2c\xee\xf1\x71\x8d\x36\x18\x7c\x52\x07\x0d\x6c\xa5\x95\x45\x6f\xf1\x82\xda\x38\x8e\xe9\x21\x11\x52\xd7\x17\x24\x18\x55\xd0\xb0\xc6\x37\xd1\x8e\x18\x10\x65\x25\xb1\x44\xd2\x92\xa6\xe6\x88\x82\x42\x39\x34\x4b\x9e\x22\x73\xdb\x0a\x87\x71\xac\x33\xeb\xd4\x41\xcd\x8e\xf3\x78\x80\x46\x80\x57\x3c\xde\x70\x95\x49\x34\x6c\x07\x3e\x20\x6f\xc3\xf8\x45\x30\xc8\xee\xf4\xae\x90\xd3\x6b\x38\x0a\xd5\x0f\xc4\xd4\xc2\xf9\x2e\xd5\x6c\x17\xbe\x47\x7f\x78\x48\x0c\x3e\xc2\xf9\x70\x28\xce\x44\xdc\x2a\xfa\xb0\xad\x3a\x50\x33\x98\xee\x1b\x05\x55\x07\x56\x73\x40\x63\x34\x25\x67\xd1\x0f\x0a\x5d\xf9\x27\x04\x9b\x7a\x6a\x8f\xcf\x30\x27\xd4\x2d\x84\xcd\x63\x99\xb1\x48\x2c\xbd\xd3\xff\xae\x40\x09\x49\xa1\xba\x49\x51\x42\xfa\x78\xc3\xe9\x31\x58\xc5\xa7\x40\x9b\xcd\xc9\x9d\x35\xac\xae\x83\x50\x24\x53\x4b\x75\xe8\xea\xe3\x3c\x27\x09\x8c\x0d\x00\x08\x5a\x61\xaf\x16\x7a\x70\x68\x2d\x7e\x25\xa1\x5c\xc1\x1d\xc9\xb0\x41\x43\x0b\xd0\x37\x7a\x58\x7b\xfb\xfd\x66\xda\xc8\x4e\x03\x87\xb5\x45\x73\x91\xe9\x92\x0b\x35\x66\x1c\xc3\x9d\x11\x25\x37\x42\x6e\xc9\x65\xb9\x96\x20\x94\xdf\xbd\x83\xf5\x39\x56\xc7\xf4\xc7\x7e\x97\x50\x2d\xf7\xf8\xb8\xeb\xca\x9a\x5a\x62\xf0\x6d\x28\x3d\xb5\xd4\xe5\x55\xe7\x73\x48\x9e\xbd\xf6\x1a\xe8\xf9\xb8\xa7\x14\x51\x74\x2d\x05\x0d\xcd\xfb\xa5\x0a\x9d\x31\xaa\x55\x30\xf9\x0f\x62\x55\x72\x7b\xaa\x54\x21\xc7\x5b\x5a\xa5\xbe\xda\x71\xad\x42\x84\xb7\xc5\x22\x30\x27\xca\x45\xa6\xbd\x60\x95\xdc\x9e\x36\x51\xc3\x19\x94\xdb\x91\xf9\x0a\x2f\x85\x93\x44\x1b\x7d\x7f\x1c\x14\x2d\x78\x1c\x13\xed\x54\x41\x82\x7c\xe3\x12\x9f\x34\x60\xa3\x85\x1c\x12\xad\x47\x70\xa2\x66\xb6\x22\x16\xfb\x46\xfa\x59\xc5\x6c\x35\x36\x66\xef\x57\xec\xed\x8d\xd8\x09\x36\xba\x11\x4f\xdc\x75\x47\xe5\x1a\xdd\x88\x2f\xa6\x6c\xac\x8e\xc3\x7a\xb5\x25\xfe\xc4\x46\xec\xf0\xbc\x7b\x23\x26\x09\xdc\xa0\xac\xd0\x58\x16\xd0\xef\xdd\x31\x0f\xbf\xec\xcb\x0c\xce\x3b\xd3\xf8\xf6\xe3\xec\xb5\x05\x61\xa5\x3b\xcb\x6a\x0e\x1b\x0f\xd8\xcb\x7f\x5e\x66\xfe\x35\x2c\x96\xb0\x19\xbe\x96\xc3\xcf\x02\x84\x15\x6e\xbd\xd2\x59\x86\x19\x2c\xb4\x2b\x88\xde\x2e\x0d\xdd\x81\x4a\xee\x60\xba\x9a\xc1\x53\x21\xd2\xc2\x9b\x4a\x09\x92\xc4\x6a\xa3\x70\x95\xf9\x7b\x1d\xfd\xcc\x89\xaf\xb9\xd2\x4a\xa4\x5c\xde\x20\xcf\xd0\xfc\x86\x5b\xfa\xcd\xe0\xda\x44\x56\x87\x7e\x11\x0e\x52\xae\x60\x81\x5d\x88\x34\x45\x6b\x31\xa3\xdc\x28\x5c\x81\xa6\xcd\xdc\xde\x70\xe1\xaa\x2f\xf6\x6f\xe1\x8a\xbf\xb8\x5c\x63\xb8\x75\x50\xb1\xdf\x7e\xf9\x3e\x3b\x6a\xf8\x06\xba\xe9\x6a\xb6\x8b\xe0\xaf\xaf\xbd\x76\xa9\x7b\x66\x0d\xfb\x37\x00\x00\xff\xff\x09\x00\x51\xc2\x47\x0e\x00\x00")

func nameServiceGeneratedTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedTransport_grpcGotemplate,
		"NAME-service/generated/transport_grpc.gotemplate",
	)
}

func nameServiceGeneratedTransport_grpcGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/transport_grpc.gotemplate", size: 3655, mode: os.FileMode(436), modTime: time.Unix(1478025927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x1a\x29\x91\xe2\x05\xac\xf0\x31\x08\x93\xb3\xce\x1f\x1a\xeb\x76\xa7\x61\x80\x37\xa7\xae\xb2\x03\x52\xed\x42\xff\x8c\xa1\x15\x67\xb5\xd7\x93\xfa\x20\x3b\xbd\x9c\xa3\x9b\x53\x89\x77\x6f\x85\x48\xd5\xe4\x6e\x2e\x87\xdd\x20\xab\x46\x2c\x73\x1e\xa5\x35\x49\x5b\xce\xd7\x0d\xad\xc3\xd3\x0e\xff\xc2\xf3\x9f\x48\x8a\x58\x7e\x01\x00\x00\xff\xff\xaf\xf5\x0f\x3e\x69\x00\x00\x00")

func nameServiceGeneratedTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedTransport_httpGotemplate,
		"NAME-service/generated/transport_http.gotemplate",
	)
}

func nameServiceGeneratedTransport_httpGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/transport_http.gotemplate", size: 105, mode: os.FileMode(436), modTime: time.Unix(1477950707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceHandlersClientClient_handlerGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\xc1\x6e\xdb\x30\x0c\x86\xcf\xd1\x53\x10\x45\x50\xc4\x41\x6a\xdf\x0b\xec\x90\x0e\x6d\xd1\xc3\x86\xa2\xdb\x0b\x28\x36\xe3\x08\xb3\x29\x4d\xa2\xb7\x0e\x84\xde\x7d\x90\xa2\xa4\xcb\x52\xac\xc3\xb0\x4b\xc0\x88\x3f\x7f\x92\x1f\x2d\x02\xdf\x0d\xef\x60\xce\x38\xba\x41\x33\xde\x3e\x63\x3b\xb1\xf5\x70\xfd\x0e\xea\x18\x95\xd3\xed\x17\xdd\x23\xb4\x83\x41\xe2\x9d\xa6\x6e\x40\xaf\x94\x19\x9d\xf5\x0c\x0b\x35\x73\x1b\xb8\x10\x39\xab\xaf\x1f\x6f\x1e\xb2\xe6\x51\xf3\x0e\xae\x62\xbc\x50\x95\x52\x4a\xc4\x6b\xea\x11\xe6\x26\x35\x38\xaf\xfa\x84\xfe\x9b\x69\xb1\xfe\x80\xbc\xb3\x5d\x88\x51\x35\x0d\x88\xcc\x4d\xfd\x51\x8f\x18\x23\x98\xd1\x0d\x38\x22\x71\x80\x83\x56\x6d\x27\x6a\x7f\x15\x2d\x44\xf2\x52\x86\x3a\x7c\x7e\xa5\xc9\xfb\xbc\xcb\xda\xf7\x21\xf7\x49\x01\x1c\x8b\x45\xee\x6d\x8a\xa0\xbe\x9b\xa8\x65\x63\x29\xe5\xd3\x3b\x52\x17\x63\x05\x8b\xa5\xdb\xd4\x47\xd5\xdc\xd4\x4f\xf8\x75\xc2\xc0\x9f\x7f\x38\x2c\x1e\x2b\x40\xef\xad\xaf\x44\x89\x34\x4b\x25\x72\x55\x28\x8f\xc8\xbb\xb4\xf8\x3f\x4c\x96\x18\xaa\x59\xb2\x2a\x04\x9d\xf6\x7a\xcc\x14\x93\x6b\x9d\xb5\x59\x93\x45\x66\x0b\x64\xb9\xa8\xea\x87\x70\xa3\x03\xa6\x09\x8b\xa4\x69\x60\xdd\x75\xd0\x4e\x81\xed\x08\x9b\x29\x18\xc2\x10\x60\xb0\xbd\x69\x61\x6b\x3d\x18\x62\xf4\xce\x23\x1b\xea\x13\xdb\xbd\xcf\xdd\xa0\xfb\xb5\xef\x63\x5c\x95\x36\x48\xdd\xcb\x60\x87\x3f\x25\x8e\x51\x2d\x9b\x94\xf3\x7b\x40\x69\xd6\x37\xd1\xc9\xff\xa5\xf5\xb7\xac\x34\x75\xb0\xc8\x3f\xaf\x52\xab\x4e\xde\x9f\xd0\xa1\x66\xec\xaa\xd3\xe7\x5b\x9a\xc6\xaa\x98\x1e\x56\xdc\x67\xca\xb7\x7b\x0d\x97\xbf\x67\xee\xed\x29\xce\x21\x20\xfc\xd9\xe2\x2d\x87\x97\x83\x9c\x9f\x63\x96\x8f\xc1\x93\x27\xb8\x2c\x47\x59\x01\x99\x41\x25\xd1\x5e\x72\x0c\x7e\x06\x00\x00\xff\xff\x59\xb7\xfe\xb0\x1a\x04\x00\x00")

func nameServiceHandlersClientClient_handlerGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceHandlersClientClient_handlerGotemplate,
		"NAME-service/handlers/client/client_handler.gotemplate",
	)
}

func nameServiceHandlersClientClient_handlerGotemplate() (*asset, error) {
	bytes, err := nameServiceHandlersClientClient_handlerGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/handlers/client/client_handler.gotemplate", size: 1050, mode: os.FileMode(436), modTime: time.Unix(1478025927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceHandlersServerServer_handlerGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x4d\x6e\xdb\x30\x10\x85\xd7\x9c\x53\x0c\x0c\xa1\x90\x0a\x87\xda\x07\xc8\xa6\x05\x5a\x74\x91\x20\x68\x7b\x01\x46\x9e\xc8\x4c\x65\x92\x25\xc7\x8e\x83\x01\xcf\xd4\x43\xf4\x62\x05\x29\xb9\x36\xfa\x83\x76\x25\x51\xfa\xe6\xe9\xbd\xa7\x09\x66\xf8\x62\x46\xc2\xad\x71\x9b\x89\x22\x80\xdd\x05\x1f\x19\x5b\x50\xab\xd1\x4f\xc6\x8d\xda\xc7\xb1\x3f\xf6\x8e\xb8\x1f\xbc\x63\x3a\xf2\x0a\x40\x85\x07\x5c\x89\xe8\xfb\x37\x1f\x2a\x7f\x6f\x78\x8b\x57\x39\xaf\xa0\x03\xe8\x7b\xbc\xa3\xe7\x4f\x14\x0f\x76\x20\x8c\xc4\xfb\xe8\x12\x1a\x74\xe6\xfb\xb7\x03\xad\x31\xb1\x61\x9a\x28\x25\xb4\xbb\x30\xd1\x8e\x1c\x1b\xb6\xde\xa1\x7f\xc4\x65\x4a\xc3\xe3\xde\x0d\x17\x32\x6d\x87\xe1\x41\x8b\xbc\xf7\x77\x66\x47\xa8\x4f\x5c\x39\xe5\x5c\x4e\x14\x51\x40\xcd\x5f\xc3\x62\x6d\x4e\x76\x06\xec\x40\x92\x21\x03\xf0\x4b\xa0\xbf\x11\x98\x38\xee\x07\x96\x0c\x20\xf2\x6c\x79\x8b\x0d\x13\x5e\xdf\xa0\xc6\x9c\x41\x89\x44\xe3\x46\xc2\xc6\x96\x67\x0d\xd3\x4f\x23\xb7\xc4\x5b\xbf\x49\x05\x52\x7d\x8f\x22\x8d\x5d\xcc\x9d\x53\xa6\x73\x3c\xa5\x6a\xc0\x36\x15\x92\xe9\x8f\x5e\xba\x4b\x95\x76\xe0\x23\x2e\x3f\x40\xbf\x9d\xaf\x6b\xb4\x0e\x5f\x5f\xf6\xd2\x58\xfd\x91\xbe\xee\x29\xf1\xe7\x97\x70\x6a\xa7\xc3\xf6\x77\x28\x05\xef\x12\x5d\x50\x6b\xa4\x18\x7d\xec\x04\x94\x52\x07\x13\x31\x52\x0a\xf8\xef\xb9\x82\x57\xf4\xe6\x3f\xe0\x2a\x7e\x6e\xf1\xa9\xb6\xf8\x0b\x79\x4b\x29\x99\x91\xf4\x3b\x4b\xd3\x26\x95\xad\xaa\x43\x73\xab\x27\xf5\xa7\x45\xf0\x1a\x17\x45\x72\x9b\x13\xba\x58\xaa\x8b\xf0\xaa\x58\x5b\xa3\xb3\x13\xd4\x17\x15\xcc\x19\x44\xae\x70\xbe\xfb\x11\x00\x00\xff\xff\x14\x16\x8e\x3b\x02\x03\x00\x00")

func nameServiceHandlersServerServer_handlerGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceHandlersServerServer_handlerGotemplate,
		"NAME-service/handlers/server/server_handler.gotemplate",
	)
}

func nameServiceHandlersServerServer_handlerGotemplate() (*asset, error) {
	bytes, err := nameServiceHandlersServerServer_handlerGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/handlers/server/server_handler.gotemplate", size: 770, mode: os.FileMode(436), modTime: time.Unix(1478036246, 0)}
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
	"NAME-service/NAME-client/client_main.gotemplate": nameServiceNameClientClient_mainGotemplate,
	"NAME-service/NAME-server/server_main.gotemplate": nameServiceNameServerServer_mainGotemplate,
	"NAME-service/generated/client/grpc/client.gotemplate": nameServiceGeneratedClientGrpcClientGotemplate,
	"NAME-service/generated/client/http/client.gotemplate": nameServiceGeneratedClientHttpClientGotemplate,
	"NAME-service/generated/endpoints.gotemplate": nameServiceGeneratedEndpointsGotemplate,
	"NAME-service/generated/transport_grpc.gotemplate": nameServiceGeneratedTransport_grpcGotemplate,
	"NAME-service/generated/transport_http.gotemplate": nameServiceGeneratedTransport_httpGotemplate,
	"NAME-service/handlers/client/client_handler.gotemplate": nameServiceHandlersClientClient_handlerGotemplate,
	"NAME-service/handlers/server/server_handler.gotemplate": nameServiceHandlersServerServer_handlerGotemplate,
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
	"NAME-service": &bintree{nil, map[string]*bintree{
		"NAME-client": &bintree{nil, map[string]*bintree{
			"client_main.gotemplate": &bintree{nameServiceNameClientClient_mainGotemplate, map[string]*bintree{}},
		}},
		"NAME-server": &bintree{nil, map[string]*bintree{
			"server_main.gotemplate": &bintree{nameServiceNameServerServer_mainGotemplate, map[string]*bintree{}},
		}},
		"generated": &bintree{nil, map[string]*bintree{
			"client": &bintree{nil, map[string]*bintree{
				"grpc": &bintree{nil, map[string]*bintree{
					"client.gotemplate": &bintree{nameServiceGeneratedClientGrpcClientGotemplate, map[string]*bintree{}},
				}},
				"http": &bintree{nil, map[string]*bintree{
					"client.gotemplate": &bintree{nameServiceGeneratedClientHttpClientGotemplate, map[string]*bintree{}},
				}},
			}},
			"endpoints.gotemplate": &bintree{nameServiceGeneratedEndpointsGotemplate, map[string]*bintree{}},
			"transport_grpc.gotemplate": &bintree{nameServiceGeneratedTransport_grpcGotemplate, map[string]*bintree{}},
			"transport_http.gotemplate": &bintree{nameServiceGeneratedTransport_httpGotemplate, map[string]*bintree{}},
		}},
		"handlers": &bintree{nil, map[string]*bintree{
			"client": &bintree{nil, map[string]*bintree{
				"client_handler.gotemplate": &bintree{nameServiceHandlersClientClient_handlerGotemplate, map[string]*bintree{}},
			}},
			"server": &bintree{nil, map[string]*bintree{
				"server_handler.gotemplate": &bintree{nameServiceHandlersServerServer_handlerGotemplate, map[string]*bintree{}},
			}},
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

