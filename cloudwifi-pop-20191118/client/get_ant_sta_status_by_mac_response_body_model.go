// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAntStaStatusByMacResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetAntStaStatusByMacResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetAntStaStatusByMacResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetAntStaStatusByMacResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetAntStaStatusByMacResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetAntStaStatusByMacResponseBody
	GetRequestId() *string
}

type GetAntStaStatusByMacResponseBody struct {
	// example:
	//
	// {
	//
	// "status": 1
	//
	// }
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 111
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// the status of ap is not online
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 76F569F1-078E-5A08-881D-810B97C502A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAntStaStatusByMacResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAntStaStatusByMacResponseBody) GoString() string {
	return s.String()
}

func (s *GetAntStaStatusByMacResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetAntStaStatusByMacResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetAntStaStatusByMacResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAntStaStatusByMacResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetAntStaStatusByMacResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAntStaStatusByMacResponseBody) SetData(v map[string]interface{}) *GetAntStaStatusByMacResponseBody {
	s.Data = v
	return s
}

func (s *GetAntStaStatusByMacResponseBody) SetErrorCode(v int32) *GetAntStaStatusByMacResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAntStaStatusByMacResponseBody) SetErrorMessage(v string) *GetAntStaStatusByMacResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAntStaStatusByMacResponseBody) SetIsSuccess(v bool) *GetAntStaStatusByMacResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetAntStaStatusByMacResponseBody) SetRequestId(v string) *GetAntStaStatusByMacResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAntStaStatusByMacResponseBody) Validate() error {
	return dara.Validate(s)
}
