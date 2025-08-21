// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeInstanceAutoRenewAttributeResponseBody
	GetCode() *int32
	SetInstanceRenewAttributes(v *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) *DescribeInstanceAutoRenewAttributeResponseBody
	GetInstanceRenewAttributes() *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes
	SetRequestId(v string) *DescribeInstanceAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type DescribeInstanceAutoRenewAttributeResponseBody struct {
	// The returned service code. A value of 0 indicates that the operation was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The renewal status of the instance.
	InstanceRenewAttributes *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes `json:"InstanceRenewAttributes,omitempty" xml:"InstanceRenewAttributes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4DD66F05-3116-4BAA-B588-52EB2E7F431D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetInstanceRenewAttributes() *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes {
	return s.InstanceRenewAttributes
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetCode(v int32) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetInstanceRenewAttributes(v *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.InstanceRenewAttributes = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes struct {
	InstanceRenewAttribute []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute `json:"InstanceRenewAttribute,omitempty" xml:"InstanceRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) GetInstanceRenewAttribute() []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	return s.InstanceRenewAttribute
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) SetInstanceRenewAttribute(v []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes {
	s.InstanceRenewAttribute = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute struct {
	// The renewal type of the instance.
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false**: disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The unit of the auto-renewal period.
	//
	// example:
	//
	// 0
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5ci7l7k1m9m2zmhp4iw3o****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetDuration() *string {
	return s.Duration
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetAutoRenewal(v bool) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetDuration(v string) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetInstanceId(v string) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) Validate() error {
	return dara.Validate(s)
}
