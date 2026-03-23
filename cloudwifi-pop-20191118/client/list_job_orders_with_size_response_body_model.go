// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobOrdersWithSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListJobOrdersWithSizeResponseBodyData) *ListJobOrdersWithSizeResponseBody
	GetData() *ListJobOrdersWithSizeResponseBodyData
	SetErrorCode(v int32) *ListJobOrdersWithSizeResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *ListJobOrdersWithSizeResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *ListJobOrdersWithSizeResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ListJobOrdersWithSizeResponseBody
	GetRequestId() *string
}

type ListJobOrdersWithSizeResponseBody struct {
	Data         *ListJobOrdersWithSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *int32                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobOrdersWithSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobOrdersWithSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobOrdersWithSizeResponseBody) GetData() *ListJobOrdersWithSizeResponseBodyData {
	return s.Data
}

func (s *ListJobOrdersWithSizeResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *ListJobOrdersWithSizeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListJobOrdersWithSizeResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListJobOrdersWithSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobOrdersWithSizeResponseBody) SetData(v *ListJobOrdersWithSizeResponseBodyData) *ListJobOrdersWithSizeResponseBody {
	s.Data = v
	return s
}

func (s *ListJobOrdersWithSizeResponseBody) SetErrorCode(v int32) *ListJobOrdersWithSizeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobOrdersWithSizeResponseBody) SetErrorMessage(v string) *ListJobOrdersWithSizeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListJobOrdersWithSizeResponseBody) SetIsSuccess(v bool) *ListJobOrdersWithSizeResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListJobOrdersWithSizeResponseBody) SetRequestId(v string) *ListJobOrdersWithSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobOrdersWithSizeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobOrdersWithSizeResponseBodyData struct {
	Count *int64                   `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListJobOrdersWithSizeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobOrdersWithSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobOrdersWithSizeResponseBodyData) GetCount() *int64 {
	return s.Count
}

func (s *ListJobOrdersWithSizeResponseBodyData) GetData() []map[string]interface{} {
	return s.Data
}

func (s *ListJobOrdersWithSizeResponseBodyData) SetCount(v int64) *ListJobOrdersWithSizeResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListJobOrdersWithSizeResponseBodyData) SetData(v []map[string]interface{}) *ListJobOrdersWithSizeResponseBodyData {
	s.Data = v
	return s
}

func (s *ListJobOrdersWithSizeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
