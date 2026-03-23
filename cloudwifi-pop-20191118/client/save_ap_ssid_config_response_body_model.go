// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApSsidConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *SaveApSsidConfigResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *SaveApSsidConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *SaveApSsidConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *SaveApSsidConfigResponseBody
	GetIsSuccess() *bool
}

type SaveApSsidConfigResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s SaveApSsidConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveApSsidConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveApSsidConfigResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SaveApSsidConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *SaveApSsidConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveApSsidConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *SaveApSsidConfigResponseBody) SetData(v map[string]interface{}) *SaveApSsidConfigResponseBody {
	s.Data = v
	return s
}

func (s *SaveApSsidConfigResponseBody) SetErrorCode(v int32) *SaveApSsidConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveApSsidConfigResponseBody) SetErrorMessage(v string) *SaveApSsidConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveApSsidConfigResponseBody) SetIsSuccess(v bool) *SaveApSsidConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SaveApSsidConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
