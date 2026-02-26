// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliveryInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *DeliveryInfo
	GetDisplayName() *string
	SetId(v string) *DeliveryInfo
	GetId() *string
	SetPostFee(v int64) *DeliveryInfo
	GetPostFee() *int64
	SetServiceType(v int64) *DeliveryInfo
	GetServiceType() *int64
}

type DeliveryInfo struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 20
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	PostFee *int64 `json:"postFee,omitempty" xml:"postFee,omitempty"`
	// serviceType
	//
	// example:
	//
	// -4
	ServiceType *int64 `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeliveryInfo) String() string {
	return dara.Prettify(s)
}

func (s DeliveryInfo) GoString() string {
	return s.String()
}

func (s *DeliveryInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DeliveryInfo) GetId() *string {
	return s.Id
}

func (s *DeliveryInfo) GetPostFee() *int64 {
	return s.PostFee
}

func (s *DeliveryInfo) GetServiceType() *int64 {
	return s.ServiceType
}

func (s *DeliveryInfo) SetDisplayName(v string) *DeliveryInfo {
	s.DisplayName = &v
	return s
}

func (s *DeliveryInfo) SetId(v string) *DeliveryInfo {
	s.Id = &v
	return s
}

func (s *DeliveryInfo) SetPostFee(v int64) *DeliveryInfo {
	s.PostFee = &v
	return s
}

func (s *DeliveryInfo) SetServiceType(v int64) *DeliveryInfo {
	s.ServiceType = &v
	return s
}

func (s *DeliveryInfo) Validate() error {
	return dara.Validate(s)
}
