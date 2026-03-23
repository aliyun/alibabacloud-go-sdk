// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStaDetailedStatusByMacResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetStaDetailedStatusByMacResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetStaDetailedStatusByMacResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetStaDetailedStatusByMacResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetStaDetailedStatusByMacResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetStaDetailedStatusByMacResponseBody
	GetRequestId() *string
}

type GetStaDetailedStatusByMacResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *int32                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	IsSuccess    *bool                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetStaDetailedStatusByMacResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStaDetailedStatusByMacResponseBody) GoString() string {
	return s.String()
}

func (s *GetStaDetailedStatusByMacResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetStaDetailedStatusByMacResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetStaDetailedStatusByMacResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetStaDetailedStatusByMacResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetStaDetailedStatusByMacResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStaDetailedStatusByMacResponseBody) SetData(v map[string]interface{}) *GetStaDetailedStatusByMacResponseBody {
	s.Data = v
	return s
}

func (s *GetStaDetailedStatusByMacResponseBody) SetErrorCode(v int32) *GetStaDetailedStatusByMacResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStaDetailedStatusByMacResponseBody) SetErrorMessage(v string) *GetStaDetailedStatusByMacResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStaDetailedStatusByMacResponseBody) SetIsSuccess(v bool) *GetStaDetailedStatusByMacResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetStaDetailedStatusByMacResponseBody) SetRequestId(v string) *GetStaDetailedStatusByMacResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStaDetailedStatusByMacResponseBody) Validate() error {
	return dara.Validate(s)
}
