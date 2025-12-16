// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSortScriptFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SaveSortScriptFileRequest
	GetContent() *string
	SetVersion(v int32) *SaveSortScriptFileRequest
	GetVersion() *int32
}

type SaveSortScriptFileRequest struct {
	// The script content that is encoded in the Base64 format.
	//
	// example:
	//
	// 4769#0: *28194492991 a client request body is buffered to a temporary file /usr/local/webserver/openresty/nginx/client_body_temp/0000624474,
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The version number of the script content.
	//
	// example:
	//
	// 2022-12-01
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SaveSortScriptFileRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSortScriptFileRequest) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileRequest) GetContent() *string {
	return s.Content
}

func (s *SaveSortScriptFileRequest) GetVersion() *int32 {
	return s.Version
}

func (s *SaveSortScriptFileRequest) SetContent(v string) *SaveSortScriptFileRequest {
	s.Content = &v
	return s
}

func (s *SaveSortScriptFileRequest) SetVersion(v int32) *SaveSortScriptFileRequest {
	s.Version = &v
	return s
}

func (s *SaveSortScriptFileRequest) Validate() error {
	return dara.Validate(s)
}
