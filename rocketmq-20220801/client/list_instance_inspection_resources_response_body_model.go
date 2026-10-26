// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListInstanceInspectionResourcesResponseBodyData) *ListInstanceInspectionResourcesResponseBody
	GetData() []*ListInstanceInspectionResourcesResponseBodyData
	SetMaxResults(v int32) *ListInstanceInspectionResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceInspectionResourcesResponseBody
	GetNextToken() *string
	SetTotal(v int64) *ListInstanceInspectionResourcesResponseBody
	GetTotal() *int64
}

type ListInstanceInspectionResourcesResponseBody struct {
	// example:
	//
	// []
	Data []*ListInstanceInspectionResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 5
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstanceInspectionResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionResourcesResponseBody) GetData() []*ListInstanceInspectionResourcesResponseBodyData {
	return s.Data
}

func (s *ListInstanceInspectionResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInspectionResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInspectionResourcesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListInstanceInspectionResourcesResponseBody) SetData(v []*ListInstanceInspectionResourcesResponseBodyData) *ListInstanceInspectionResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBody) SetMaxResults(v int32) *ListInstanceInspectionResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBody) SetNextToken(v string) *ListInstanceInspectionResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBody) SetTotal(v int64) *ListInstanceInspectionResourcesResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceInspectionResourcesResponseBodyData struct {
	// example:
	//
	// orders-g1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 订单实例
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Risk
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 消费延迟超过阈值
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// ConsumerGroup
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListInstanceInspectionResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionResourcesResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListInstanceInspectionResourcesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListInstanceInspectionResourcesResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *ListInstanceInspectionResourcesResponseBodyData) GetSummary() *string {
	return s.Summary
}

func (s *ListInstanceInspectionResourcesResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListInstanceInspectionResourcesResponseBodyData) SetId(v string) *ListInstanceInspectionResourcesResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBodyData) SetName(v string) *ListInstanceInspectionResourcesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBodyData) SetResult(v string) *ListInstanceInspectionResourcesResponseBodyData {
	s.Result = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBodyData) SetSummary(v string) *ListInstanceInspectionResourcesResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBodyData) SetType(v string) *ListInstanceInspectionResourcesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListInstanceInspectionResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
