// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategyAvailableConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGtmAccessStrategyAvailableConfigRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeGtmAccessStrategyAvailableConfigRequest
	GetLang() *string
}

type DescribeGtmAccessStrategyAvailableConfigRequest struct {
	// The ID of the Global Traffic Manager (GTM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) SetInstanceId(v string) *DescribeGtmAccessStrategyAvailableConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) SetLang(v string) *DescribeGtmAccessStrategyAvailableConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigRequest) Validate() error {
	return dara.Validate(s)
}
