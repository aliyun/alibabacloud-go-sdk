// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdditionalVpcLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalVpcId(v string) *CreateAdditionalVpcLinkRequest
	GetAdditionalVpcId() *string
	SetAdditionalVswitchId(v string) *CreateAdditionalVpcLinkRequest
	GetAdditionalVswitchId() *string
	SetInstanceId(v string) *CreateAdditionalVpcLinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateAdditionalVpcLinkRequest
	GetRegionId() *string
	SetSecurityToken(v string) *CreateAdditionalVpcLinkRequest
	GetSecurityToken() *string
}

type CreateAdditionalVpcLinkRequest struct {
	// This parameter is required.
	AdditionalVpcId *string `json:"AdditionalVpcId,omitempty" xml:"AdditionalVpcId,omitempty"`
	// This parameter is required.
	AdditionalVswitchId *string `json:"AdditionalVswitchId,omitempty" xml:"AdditionalVswitchId,omitempty"`
	// This parameter is required.
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateAdditionalVpcLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAdditionalVpcLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAdditionalVpcLinkRequest) GetAdditionalVpcId() *string {
	return s.AdditionalVpcId
}

func (s *CreateAdditionalVpcLinkRequest) GetAdditionalVswitchId() *string {
	return s.AdditionalVswitchId
}

func (s *CreateAdditionalVpcLinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAdditionalVpcLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAdditionalVpcLinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateAdditionalVpcLinkRequest) SetAdditionalVpcId(v string) *CreateAdditionalVpcLinkRequest {
	s.AdditionalVpcId = &v
	return s
}

func (s *CreateAdditionalVpcLinkRequest) SetAdditionalVswitchId(v string) *CreateAdditionalVpcLinkRequest {
	s.AdditionalVswitchId = &v
	return s
}

func (s *CreateAdditionalVpcLinkRequest) SetInstanceId(v string) *CreateAdditionalVpcLinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAdditionalVpcLinkRequest) SetRegionId(v string) *CreateAdditionalVpcLinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAdditionalVpcLinkRequest) SetSecurityToken(v string) *CreateAdditionalVpcLinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAdditionalVpcLinkRequest) Validate() error {
	return dara.Validate(s)
}
