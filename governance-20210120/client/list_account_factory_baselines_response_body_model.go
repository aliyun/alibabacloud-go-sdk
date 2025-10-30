// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountFactoryBaselinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselines(v []*ListAccountFactoryBaselinesResponseBodyBaselines) *ListAccountFactoryBaselinesResponseBody
	GetBaselines() []*ListAccountFactoryBaselinesResponseBodyBaselines
	SetNextToken(v string) *ListAccountFactoryBaselinesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccountFactoryBaselinesResponseBody
	GetRequestId() *string
}

type ListAccountFactoryBaselinesResponseBody struct {
	// The baselines.
	Baselines []*ListAccountFactoryBaselinesResponseBodyBaselines `json:"Baselines,omitempty" xml:"Baselines,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3245E413-7CDD-5287-8988-6A94DE8A8369
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountFactoryBaselinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesResponseBody) GetBaselines() []*ListAccountFactoryBaselinesResponseBodyBaselines {
	return s.Baselines
}

func (s *ListAccountFactoryBaselinesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccountFactoryBaselinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountFactoryBaselinesResponseBody) SetBaselines(v []*ListAccountFactoryBaselinesResponseBodyBaselines) *ListAccountFactoryBaselinesResponseBody {
	s.Baselines = v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBody) SetNextToken(v string) *ListAccountFactoryBaselinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBody) SetRequestId(v string) *ListAccountFactoryBaselinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBody) Validate() error {
	if s.Baselines != nil {
		for _, item := range s.Baselines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccountFactoryBaselinesResponseBodyBaselines struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Default
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The time at which the baseline was created.
	//
	// example:
	//
	// 2021-11-30T09:09:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the baseline.
	//
	// example:
	//
	// Default baseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the baseline. Valid values:
	//
	// 	- System: default baseline.
	//
	// 	- Custom: custom baseline.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the baseline was updated.
	//
	// example:
	//
	// 2022-12-29T07:08:40Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAccountFactoryBaselinesResponseBodyBaselines) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselinesResponseBodyBaselines) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) GetBaselineId() *string {
	return s.BaselineId
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) GetBaselineName() *string {
	return s.BaselineName
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) GetDescription() *string {
	return s.Description
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) GetType() *string {
	return s.Type
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetBaselineId(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.BaselineId = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetBaselineName(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.BaselineName = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetCreateTime(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.CreateTime = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetDescription(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.Description = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetType(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetUpdateTime(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.UpdateTime = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) Validate() error {
	return dara.Validate(s)
}
