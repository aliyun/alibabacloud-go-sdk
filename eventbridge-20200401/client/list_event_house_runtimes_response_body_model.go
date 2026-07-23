// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventHouseRuntimesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEventHouseRuntimesResponseBody
	GetCode() *string
	SetData(v *ListEventHouseRuntimesResponseBodyData) *ListEventHouseRuntimesResponseBody
	GetData() *ListEventHouseRuntimesResponseBodyData
	SetMessage(v string) *ListEventHouseRuntimesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEventHouseRuntimesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEventHouseRuntimesResponseBody
	GetSuccess() *bool
}

type ListEventHouseRuntimesResponseBody struct {
	// The response code. Success indicates that the operation was successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The runtime list result.
	Data *ListEventHouseRuntimesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEventHouseRuntimesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventHouseRuntimesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventHouseRuntimesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEventHouseRuntimesResponseBody) GetData() *ListEventHouseRuntimesResponseBodyData {
	return s.Data
}

func (s *ListEventHouseRuntimesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEventHouseRuntimesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventHouseRuntimesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEventHouseRuntimesResponseBody) SetCode(v string) *ListEventHouseRuntimesResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventHouseRuntimesResponseBody) SetData(v *ListEventHouseRuntimesResponseBodyData) *ListEventHouseRuntimesResponseBody {
	s.Data = v
	return s
}

func (s *ListEventHouseRuntimesResponseBody) SetMessage(v string) *ListEventHouseRuntimesResponseBody {
	s.Message = &v
	return s
}

func (s *ListEventHouseRuntimesResponseBody) SetRequestId(v string) *ListEventHouseRuntimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventHouseRuntimesResponseBody) SetSuccess(v bool) *ListEventHouseRuntimesResponseBody {
	s.Success = &v
	return s
}

func (s *ListEventHouseRuntimesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventHouseRuntimesResponseBodyData struct {
	// example:
	//
	// 20
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of runtimes.
	Runtimes []*EventHouseRuntime `json:"Runtimes,omitempty" xml:"Runtimes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEventHouseRuntimesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEventHouseRuntimesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventHouseRuntimesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventHouseRuntimesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventHouseRuntimesResponseBodyData) GetRuntimes() []*EventHouseRuntime {
	return s.Runtimes
}

func (s *ListEventHouseRuntimesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEventHouseRuntimesResponseBodyData) SetMaxResults(v int32) *ListEventHouseRuntimesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListEventHouseRuntimesResponseBodyData) SetNextToken(v string) *ListEventHouseRuntimesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListEventHouseRuntimesResponseBodyData) SetRuntimes(v []*EventHouseRuntime) *ListEventHouseRuntimesResponseBodyData {
	s.Runtimes = v
	return s
}

func (s *ListEventHouseRuntimesResponseBodyData) SetTotalCount(v int32) *ListEventHouseRuntimesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListEventHouseRuntimesResponseBodyData) Validate() error {
	if s.Runtimes != nil {
		for _, item := range s.Runtimes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
