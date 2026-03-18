// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeSubQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateComputeSubQuotaResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateComputeSubQuotaResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateComputeSubQuotaResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateComputeSubQuotaResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateComputeSubQuotaResponseBody
	GetRequestId() *string
}

type UpdateComputeSubQuotaResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComputeSubQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeSubQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeSubQuotaResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateComputeSubQuotaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateComputeSubQuotaResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateComputeSubQuotaResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateComputeSubQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeSubQuotaResponseBody) SetData(v string) *UpdateComputeSubQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetErrorCode(v string) *UpdateComputeSubQuotaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetErrorMsg(v string) *UpdateComputeSubQuotaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetHttpCode(v int32) *UpdateComputeSubQuotaResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) SetRequestId(v string) *UpdateComputeSubQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeSubQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
