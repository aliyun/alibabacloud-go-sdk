// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeadLockDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeadLockDetailResponseBody
	GetCode() *string
	SetData(v string) *GetDeadLockDetailResponseBody
	GetData() *string
	SetMessage(v string) *GetDeadLockDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeadLockDetailResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDeadLockDetailResponseBody
	GetSuccess() *string
	SetSynchro(v string) *GetDeadLockDetailResponseBody
	GetSynchro() *string
}

type GetDeadLockDetailResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9CB97BC4-6479-55D0-B9D0-EA925AFE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// None
	Synchro *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetDeadLockDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeadLockDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeadLockDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeadLockDetailResponseBody) GetData() *string {
	return s.Data
}

func (s *GetDeadLockDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeadLockDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeadLockDetailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDeadLockDetailResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *GetDeadLockDetailResponseBody) SetCode(v string) *GetDeadLockDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeadLockDetailResponseBody) SetData(v string) *GetDeadLockDetailResponseBody {
	s.Data = &v
	return s
}

func (s *GetDeadLockDetailResponseBody) SetMessage(v string) *GetDeadLockDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeadLockDetailResponseBody) SetRequestId(v string) *GetDeadLockDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeadLockDetailResponseBody) SetSuccess(v string) *GetDeadLockDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeadLockDetailResponseBody) SetSynchro(v string) *GetDeadLockDetailResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetDeadLockDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
