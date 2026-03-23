// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStaStatusListByApResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetStaStatusListByApResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetStaStatusListByApResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetStaStatusListByApResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetStaStatusListByApResponseBody
	GetIsSuccess() *bool
}

type GetStaStatusListByApResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
}

func (s GetStaStatusListByApResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStaStatusListByApResponseBody) GoString() string {
	return s.String()
}

func (s *GetStaStatusListByApResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetStaStatusListByApResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetStaStatusListByApResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetStaStatusListByApResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetStaStatusListByApResponseBody) SetData(v map[string]interface{}) *GetStaStatusListByApResponseBody {
	s.Data = v
	return s
}

func (s *GetStaStatusListByApResponseBody) SetErrorCode(v int32) *GetStaStatusListByApResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStaStatusListByApResponseBody) SetErrorMessage(v string) *GetStaStatusListByApResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStaStatusListByApResponseBody) SetIsSuccess(v bool) *GetStaStatusListByApResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetStaStatusListByApResponseBody) Validate() error {
	return dara.Validate(s)
}
