// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnConditionIPBInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *DescribeCdnConditionIPBInfoRequest
	GetDataId() *string
}

type DescribeCdnConditionIPBInfoRequest struct {
	// The configuration ID. Valid values:
	//
	// 	- condition_region_config_cn
	//
	// 	- condition_region_config_en
	//
	// 	- condition_isp_config_cn
	//
	// 	- condition_isp_config_en
	//
	// 	- condition_country_config_cn
	//
	// 	- condition_country_config_en
	//
	// This parameter is required.
	//
	// example:
	//
	// condition_region_config_cn
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s DescribeCdnConditionIPBInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnConditionIPBInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnConditionIPBInfoRequest) GetDataId() *string {
	return s.DataId
}

func (s *DescribeCdnConditionIPBInfoRequest) SetDataId(v string) *DescribeCdnConditionIPBInfoRequest {
	s.DataId = &v
	return s
}

func (s *DescribeCdnConditionIPBInfoRequest) Validate() error {
	return dara.Validate(s)
}
