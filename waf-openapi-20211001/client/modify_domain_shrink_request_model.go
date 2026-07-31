// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *ModifyDomainShrinkRequest
	GetAccessType() *string
	SetDomain(v string) *ModifyDomainShrinkRequest
	GetDomain() *string
	SetDomainId(v string) *ModifyDomainShrinkRequest
	GetDomainId() *string
	SetInstanceId(v string) *ModifyDomainShrinkRequest
	GetInstanceId() *string
	SetListenShrink(v string) *ModifyDomainShrinkRequest
	GetListenShrink() *string
	SetRedirectShrink(v string) *ModifyDomainShrinkRequest
	GetRedirectShrink() *string
	SetRegionId(v string) *ModifyDomainShrinkRequest
	GetRegionId() *string
}

type ModifyDomainShrinkRequest struct {
	// The access type of the WAF instance. Valid values:
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name to operate on.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The domain name ID.
	//
	// example:
	//
	// www.aliyundoc.com-waf
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listening configuration.
	//
	// This parameter is required.
	ListenShrink *string `json:"Listen,omitempty" xml:"Listen,omitempty"`
	// The forwarding configuration.
	//
	// This parameter is required.
	RedirectShrink *string `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainShrinkRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *ModifyDomainShrinkRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyDomainShrinkRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *ModifyDomainShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDomainShrinkRequest) GetListenShrink() *string {
	return s.ListenShrink
}

func (s *ModifyDomainShrinkRequest) GetRedirectShrink() *string {
	return s.RedirectShrink
}

func (s *ModifyDomainShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDomainShrinkRequest) SetAccessType(v string) *ModifyDomainShrinkRequest {
	s.AccessType = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetDomain(v string) *ModifyDomainShrinkRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetDomainId(v string) *ModifyDomainShrinkRequest {
	s.DomainId = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetInstanceId(v string) *ModifyDomainShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetListenShrink(v string) *ModifyDomainShrinkRequest {
	s.ListenShrink = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetRedirectShrink(v string) *ModifyDomainShrinkRequest {
	s.RedirectShrink = &v
	return s
}

func (s *ModifyDomainShrinkRequest) SetRegionId(v string) *ModifyDomainShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
