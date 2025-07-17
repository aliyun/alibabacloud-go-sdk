// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpdateEnvironmentRequest
	GetAliyunLang() *string
	SetEnvironmentId(v string) *UpdateEnvironmentRequest
	GetEnvironmentId() *string
	SetEnvironmentName(v string) *UpdateEnvironmentRequest
	GetEnvironmentName() *string
	SetFeePackage(v string) *UpdateEnvironmentRequest
	GetFeePackage() *string
	SetRegionId(v string) *UpdateEnvironmentRequest
	GetRegionId() *string
}

type UpdateEnvironmentRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The name of the environment instance.
	//
	// example:
	//
	// env1
	EnvironmentName *string `json:"EnvironmentName,omitempty" xml:"EnvironmentName,omitempty"`
	// The payable resource plan. Valid values:
	//
	// 	- If the EnvironmentType parameter is set to CS, set the value to CS_Basic or CS_Pro. Default value: CS_Basic.
	//
	// 	- Otherwise, leave the parameter empty.
	//
	// example:
	//
	// CS_Basic
	FeePackage *string `json:"FeePackage,omitempty" xml:"FeePackage,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpdateEnvironmentRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateEnvironmentRequest) GetEnvironmentName() *string {
	return s.EnvironmentName
}

func (s *UpdateEnvironmentRequest) GetFeePackage() *string {
	return s.FeePackage
}

func (s *UpdateEnvironmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnvironmentRequest) SetAliyunLang(v string) *UpdateEnvironmentRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetEnvironmentId(v string) *UpdateEnvironmentRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetEnvironmentName(v string) *UpdateEnvironmentRequest {
	s.EnvironmentName = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetFeePackage(v string) *UpdateEnvironmentRequest {
	s.FeePackage = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetRegionId(v string) *UpdateEnvironmentRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
