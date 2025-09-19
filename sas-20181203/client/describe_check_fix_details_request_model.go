// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckFixDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v string) *DescribeCheckFixDetailsRequest
	GetCheckIds() *string
	SetLang(v string) *DescribeCheckFixDetailsRequest
	GetLang() *string
	SetRiskId(v int64) *DescribeCheckFixDetailsRequest
	GetRiskId() *int64
}

type DescribeCheckFixDetailsRequest struct {
	// The ID of the risk item.
	//
	// >  You can call the [DescribeRiskType](~~DescribeRiskType~~) operation to query the IDs of risk items.
	//
	// example:
	//
	// 58
	CheckIds *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the baseline.
	//
	// >  You can call the [DescribeCheckWarningSummary](https://help.aliyun.com/document_detail/116179.html) operation to query the IDs of baselines.
	//
	// example:
	//
	// 51
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
}

func (s DescribeCheckFixDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckFixDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckFixDetailsRequest) GetCheckIds() *string {
	return s.CheckIds
}

func (s *DescribeCheckFixDetailsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCheckFixDetailsRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeCheckFixDetailsRequest) SetCheckIds(v string) *DescribeCheckFixDetailsRequest {
	s.CheckIds = &v
	return s
}

func (s *DescribeCheckFixDetailsRequest) SetLang(v string) *DescribeCheckFixDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckFixDetailsRequest) SetRiskId(v int64) *DescribeCheckFixDetailsRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeCheckFixDetailsRequest) Validate() error {
	return dara.Validate(s)
}
