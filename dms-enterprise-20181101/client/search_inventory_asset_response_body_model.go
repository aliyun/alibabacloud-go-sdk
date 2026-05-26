// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInventoryAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchInventoryAssetResponseBodyData) *SearchInventoryAssetResponseBody
	GetData() *SearchInventoryAssetResponseBodyData
	SetErrorCode(v string) *SearchInventoryAssetResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchInventoryAssetResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SearchInventoryAssetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchInventoryAssetResponseBody
	GetSuccess() *bool
}

type SearchInventoryAssetResponseBody struct {
	Data *SearchInventoryAssetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SearchInventoryAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryAssetResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInventoryAssetResponseBody) GetData() *SearchInventoryAssetResponseBodyData {
	return s.Data
}

func (s *SearchInventoryAssetResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchInventoryAssetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchInventoryAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchInventoryAssetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchInventoryAssetResponseBody) SetData(v *SearchInventoryAssetResponseBodyData) *SearchInventoryAssetResponseBody {
	s.Data = v
	return s
}

func (s *SearchInventoryAssetResponseBody) SetErrorCode(v string) *SearchInventoryAssetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchInventoryAssetResponseBody) SetErrorMessage(v string) *SearchInventoryAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchInventoryAssetResponseBody) SetRequestId(v string) *SearchInventoryAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchInventoryAssetResponseBody) SetSuccess(v bool) *SearchInventoryAssetResponseBody {
	s.Success = &v
	return s
}

func (s *SearchInventoryAssetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchInventoryAssetResponseBodyData struct {
	Items []*TableKnowledgeVO `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchInventoryAssetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchInventoryAssetResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchInventoryAssetResponseBodyData) GetItems() []*TableKnowledgeVO {
	return s.Items
}

func (s *SearchInventoryAssetResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchInventoryAssetResponseBodyData) SetItems(v []*TableKnowledgeVO) *SearchInventoryAssetResponseBodyData {
	s.Items = v
	return s
}

func (s *SearchInventoryAssetResponseBodyData) SetTotalCount(v int32) *SearchInventoryAssetResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SearchInventoryAssetResponseBodyData) Validate() error {
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
