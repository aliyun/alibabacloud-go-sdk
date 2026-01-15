// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTestFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerModuleId(v string) *ValidateTestFileRequest
	GetCustomerModuleId() *string
	SetFilePath(v string) *ValidateTestFileRequest
	GetFilePath() *string
}

type ValidateTestFileRequest struct {
	// Model ID.
	//
	// example:
	//
	// 456
	CustomerModuleId *string `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
	// File path.
	//
	// example:
	//
	// saf/xxxxxx/xxxxx.pmml
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
}

func (s ValidateTestFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateTestFileRequest) GoString() string {
	return s.String()
}

func (s *ValidateTestFileRequest) GetCustomerModuleId() *string {
	return s.CustomerModuleId
}

func (s *ValidateTestFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ValidateTestFileRequest) SetCustomerModuleId(v string) *ValidateTestFileRequest {
	s.CustomerModuleId = &v
	return s
}

func (s *ValidateTestFileRequest) SetFilePath(v string) *ValidateTestFileRequest {
	s.FilePath = &v
	return s
}

func (s *ValidateTestFileRequest) Validate() error {
	return dara.Validate(s)
}
