// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsertTaskDetailResponseBody
	GetCode() *string
	SetData(v string) *InsertTaskDetailResponseBody
	GetData() *string
	SetMessage(v string) *InsertTaskDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsertTaskDetailResponseBody
	GetSuccess() *bool
}

type InsertTaskDetailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *InsertTaskDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsertTaskDetailResponseBody) GetData() *string {
	return s.Data
}

func (s *InsertTaskDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsertTaskDetailResponseBody) SetCode(v string) *InsertTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetData(v string) *InsertTaskDetailResponseBody {
	s.Data = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetMessage(v string) *InsertTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetRequestId(v string) *InsertTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertTaskDetailResponseBody) SetSuccess(v bool) *InsertTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *InsertTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
