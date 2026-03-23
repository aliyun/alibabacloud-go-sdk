// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootApResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *RebootApResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *RebootApResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *RebootApResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *RebootApResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *RebootApResponseBody
	GetRequestId() *string
}

type RebootApResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootApResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootApResponseBody) GoString() string {
	return s.String()
}

func (s *RebootApResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *RebootApResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *RebootApResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RebootApResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *RebootApResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootApResponseBody) SetData(v map[string]interface{}) *RebootApResponseBody {
	s.Data = v
	return s
}

func (s *RebootApResponseBody) SetErrorCode(v int32) *RebootApResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RebootApResponseBody) SetErrorMessage(v string) *RebootApResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RebootApResponseBody) SetIsSuccess(v bool) *RebootApResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *RebootApResponseBody) SetRequestId(v string) *RebootApResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootApResponseBody) Validate() error {
	return dara.Validate(s)
}
