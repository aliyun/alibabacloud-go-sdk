// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApThirdAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *SaveApThirdAppResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *SaveApThirdAppResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApThirdAppResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApThirdAppResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SaveApThirdAppResponseBody
	GetRequestId() *string
}

type SaveApThirdAppResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveApThirdAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApThirdAppResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApThirdAppResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SaveApThirdAppResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApThirdAppResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApThirdAppResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApThirdAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveApThirdAppResponseBody) SetData(v map[string]interface{}) *SaveApThirdAppResponseBody {
	s.Data = v
	return s
}

func (s *SaveApThirdAppResponseBody) SetErrorCode(v int32) *SaveApThirdAppResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApThirdAppResponseBody) SetErrorMessage(v string) *SaveApThirdAppResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApThirdAppResponseBody) SetIsSuccess(v bool) *SaveApThirdAppResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApThirdAppResponseBody) SetRequestId(v string) *SaveApThirdAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveApThirdAppResponseBody) Validate() error {
	return dara.Validate(s)
}
