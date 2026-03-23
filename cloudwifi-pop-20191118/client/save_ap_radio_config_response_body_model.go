// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApRadioConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *SaveApRadioConfigResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *SaveApRadioConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApRadioConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApRadioConfigResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *SaveApRadioConfigResponseBody
	GetRequestId() *string
}

type SaveApRadioConfigResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveApRadioConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApRadioConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApRadioConfigResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SaveApRadioConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApRadioConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApRadioConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApRadioConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveApRadioConfigResponseBody) SetData(v map[string]interface{}) *SaveApRadioConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveApRadioConfigResponseBody) SetErrorCode(v int32) *SaveApRadioConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApRadioConfigResponseBody) SetErrorMessage(v string) *SaveApRadioConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApRadioConfigResponseBody) SetIsSuccess(v bool) *SaveApRadioConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApRadioConfigResponseBody) SetRequestId(v string) *SaveApRadioConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveApRadioConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
