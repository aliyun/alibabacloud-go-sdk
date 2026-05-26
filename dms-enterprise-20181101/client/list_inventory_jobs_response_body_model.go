// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventoryJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListInventoryJobsResponseBodyData) *ListInventoryJobsResponseBody
	GetData() *ListInventoryJobsResponseBodyData
	SetErrorCode(v string) *ListInventoryJobsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListInventoryJobsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListInventoryJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInventoryJobsResponseBody
	GetSuccess() *bool
}

type ListInventoryJobsResponseBody struct {
	Data *ListInventoryJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInventoryJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInventoryJobsResponseBody) GetData() *ListInventoryJobsResponseBodyData {
	return s.Data
}

func (s *ListInventoryJobsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInventoryJobsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInventoryJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInventoryJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInventoryJobsResponseBody) SetData(v *ListInventoryJobsResponseBodyData) *ListInventoryJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListInventoryJobsResponseBody) SetErrorCode(v string) *ListInventoryJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInventoryJobsResponseBody) SetErrorMessage(v string) *ListInventoryJobsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInventoryJobsResponseBody) SetRequestId(v string) *ListInventoryJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInventoryJobsResponseBody) SetSuccess(v bool) *ListInventoryJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListInventoryJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInventoryJobsResponseBodyData struct {
	Items []*KnowledgeJobInfoVO `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInventoryJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInventoryJobsResponseBodyData) GetItems() []*KnowledgeJobInfoVO {
	return s.Items
}

func (s *ListInventoryJobsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInventoryJobsResponseBodyData) SetItems(v []*KnowledgeJobInfoVO) *ListInventoryJobsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListInventoryJobsResponseBodyData) SetTotalCount(v int32) *ListInventoryJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInventoryJobsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
