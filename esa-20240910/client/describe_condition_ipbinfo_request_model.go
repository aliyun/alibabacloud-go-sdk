// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConditionIPBInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *DescribeConditionIPBInfoRequest
	GetDataId() *string
}

type DescribeConditionIPBInfoRequest struct {
	// The configuration ID. Valid values:
	//
	// - condition_region_config_cn: Provides a mapping list of Chinese region names and their corresponding codes.
	//
	// - condition_region_config_en: Provides a mapping list of English region names and their corresponding codes.
	//
	// - condition_isp_config_cn: Provides a mapping list of Chinese ISP names and their corresponding codes.
	//
	// - condition_isp_config_en: Provides a mapping list of English ISP names and their corresponding codes.
	//
	// - condition_country_config_cn: Provides a mapping list of Chinese country names and their corresponding codes.
	//
	// - condition_country_config_en: Provides a mapping list of English country names and their corresponding codes.
	//
	// This parameter is required.
	//
	// example:
	//
	// condition_region_config_cn
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s DescribeConditionIPBInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeConditionIPBInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeConditionIPBInfoRequest) GetDataId() *string {
	return s.DataId
}

func (s *DescribeConditionIPBInfoRequest) SetDataId(v string) *DescribeConditionIPBInfoRequest {
	s.DataId = &v
	return s
}

func (s *DescribeConditionIPBInfoRequest) Validate() error {
	return dara.Validate(s)
}
