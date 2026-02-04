// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionTrackEventTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventTypes(v []*ListActionTrackEventTypesResponseBodyEventTypes) *ListActionTrackEventTypesResponseBody
	GetEventTypes() []*ListActionTrackEventTypesResponseBodyEventTypes
	SetMaxResults(v int64) *ListActionTrackEventTypesResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListActionTrackEventTypesResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListActionTrackEventTypesResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListActionTrackEventTypesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListActionTrackEventTypesResponseBody
	GetTotalCount() *int64
}

type ListActionTrackEventTypesResponseBody struct {
	EventTypes []*ListActionTrackEventTypesResponseBodyEventTypes `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// 分页查询时每页行数。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于上一次翻页查询。
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListActionTrackEventTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListActionTrackEventTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionTrackEventTypesResponseBody) GetEventTypes() []*ListActionTrackEventTypesResponseBodyEventTypes {
	return s.EventTypes
}

func (s *ListActionTrackEventTypesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListActionTrackEventTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionTrackEventTypesResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListActionTrackEventTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActionTrackEventTypesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListActionTrackEventTypesResponseBody) SetEventTypes(v []*ListActionTrackEventTypesResponseBodyEventTypes) *ListActionTrackEventTypesResponseBody {
	s.EventTypes = v
	return s
}

func (s *ListActionTrackEventTypesResponseBody) SetMaxResults(v int64) *ListActionTrackEventTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActionTrackEventTypesResponseBody) SetNextToken(v string) *ListActionTrackEventTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListActionTrackEventTypesResponseBody) SetPreviousToken(v string) *ListActionTrackEventTypesResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListActionTrackEventTypesResponseBody) SetRequestId(v string) *ListActionTrackEventTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActionTrackEventTypesResponseBody) SetTotalCount(v int64) *ListActionTrackEventTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListActionTrackEventTypesResponseBody) Validate() error {
	if s.EventTypes != nil {
		for _, item := range s.EventTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActionTrackEventTypesResponseBodyEventTypes struct {
	// example:
	//
	// urn:alibaba:idaas:event:user:create
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
}

func (s ListActionTrackEventTypesResponseBodyEventTypes) String() string {
	return dara.Prettify(s)
}

func (s ListActionTrackEventTypesResponseBodyEventTypes) GoString() string {
	return s.String()
}

func (s *ListActionTrackEventTypesResponseBodyEventTypes) GetEventType() *string {
	return s.EventType
}

func (s *ListActionTrackEventTypesResponseBodyEventTypes) SetEventType(v string) *ListActionTrackEventTypesResponseBodyEventTypes {
	s.EventType = &v
	return s
}

func (s *ListActionTrackEventTypesResponseBodyEventTypes) Validate() error {
	return dara.Validate(s)
}
