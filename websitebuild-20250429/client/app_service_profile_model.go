// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppServiceProfile interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *AppServiceProfile
	GetBizId() *string
	SetDesignType(v string) *AppServiceProfile
	GetDesignType() *string
	SetDesignTypeText(v string) *AppServiceProfile
	GetDesignTypeText() *string
	SetInstanceId(v string) *AppServiceProfile
	GetInstanceId() *string
	SetOrderId(v string) *AppServiceProfile
	GetOrderId() *string
	SetServiceSpec(v string) *AppServiceProfile
	GetServiceSpec() *string
	SetServiceSpecText(v string) *AppServiceProfile
	GetServiceSpecText() *string
}

type AppServiceProfile struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	DesignType      *string `json:"DesignType,omitempty" xml:"DesignType,omitempty"`
	DesignTypeText  *string `json:"DesignTypeText,omitempty" xml:"DesignTypeText,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId         *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ServiceSpec     *string `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	ServiceSpecText *string `json:"ServiceSpecText,omitempty" xml:"ServiceSpecText,omitempty"`
}

func (s AppServiceProfile) String() string {
	return dara.Prettify(s)
}

func (s AppServiceProfile) GoString() string {
	return s.String()
}

func (s *AppServiceProfile) GetBizId() *string {
	return s.BizId
}

func (s *AppServiceProfile) GetDesignType() *string {
	return s.DesignType
}

func (s *AppServiceProfile) GetDesignTypeText() *string {
	return s.DesignTypeText
}

func (s *AppServiceProfile) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AppServiceProfile) GetOrderId() *string {
	return s.OrderId
}

func (s *AppServiceProfile) GetServiceSpec() *string {
	return s.ServiceSpec
}

func (s *AppServiceProfile) GetServiceSpecText() *string {
	return s.ServiceSpecText
}

func (s *AppServiceProfile) SetBizId(v string) *AppServiceProfile {
	s.BizId = &v
	return s
}

func (s *AppServiceProfile) SetDesignType(v string) *AppServiceProfile {
	s.DesignType = &v
	return s
}

func (s *AppServiceProfile) SetDesignTypeText(v string) *AppServiceProfile {
	s.DesignTypeText = &v
	return s
}

func (s *AppServiceProfile) SetInstanceId(v string) *AppServiceProfile {
	s.InstanceId = &v
	return s
}

func (s *AppServiceProfile) SetOrderId(v string) *AppServiceProfile {
	s.OrderId = &v
	return s
}

func (s *AppServiceProfile) SetServiceSpec(v string) *AppServiceProfile {
	s.ServiceSpec = &v
	return s
}

func (s *AppServiceProfile) SetServiceSpecText(v string) *AppServiceProfile {
	s.ServiceSpecText = &v
	return s
}

func (s *AppServiceProfile) Validate() error {
	return dara.Validate(s)
}
