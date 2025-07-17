// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFileUploadLeaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryType(v string) *ApplyFileUploadLeaseRequest
	GetCategoryType() *string
	SetFileName(v string) *ApplyFileUploadLeaseRequest
	GetFileName() *string
	SetMd5(v string) *ApplyFileUploadLeaseRequest
	GetMd5() *string
	SetSizeInBytes(v string) *ApplyFileUploadLeaseRequest
	GetSizeInBytes() *string
	SetUseInternalEndpoint(v bool) *ApplyFileUploadLeaseRequest
	GetUseInternalEndpoint() *bool
}

type ApplyFileUploadLeaseRequest struct {
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The name of the uploaded document, including the extension. Supported formats: pdf, doc, docx, md, txt, ppt, and pptx. The document name must be 4 to 128 characters in length.
	//
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The MD5 value of the uploaded document. This parameter is verified by the server (not in the current version).
	//
	// This parameter is required.
	//
	// example:
	//
	// 19657c391f6c70bcea63c154d8606bb3
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The size of the uploaded document, in bytes. This parameter is verified by the server (not in the current version). Valid values: 1 to 100000000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	SizeInBytes         *string `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	UseInternalEndpoint *bool   `json:"UseInternalEndpoint,omitempty" xml:"UseInternalEndpoint,omitempty"`
}

func (s ApplyFileUploadLeaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyFileUploadLeaseRequest) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *ApplyFileUploadLeaseRequest) GetFileName() *string {
	return s.FileName
}

func (s *ApplyFileUploadLeaseRequest) GetMd5() *string {
	return s.Md5
}

func (s *ApplyFileUploadLeaseRequest) GetSizeInBytes() *string {
	return s.SizeInBytes
}

func (s *ApplyFileUploadLeaseRequest) GetUseInternalEndpoint() *bool {
	return s.UseInternalEndpoint
}

func (s *ApplyFileUploadLeaseRequest) SetCategoryType(v string) *ApplyFileUploadLeaseRequest {
	s.CategoryType = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetFileName(v string) *ApplyFileUploadLeaseRequest {
	s.FileName = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetMd5(v string) *ApplyFileUploadLeaseRequest {
	s.Md5 = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetSizeInBytes(v string) *ApplyFileUploadLeaseRequest {
	s.SizeInBytes = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetUseInternalEndpoint(v bool) *ApplyFileUploadLeaseRequest {
	s.UseInternalEndpoint = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) Validate() error {
	return dara.Validate(s)
}
