// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApBasicConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *SaveApBasicConfigResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *SaveApBasicConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApBasicConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApBasicConfigResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SaveApBasicConfigResponseBody
	GetRequestId() *string
}

type SaveApBasicConfigResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveApBasicConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApBasicConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApBasicConfigResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SaveApBasicConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApBasicConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApBasicConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApBasicConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveApBasicConfigResponseBody) SetData(v map[string]interface{}) *SaveApBasicConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveApBasicConfigResponseBody) SetErrorCode(v int32) *SaveApBasicConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApBasicConfigResponseBody) SetErrorMessage(v string) *SaveApBasicConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApBasicConfigResponseBody) SetIsSuccess(v bool) *SaveApBasicConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApBasicConfigResponseBody) SetRequestId(v string) *SaveApBasicConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveApBasicConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
