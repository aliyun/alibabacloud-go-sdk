// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGtmInstanceStatusRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeGtmInstanceStatusRequest
	GetLang() *string
}

type DescribeGtmInstanceStatusRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGtmInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmInstanceStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmInstanceStatusRequest) SetInstanceId(v string) *DescribeGtmInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceStatusRequest) SetLang(v string) *DescribeGtmInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
