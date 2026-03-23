// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagMetaAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTagMetaAssetResponseBodyData) *ListTagMetaAssetResponseBody
	GetData() []*ListTagMetaAssetResponseBodyData
	SetErrorCode(v string) *ListTagMetaAssetResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTagMetaAssetResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTagMetaAssetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTagMetaAssetResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTagMetaAssetResponseBody
	GetTotalCount() *int64
}

type ListTagMetaAssetResponseBody struct {
	Data []*ListTagMetaAssetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagMetaAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagMetaAssetResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagMetaAssetResponseBody) GetData() []*ListTagMetaAssetResponseBodyData {
	return s.Data
}

func (s *ListTagMetaAssetResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTagMetaAssetResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTagMetaAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagMetaAssetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTagMetaAssetResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTagMetaAssetResponseBody) SetData(v []*ListTagMetaAssetResponseBodyData) *ListTagMetaAssetResponseBody {
	s.Data = v
	return s
}

func (s *ListTagMetaAssetResponseBody) SetErrorCode(v string) *ListTagMetaAssetResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTagMetaAssetResponseBody) SetErrorMessage(v string) *ListTagMetaAssetResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTagMetaAssetResponseBody) SetRequestId(v string) *ListTagMetaAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagMetaAssetResponseBody) SetSuccess(v bool) *ListTagMetaAssetResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagMetaAssetResponseBody) SetTotalCount(v int64) *ListTagMetaAssetResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagMetaAssetResponseBody) Validate() error {
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

type ListTagMetaAssetResponseBodyData struct {
	// example:
	//
	// {"schemaName": "test"}
	MetaEntityAttrs map[string]interface{} `json:"MetaEntityAttrs,omitempty" xml:"MetaEntityAttrs,omitempty"`
	// example:
	//
	// META_DATABASE
	MetaType *string `json:"MetaType,omitempty" xml:"MetaType,omitempty"`
}

func (s ListTagMetaAssetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTagMetaAssetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagMetaAssetResponseBodyData) GetMetaEntityAttrs() map[string]interface{} {
	return s.MetaEntityAttrs
}

func (s *ListTagMetaAssetResponseBodyData) GetMetaType() *string {
	return s.MetaType
}

func (s *ListTagMetaAssetResponseBodyData) SetMetaEntityAttrs(v map[string]interface{}) *ListTagMetaAssetResponseBodyData {
	s.MetaEntityAttrs = v
	return s
}

func (s *ListTagMetaAssetResponseBodyData) SetMetaType(v string) *ListTagMetaAssetResponseBodyData {
	s.MetaType = &v
	return s
}

func (s *ListTagMetaAssetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
