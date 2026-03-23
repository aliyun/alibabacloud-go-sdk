// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageVisitDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *GetPageVisitDataResponseBody
	GetData() map[string]interface{}
	SetErrorCode(v int32) *GetPageVisitDataResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *GetPageVisitDataResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *GetPageVisitDataResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetPageVisitDataResponseBody
	GetRequestId() *string
}

type GetPageVisitDataResponseBody struct {
	// example:
	//
	// {
	//
	//       "totalPv": 535,
	//
	//       "totalUv": 246,
	//
	//       "detailVo": [
	//
	// {
	//
	// "uv":17,
	//
	// "pv":56,
	//
	// "ds":"20230710"
	//
	// }
	//
	// ]
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
	// example:
	//
	// 76F569F1-078E-5A08-881D-810B97C502A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPageVisitDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPageVisitDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageVisitDataResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetPageVisitDataResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *GetPageVisitDataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPageVisitDataResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetPageVisitDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPageVisitDataResponseBody) SetData(v map[string]interface{}) *GetPageVisitDataResponseBody {
	s.Data = v
	return s
}

func (s *GetPageVisitDataResponseBody) SetErrorCode(v int32) *GetPageVisitDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPageVisitDataResponseBody) SetErrorMessage(v string) *GetPageVisitDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPageVisitDataResponseBody) SetIsSuccess(v bool) *GetPageVisitDataResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetPageVisitDataResponseBody) SetRequestId(v string) *GetPageVisitDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageVisitDataResponseBody) Validate() error {
	return dara.Validate(s)
}
