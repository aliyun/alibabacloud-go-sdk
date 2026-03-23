// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApDetailedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApDetailedConfigResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApDetailedConfigResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApDetailedConfigResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApDetailedConfigResponseBody
	GetIsSuccess() *bool
}

type GetApDetailedConfigResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApDetailedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApDetailedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApDetailedConfigResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApDetailedConfigResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApDetailedConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApDetailedConfigResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApDetailedConfigResponseBody) SetData(v map[string]interface{}) *GetApDetailedConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetApDetailedConfigResponseBody) SetErrorCode(v int32) *GetApDetailedConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApDetailedConfigResponseBody) SetErrorMessage(v string) *GetApDetailedConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApDetailedConfigResponseBody) SetIsSuccess(v bool) *GetApDetailedConfigResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApDetailedConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
