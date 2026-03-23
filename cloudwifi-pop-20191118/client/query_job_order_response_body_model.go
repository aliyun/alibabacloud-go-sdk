// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *QueryJobOrderResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *QueryJobOrderResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *QueryJobOrderResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *QueryJobOrderResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *QueryJobOrderResponseBody
	GetRequestId() *string
}

type QueryJobOrderResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryJobOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryJobOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobOrderResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryJobOrderResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *QueryJobOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryJobOrderResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *QueryJobOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryJobOrderResponseBody) SetData(v map[string]interface{}) *QueryJobOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryJobOrderResponseBody) SetErrorCode(v int32) *QueryJobOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryJobOrderResponseBody) SetErrorMessage(v string) *QueryJobOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryJobOrderResponseBody) SetIsSuccess(v bool) *QueryJobOrderResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *QueryJobOrderResponseBody) SetRequestId(v string) *QueryJobOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryJobOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
