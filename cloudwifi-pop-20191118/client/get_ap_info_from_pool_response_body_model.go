// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApInfoFromPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetApInfoFromPoolResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetApInfoFromPoolResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetApInfoFromPoolResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetApInfoFromPoolResponseBody
	GetIsSuccess() *bool
}

type GetApInfoFromPoolResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetApInfoFromPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApInfoFromPoolResponseBody) GoString() string {
	return s.String()
}

func (s *GetApInfoFromPoolResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetApInfoFromPoolResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetApInfoFromPoolResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApInfoFromPoolResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetApInfoFromPoolResponseBody) SetData(v map[string]interface{}) *GetApInfoFromPoolResponseBody {
	s.Data = v
	return s
}

func (s *GetApInfoFromPoolResponseBody) SetErrorCode(v int32) *GetApInfoFromPoolResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetApInfoFromPoolResponseBody) SetErrorMessage(v string) *GetApInfoFromPoolResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetApInfoFromPoolResponseBody) SetIsSuccess(v bool) *GetApInfoFromPoolResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetApInfoFromPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
