// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *BaseResponse
	GetData() map[string]interface{}
	SetErrorCode(v string) *BaseResponse
	GetErrorCode() *string
	SetErrorMessage(v string) *BaseResponse
	GetErrorMessage() *string
}

type BaseResponse struct {
	Data         map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode    *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string                `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}

func (s BaseResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseResponse) GoString() string {
	return s.String()
}

func (s *BaseResponse) GetData() map[string]interface{} {
	return s.Data
}

func (s *BaseResponse) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BaseResponse) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BaseResponse) SetData(v map[string]interface{}) *BaseResponse {
	s.Data = v
	return s
}

func (s *BaseResponse) SetErrorCode(v string) *BaseResponse {
	s.ErrorCode = &v
	return s
}

func (s *BaseResponse) SetErrorMessage(v string) *BaseResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BaseResponse) Validate() error {
	return dara.Validate(s)
}
