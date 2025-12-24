// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GenerateUploadAuthRequest
	GetInstanceId() *string
	SetPurpose(v string) *GenerateUploadAuthRequest
	GetPurpose() *string
	SetType(v string) *GenerateUploadAuthRequest
	GetType() *string
}

type GenerateUploadAuthRequest struct {
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_111ccc11xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 文件用途
	//
	// example:
	//
	// user_import
	Purpose *string `json:"Purpose,omitempty" xml:"Purpose,omitempty"`
	// 文件类型，目前只支持image,最大1M
	//
	// example:
	//
	// image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GenerateUploadAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadAuthRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadAuthRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateUploadAuthRequest) GetPurpose() *string {
	return s.Purpose
}

func (s *GenerateUploadAuthRequest) GetType() *string {
	return s.Type
}

func (s *GenerateUploadAuthRequest) SetInstanceId(v string) *GenerateUploadAuthRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateUploadAuthRequest) SetPurpose(v string) *GenerateUploadAuthRequest {
	s.Purpose = &v
	return s
}

func (s *GenerateUploadAuthRequest) SetType(v string) *GenerateUploadAuthRequest {
	s.Type = &v
	return s
}

func (s *GenerateUploadAuthRequest) Validate() error {
	return dara.Validate(s)
}
