// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartUnionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *StartUnionTaskResponseBody
	GetCode() *int64
	SetData(v bool) *StartUnionTaskResponseBody
	GetData() *bool
	SetErrorMsg(v string) *StartUnionTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *StartUnionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartUnionTaskResponseBody
	GetSuccess() *bool
}

type StartUnionTaskResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartUnionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartUnionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartUnionTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *StartUnionTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StartUnionTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *StartUnionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartUnionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartUnionTaskResponseBody) SetCode(v int64) *StartUnionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartUnionTaskResponseBody) SetData(v bool) *StartUnionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartUnionTaskResponseBody) SetErrorMsg(v string) *StartUnionTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StartUnionTaskResponseBody) SetRequestId(v string) *StartUnionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartUnionTaskResponseBody) SetSuccess(v bool) *StartUnionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StartUnionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
