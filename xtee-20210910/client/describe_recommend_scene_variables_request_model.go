// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendSceneVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRecommendSceneVariablesRequest
	GetLang() *string
	SetRegId(v string) *DescribeRecommendSceneVariablesRequest
	GetRegId() *string
	SetSampleId(v int64) *DescribeRecommendSceneVariablesRequest
	GetSampleId() *int64
}

type DescribeRecommendSceneVariablesRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region Code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Sample ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 5467
	SampleId *int64 `json:"sampleId,omitempty" xml:"sampleId,omitempty"`
}

func (s DescribeRecommendSceneVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendSceneVariablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendSceneVariablesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecommendSceneVariablesRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRecommendSceneVariablesRequest) GetSampleId() *int64 {
	return s.SampleId
}

func (s *DescribeRecommendSceneVariablesRequest) SetLang(v string) *DescribeRecommendSceneVariablesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecommendSceneVariablesRequest) SetRegId(v string) *DescribeRecommendSceneVariablesRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRecommendSceneVariablesRequest) SetSampleId(v int64) *DescribeRecommendSceneVariablesRequest {
	s.SampleId = &v
	return s
}

func (s *DescribeRecommendSceneVariablesRequest) Validate() error {
	return dara.Validate(s)
}
