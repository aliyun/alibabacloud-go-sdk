// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskLevelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DescribeRiskLevelsRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeRiskLevelsRequest
	GetLang() *string
	SetTemplateId(v int64) *DescribeRiskLevelsRequest
	GetTemplateId() *int64
}

type DescribeRiskLevelsRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- zh_cn: Chinese (default)
	//
	// 	- en_us: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the industry-specific rule template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeRiskLevelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLevelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskLevelsRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeRiskLevelsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskLevelsRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeRiskLevelsRequest) SetFeatureType(v int32) *DescribeRiskLevelsRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeRiskLevelsRequest) SetLang(v string) *DescribeRiskLevelsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskLevelsRequest) SetTemplateId(v int64) *DescribeRiskLevelsRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeRiskLevelsRequest) Validate() error {
	return dara.Validate(s)
}
