// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmInstanceStatusRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmInstanceStatusRequest
	GetLang() *string
}

type DescribeDnsGtmInstanceStatusRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmInstanceStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmInstanceStatusRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusRequest) SetLang(v string) *DescribeDnsGtmInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
