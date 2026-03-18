// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeQuotaPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteComputeQuotaPlanResponseBody
	GetData() *string
	SetErrorCode(v string) *DeleteComputeQuotaPlanResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *DeleteComputeQuotaPlanResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *DeleteComputeQuotaPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteComputeQuotaPlanResponseBody
	GetRequestId() *string
}

type DeleteComputeQuotaPlanResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32  `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteComputeQuotaPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComputeQuotaPlanResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteComputeQuotaPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteComputeQuotaPlanResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteComputeQuotaPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteComputeQuotaPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComputeQuotaPlanResponseBody) SetData(v string) *DeleteComputeQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetErrorCode(v string) *DeleteComputeQuotaPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetErrorMsg(v string) *DeleteComputeQuotaPlanResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetHttpCode(v int32) *DeleteComputeQuotaPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) SetRequestId(v string) *DeleteComputeQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComputeQuotaPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
