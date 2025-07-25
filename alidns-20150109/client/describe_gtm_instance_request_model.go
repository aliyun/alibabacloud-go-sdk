// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGtmInstanceRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeGtmInstanceRequest
	GetLang() *string
	SetNeedDetailAttributes(v bool) *DescribeGtmInstanceRequest
	GetNeedDetailAttributes() *bool
}

type DescribeGtmInstanceRequest struct {
	// The ID of the GTM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the values of specific response parameters.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether additional information is required. Default value: **false**. If the value is **true**, the AccessStrategyNum and AddressPoolNum parameters are returned.
	//
	// example:
	//
	// false
	NeedDetailAttributes *bool `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
}

func (s DescribeGtmInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmInstanceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmInstanceRequest) GetNeedDetailAttributes() *bool {
	return s.NeedDetailAttributes
}

func (s *DescribeGtmInstanceRequest) SetInstanceId(v string) *DescribeGtmInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceRequest) SetLang(v string) *DescribeGtmInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceRequest) SetNeedDetailAttributes(v bool) *DescribeGtmInstanceRequest {
	s.NeedDetailAttributes = &v
	return s
}

func (s *DescribeGtmInstanceRequest) Validate() error {
	return dara.Validate(s)
}
