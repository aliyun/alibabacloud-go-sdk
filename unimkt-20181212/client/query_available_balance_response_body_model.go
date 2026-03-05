// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryAvailableBalanceResponseBody
	GetCode() *int32
	SetData(v int64) *QueryAvailableBalanceResponseBody
	GetData() *int64
	SetErrorMsg(v string) *QueryAvailableBalanceResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryAvailableBalanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAvailableBalanceResponseBody
	GetSuccess() *bool
}

type QueryAvailableBalanceResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvailableBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailableBalanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryAvailableBalanceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *QueryAvailableBalanceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryAvailableBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAvailableBalanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAvailableBalanceResponseBody) SetCode(v int32) *QueryAvailableBalanceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAvailableBalanceResponseBody) SetData(v int64) *QueryAvailableBalanceResponseBody {
	s.Data = &v
	return s
}

func (s *QueryAvailableBalanceResponseBody) SetErrorMsg(v string) *QueryAvailableBalanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryAvailableBalanceResponseBody) SetRequestId(v string) *QueryAvailableBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailableBalanceResponseBody) SetSuccess(v bool) *QueryAvailableBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAvailableBalanceResponseBody) Validate() error {
	return dara.Validate(s)
}
