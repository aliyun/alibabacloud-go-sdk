// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileUploadParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GenerateFileUploadParamsRequest
	GetBusinessType() *string
	SetFileName(v string) *GenerateFileUploadParamsRequest
	GetFileName() *string
	SetInstanceId(v string) *GenerateFileUploadParamsRequest
	GetInstanceId() *string
}

type GenerateFileUploadParamsRequest struct {
	// example:
	//
	// CloneVoice
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// test.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateFileUploadParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileUploadParamsRequest) GoString() string {
	return s.String()
}

func (s *GenerateFileUploadParamsRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GenerateFileUploadParamsRequest) GetFileName() *string {
	return s.FileName
}

func (s *GenerateFileUploadParamsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateFileUploadParamsRequest) SetBusinessType(v string) *GenerateFileUploadParamsRequest {
	s.BusinessType = &v
	return s
}

func (s *GenerateFileUploadParamsRequest) SetFileName(v string) *GenerateFileUploadParamsRequest {
	s.FileName = &v
	return s
}

func (s *GenerateFileUploadParamsRequest) SetInstanceId(v string) *GenerateFileUploadParamsRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateFileUploadParamsRequest) Validate() error {
	return dara.Validate(s)
}
