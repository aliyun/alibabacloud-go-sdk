// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []map[string]interface{}) *ListJobOrdersResponseBody
	GetData() []map[string]interface{}
	SetErrorCode(v int32) *ListJobOrdersResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *ListJobOrdersResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *ListJobOrdersResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *ListJobOrdersResponseBody
	GetRequestId() *string
}

type ListJobOrdersResponseBody struct {
	Data         []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *int32                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobOrdersResponseBody) GetData() []map[string]interface{} {
	return s.Data
}

func (s *ListJobOrdersResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *ListJobOrdersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListJobOrdersResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListJobOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobOrdersResponseBody) SetData(v []map[string]interface{}) *ListJobOrdersResponseBody {
	s.Data = v
	return s
}

func (s *ListJobOrdersResponseBody) SetErrorCode(v int32) *ListJobOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobOrdersResponseBody) SetErrorMessage(v string) *ListJobOrdersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListJobOrdersResponseBody) SetIsSuccess(v bool) *ListJobOrdersResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListJobOrdersResponseBody) SetRequestId(v string) *ListJobOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobOrdersResponseBody) Validate() error {
	return dara.Validate(s)
}
