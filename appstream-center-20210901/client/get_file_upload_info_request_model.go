// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *GetFileUploadInfoRequest
	GetFileType() *string
}

type GetFileUploadInfoRequest struct {
	// The file type. Currently supported values: BrowserBookmarks and BrowserRestrictionUrls.
	//
	// This parameter is required.
	//
	// example:
	//
	// BrowserBookmarks
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GetFileUploadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoRequest) GetFileType() *string {
	return s.FileType
}

func (s *GetFileUploadInfoRequest) SetFileType(v string) *GetFileUploadInfoRequest {
	s.FileType = &v
	return s
}

func (s *GetFileUploadInfoRequest) Validate() error {
	return dara.Validate(s)
}
