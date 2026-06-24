// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateAllRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *DescribeTemplateAllRulesRequest
	GetFeatureType() *int32
	SetLang(v string) *DescribeTemplateAllRulesRequest
	GetLang() *string
	SetTemplateId(v int64) *DescribeTemplateAllRulesRequest
	GetTemplateId() *int64
}

type DescribeTemplateAllRulesRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the industry template.
	//
	// > You can call the [DescribeCategoryTemplateList](https://help.aliyun.com/document_detail/2399296.html) operation to obtain the ID of the industry template. If you do not specify this parameter, the list of models for the primary template is returned by default.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTemplateAllRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateAllRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAllRulesRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeTemplateAllRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTemplateAllRulesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeTemplateAllRulesRequest) SetFeatureType(v int32) *DescribeTemplateAllRulesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeTemplateAllRulesRequest) SetLang(v string) *DescribeTemplateAllRulesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTemplateAllRulesRequest) SetTemplateId(v int64) *DescribeTemplateAllRulesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeTemplateAllRulesRequest) Validate() error {
	return dara.Validate(s)
}
