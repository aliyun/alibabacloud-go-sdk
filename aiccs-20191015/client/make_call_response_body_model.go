// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MakeCallResponseBody
	GetCode() *string
	SetData(v string) *MakeCallResponseBody
	GetData() *string
	SetMessage(v string) *MakeCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *MakeCallResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MakeCallResponseBody
	GetSuccess() *bool
}

type MakeCallResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MakeCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MakeCallResponseBody) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *MakeCallResponseBody) GetData() *string {
	return s.Data
}

func (s *MakeCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MakeCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MakeCallResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MakeCallResponseBody) SetCode(v string) *MakeCallResponseBody {
	s.Code = &v
	return s
}

func (s *MakeCallResponseBody) SetData(v string) *MakeCallResponseBody {
	s.Data = &v
	return s
}

func (s *MakeCallResponseBody) SetMessage(v string) *MakeCallResponseBody {
	s.Message = &v
	return s
}

func (s *MakeCallResponseBody) SetRequestId(v string) *MakeCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeCallResponseBody) SetSuccess(v bool) *MakeCallResponseBody {
	s.Success = &v
	return s
}

func (s *MakeCallResponseBody) Validate() error {
	return dara.Validate(s)
}
