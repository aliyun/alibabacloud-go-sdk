// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDnsGtmInstanceRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeDnsGtmInstanceRequest
	GetLang() *string
}

type DescribeDnsGtmInstanceRequest struct {
	// The ID of the instance about which you want to query the information.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language in which you want the values of some response parameters to be returned. These response parameters support multiple languages. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeDnsGtmInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmInstanceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmInstanceRequest) SetInstanceId(v string) *DescribeDnsGtmInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceRequest) SetLang(v string) *DescribeDnsGtmInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstanceRequest) Validate() error {
	return dara.Validate(s)
}
