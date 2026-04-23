// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchInventoryKnowledgeResponseBodyData) *SearchInventoryKnowledgeResponseBody
	GetData() *SearchInventoryKnowledgeResponseBodyData
	SetErrorCode(v string) *SearchInventoryKnowledgeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchInventoryKnowledgeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SearchInventoryKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchInventoryKnowledgeResponseBody
	GetSuccess() *bool
}

type SearchInventoryKnowledgeResponseBody struct {
	Data *SearchInventoryKnowledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SearchInventoryKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInventoryKnowledgeResponseBody) GetData() *SearchInventoryKnowledgeResponseBodyData {
	return s.Data
}

func (s *SearchInventoryKnowledgeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchInventoryKnowledgeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchInventoryKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchInventoryKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchInventoryKnowledgeResponseBody) SetData(v *SearchInventoryKnowledgeResponseBodyData) *SearchInventoryKnowledgeResponseBody {
	s.Data = v
	return s
}

func (s *SearchInventoryKnowledgeResponseBody) SetErrorCode(v string) *SearchInventoryKnowledgeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchInventoryKnowledgeResponseBody) SetErrorMessage(v string) *SearchInventoryKnowledgeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchInventoryKnowledgeResponseBody) SetRequestId(v string) *SearchInventoryKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchInventoryKnowledgeResponseBody) SetSuccess(v bool) *SearchInventoryKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *SearchInventoryKnowledgeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchInventoryKnowledgeResponseBodyData struct {
	Items []*KnowledgeBaseVO `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchInventoryKnowledgeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryKnowledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchInventoryKnowledgeResponseBodyData) GetItems() []*KnowledgeBaseVO {
	return s.Items
}

func (s *SearchInventoryKnowledgeResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchInventoryKnowledgeResponseBodyData) SetItems(v []*KnowledgeBaseVO) *SearchInventoryKnowledgeResponseBodyData {
	s.Items = v
	return s
}

func (s *SearchInventoryKnowledgeResponseBodyData) SetTotalCount(v int32) *SearchInventoryKnowledgeResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SearchInventoryKnowledgeResponseBodyData) Validate() error {
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
