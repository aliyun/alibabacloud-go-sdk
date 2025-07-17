// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *InitEnvironmentRequest
	GetAliyunLang() *string
	SetCreateAuthToken(v bool) *InitEnvironmentRequest
	GetCreateAuthToken() *bool
	SetEnvironmentId(v string) *InitEnvironmentRequest
	GetEnvironmentId() *string
	SetManagedType(v string) *InitEnvironmentRequest
	GetManagedType() *string
	SetRegionId(v string) *InitEnvironmentRequest
	GetRegionId() *string
}

type InitEnvironmentRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// Specifies whether to create a token to improve data security.
	//
	// example:
	//
	// false
	CreateAuthToken *bool `json:"CreateAuthToken,omitempty" xml:"CreateAuthToken,omitempty"`
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// Whether agents or exporters are managed. Valid values:
	//
	// 	- none: No. By default, no managed agents or exporters are provided for ACK clusters.
	//
	// 	- agent: Agents are managed. By default, managed agents are provided for ASK clusters, ACS clusters, and ACK One clusters.
	//
	// 	- agent-exproter: Agents and exporters are managed. By default, managed agents and exporters are provided for cloud services.
	//
	// example:
	//
	// agent
	ManagedType *string `json:"ManagedType,omitempty" xml:"ManagedType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s InitEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *InitEnvironmentRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *InitEnvironmentRequest) GetCreateAuthToken() *bool {
	return s.CreateAuthToken
}

func (s *InitEnvironmentRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *InitEnvironmentRequest) GetManagedType() *string {
	return s.ManagedType
}

func (s *InitEnvironmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitEnvironmentRequest) SetAliyunLang(v string) *InitEnvironmentRequest {
	s.AliyunLang = &v
	return s
}

func (s *InitEnvironmentRequest) SetCreateAuthToken(v bool) *InitEnvironmentRequest {
	s.CreateAuthToken = &v
	return s
}

func (s *InitEnvironmentRequest) SetEnvironmentId(v string) *InitEnvironmentRequest {
	s.EnvironmentId = &v
	return s
}

func (s *InitEnvironmentRequest) SetManagedType(v string) *InitEnvironmentRequest {
	s.ManagedType = &v
	return s
}

func (s *InitEnvironmentRequest) SetRegionId(v string) *InitEnvironmentRequest {
	s.RegionId = &v
	return s
}

func (s *InitEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
