// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAddressPoolAvailableConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest
	GetLang() *string
}

type DescribeDnsGtmAddressPoolAvailableConfigRequest struct {
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

func (s DescribeDnsGtmAddressPoolAvailableConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) SetInstanceId(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) SetLang(v string) *DescribeDnsGtmAddressPoolAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigRequest) Validate() error {
	return dara.Validate(s)
}
