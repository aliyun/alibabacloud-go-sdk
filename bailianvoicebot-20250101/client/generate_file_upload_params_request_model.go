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
	SetBusinessUnitId(v string) *GenerateFileUploadParamsRequest
	GetBusinessUnitId() *string
	SetFileName(v string) *GenerateFileUploadParamsRequest
	GetFileName() *string
}

type GenerateFileUploadParamsRequest struct {
	// example:
	//
	// CloneVoice
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// test.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
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

func (s *GenerateFileUploadParamsRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *GenerateFileUploadParamsRequest) GetFileName() *string {
	return s.FileName
}

func (s *GenerateFileUploadParamsRequest) SetBusinessType(v string) *GenerateFileUploadParamsRequest {
	s.BusinessType = &v
	return s
}

func (s *GenerateFileUploadParamsRequest) SetBusinessUnitId(v string) *GenerateFileUploadParamsRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *GenerateFileUploadParamsRequest) SetFileName(v string) *GenerateFileUploadParamsRequest {
	s.FileName = &v
	return s
}

func (s *GenerateFileUploadParamsRequest) Validate() error {
	return dara.Validate(s)
}
