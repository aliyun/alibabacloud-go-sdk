// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyShareImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyShareImageResponseBody
	GetCode() *string
	SetData(v bool) *ModifyShareImageResponseBody
	GetData() *bool
	SetMessage(v string) *ModifyShareImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyShareImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyShareImageResponseBody
	GetSuccess() *bool
}

type ModifyShareImageResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyShareImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyShareImageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyShareImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyShareImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyShareImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyShareImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyShareImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyShareImageResponseBody) SetCode(v string) *ModifyShareImageResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyShareImageResponseBody) SetData(v bool) *ModifyShareImageResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyShareImageResponseBody) SetMessage(v string) *ModifyShareImageResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyShareImageResponseBody) SetRequestId(v string) *ModifyShareImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyShareImageResponseBody) SetSuccess(v bool) *ModifyShareImageResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyShareImageResponseBody) Validate() error {
	return dara.Validate(s)
}
