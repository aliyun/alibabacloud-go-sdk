// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateComputeQuotaPlanResponseBody
	GetData() *string
	SetErrorCode(v string) *UpdateComputeQuotaPlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateComputeQuotaPlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *UpdateComputeQuotaPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateComputeQuotaPlanResponseBody
	GetRequestId() *string
}

type UpdateComputeQuotaPlanResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComputeQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateComputeQuotaPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateComputeQuotaPlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateComputeQuotaPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateComputeQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeQuotaPlanResponseBody) SetData(v string) *UpdateComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetErrorCode(v string) *UpdateComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetErrorMsg(v string) *UpdateComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetHttpCode(v int32) *UpdateComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) SetRequestId(v string) *UpdateComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
