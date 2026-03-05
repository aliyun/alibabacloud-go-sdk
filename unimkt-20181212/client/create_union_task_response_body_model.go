// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUnionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateUnionTaskResponseBody
	GetCode() *int32
	SetData(v int64) *CreateUnionTaskResponseBody
	GetData() *int64
	SetErrorMsg(v string) *CreateUnionTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateUnionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUnionTaskResponseBody
	GetSuccess() *bool
}

type CreateUnionTaskResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUnionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUnionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUnionTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateUnionTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateUnionTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateUnionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUnionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUnionTaskResponseBody) SetCode(v int32) *CreateUnionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUnionTaskResponseBody) SetData(v int64) *CreateUnionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateUnionTaskResponseBody) SetErrorMsg(v string) *CreateUnionTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateUnionTaskResponseBody) SetRequestId(v string) *CreateUnionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUnionTaskResponseBody) SetSuccess(v bool) *CreateUnionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUnionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
