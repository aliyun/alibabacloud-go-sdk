// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTableAssetKnowledgeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchTableAssetKnowledgeResponseBodyData) *SearchTableAssetKnowledgeResponseBody
	GetData() *SearchTableAssetKnowledgeResponseBodyData
	SetErrorCode(v string) *SearchTableAssetKnowledgeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchTableAssetKnowledgeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SearchTableAssetKnowledgeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchTableAssetKnowledgeResponseBody
	GetSuccess() *bool
}

type SearchTableAssetKnowledgeResponseBody struct {
	Data *SearchTableAssetKnowledgeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SearchTableAssetKnowledgeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTableAssetKnowledgeResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTableAssetKnowledgeResponseBody) GetData() *SearchTableAssetKnowledgeResponseBodyData {
	return s.Data
}

func (s *SearchTableAssetKnowledgeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchTableAssetKnowledgeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchTableAssetKnowledgeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTableAssetKnowledgeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchTableAssetKnowledgeResponseBody) SetData(v *SearchTableAssetKnowledgeResponseBodyData) *SearchTableAssetKnowledgeResponseBody {
	s.Data = v
	return s
}

func (s *SearchTableAssetKnowledgeResponseBody) SetErrorCode(v string) *SearchTableAssetKnowledgeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchTableAssetKnowledgeResponseBody) SetErrorMessage(v string) *SearchTableAssetKnowledgeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchTableAssetKnowledgeResponseBody) SetRequestId(v string) *SearchTableAssetKnowledgeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTableAssetKnowledgeResponseBody) SetSuccess(v bool) *SearchTableAssetKnowledgeResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTableAssetKnowledgeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchTableAssetKnowledgeResponseBodyData struct {
	Items []*KnowledgeBaseVO `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchTableAssetKnowledgeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchTableAssetKnowledgeResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTableAssetKnowledgeResponseBodyData) GetItems() []*KnowledgeBaseVO {
	return s.Items
}

func (s *SearchTableAssetKnowledgeResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchTableAssetKnowledgeResponseBodyData) SetItems(v []*KnowledgeBaseVO) *SearchTableAssetKnowledgeResponseBodyData {
	s.Items = v
	return s
}

func (s *SearchTableAssetKnowledgeResponseBodyData) SetTotalCount(v int32) *SearchTableAssetKnowledgeResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SearchTableAssetKnowledgeResponseBodyData) Validate() error {
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
