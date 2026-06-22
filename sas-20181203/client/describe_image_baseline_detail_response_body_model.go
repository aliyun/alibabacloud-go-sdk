// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineDetail(v *DescribeImageBaselineDetailResponseBodyBaselineDetail) *DescribeImageBaselineDetailResponseBody
	GetBaselineDetail() *DescribeImageBaselineDetailResponseBodyBaselineDetail
	SetRequestId(v string) *DescribeImageBaselineDetailResponseBody
	GetRequestId() *string
}

type DescribeImageBaselineDetailResponseBody struct {
	// The image baseline details.
	BaselineDetail *DescribeImageBaselineDetailResponseBodyBaselineDetail `json:"BaselineDetail,omitempty" xml:"BaselineDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageBaselineDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineDetailResponseBody) GetBaselineDetail() *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	return s.BaselineDetail
}

func (s *DescribeImageBaselineDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageBaselineDetailResponseBody) SetBaselineDetail(v *DescribeImageBaselineDetailResponseBodyBaselineDetail) *DescribeImageBaselineDetailResponseBody {
	s.BaselineDetail = v
	return s
}

func (s *DescribeImageBaselineDetailResponseBody) SetRequestId(v string) *DescribeImageBaselineDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBody) Validate() error {
	if s.BaselineDetail != nil {
		if err := s.BaselineDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageBaselineDetailResponseBodyBaselineDetail struct {
	// The remediation suggestion for the baseline check risk item.
	//
	// example:
	//
	// 删除泄漏的AK信息。
	Advice *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	// The alias of the baseline classification.
	//
	// example:
	//
	// ak_leak
	BaselineClassAlias *string `json:"BaselineClassAlias,omitempty" xml:"BaselineClassAlias,omitempty"`
	// The alias of the baseline sub-item.
	//
	// example:
	//
	// Access Key泄露
	BaselineItemAlias *string `json:"BaselineItemAlias,omitempty" xml:"BaselineItemAlias,omitempty"`
	// The key of the baseline sub-item.
	//
	// example:
	//
	// ak_leak
	BaselineItemKey *string `json:"BaselineItemKey,omitempty" xml:"BaselineItemKey,omitempty"`
	// The alias of the baseline name.
	//
	// example:
	//
	// Access Key泄露
	BaselineNameAlias *string `json:"BaselineNameAlias,omitempty" xml:"BaselineNameAlias,omitempty"`
	// The description of the baseline risk item.
	//
	// example:
	//
	// Access Key泄露，会导致Access Key被冒用的风险。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The risk level of the check item. Valid values:
	//
	// - **high**: High risk level.
	//
	// - **medium**: Medium risk level.
	//
	// - **low**: Low risk level.
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The baseline issue prompt.
	//
	// example:
	//
	// /usr/aksk.txt:yourAccessKeyID
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The asynchronous request ID.
	//
	// example:
	//
	// async__c6f3b0b54613383b40bdce593ffe****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
}

func (s DescribeImageBaselineDetailResponseBodyBaselineDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineDetailResponseBodyBaselineDetail) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetAdvice() *string {
	return s.Advice
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetBaselineClassAlias() *string {
	return s.BaselineClassAlias
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetBaselineItemAlias() *string {
	return s.BaselineItemAlias
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetBaselineItemKey() *string {
	return s.BaselineItemKey
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetBaselineNameAlias() *string {
	return s.BaselineNameAlias
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetLevel() *string {
	return s.Level
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) GetResultId() *string {
	return s.ResultId
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetAdvice(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.Advice = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetBaselineClassAlias(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.BaselineClassAlias = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetBaselineItemAlias(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.BaselineItemAlias = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetBaselineItemKey(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.BaselineItemKey = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetBaselineNameAlias(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.BaselineNameAlias = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetDescription(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.Description = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetLevel(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.Level = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetPrompt(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.Prompt = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) SetResultId(v string) *DescribeImageBaselineDetailResponseBodyBaselineDetail {
	s.ResultId = &v
	return s
}

func (s *DescribeImageBaselineDetailResponseBodyBaselineDetail) Validate() error {
	return dara.Validate(s)
}
