// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceSystemCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmInstanceSystemCnameRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmInstanceSystemCnameRequest
	GetLang() *string
}

type DescribeDnsGtmInstanceSystemCnameRequest struct {
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

func (s DescribeDnsGtmInstanceSystemCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceSystemCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceSystemCnameRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) SetLang(v string) *DescribeDnsGtmInstanceSystemCnameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceSystemCnameRequest) Validate() error {
	return dara.Validate(s)
}
