// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateFileInfoRequest
	GetDescription() *string
	SetFileId(v string) *UpdateFileInfoRequest
	GetFileId() *string
}

type UpdateFileInfoRequest struct {
	// The file description.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试使用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file. Each file has a unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-16713accddtgtj6340jgnclhwsg1813f718db2f7
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s UpdateFileInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFileInfoRequest) GetFileId() *string {
	return s.FileId
}

func (s *UpdateFileInfoRequest) SetDescription(v string) *UpdateFileInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateFileInfoRequest) SetFileId(v string) *UpdateFileInfoRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileInfoRequest) Validate() error {
	return dara.Validate(s)
}
