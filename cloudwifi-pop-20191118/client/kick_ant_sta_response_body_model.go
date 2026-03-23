// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickAntStaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *KickAntStaResponseBody
	GetData() *string
	SetErrorCode(v int32) *KickAntStaResponseBody
	GetErrorCode() *int32
	SetErrorMessage(v string) *KickAntStaResponseBody
	GetErrorMessage() *string
	SetIsSuccess(v bool) *KickAntStaResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *KickAntStaResponseBody
	GetRequestId() *string
}

type KickAntStaResponseBody struct {
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
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 0
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

func (s KickAntStaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KickAntStaResponseBody) GoString() string {
	return s.String()
}

func (s *KickAntStaResponseBody) GetData() *string {
	return s.Data
}

func (s *KickAntStaResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *KickAntStaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *KickAntStaResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *KickAntStaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KickAntStaResponseBody) SetData(v string) *KickAntStaResponseBody {
	s.Data = &v
	return s
}

func (s *KickAntStaResponseBody) SetErrorCode(v int32) *KickAntStaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *KickAntStaResponseBody) SetErrorMessage(v string) *KickAntStaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *KickAntStaResponseBody) SetIsSuccess(v bool) *KickAntStaResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *KickAntStaResponseBody) SetRequestId(v string) *KickAntStaResponseBody {
	s.RequestId = &v
	return s
}

func (s *KickAntStaResponseBody) Validate() error {
	return dara.Validate(s)
}
