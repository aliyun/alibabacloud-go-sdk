// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeywordLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateKeywordLibResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateKeywordLibResponseBody
	GetData() *bool
	SetMsg(v string) *UpdateKeywordLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *UpdateKeywordLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateKeywordLibResponseBody
	GetSuccess() *bool
}

type UpdateKeywordLibResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateKeywordLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateKeywordLibResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateKeywordLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *UpdateKeywordLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKeywordLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateKeywordLibResponseBody) SetCode(v int32) *UpdateKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetData(v bool) *UpdateKeywordLibResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetMsg(v string) *UpdateKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetRequestId(v string) *UpdateKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) SetSuccess(v bool) *UpdateKeywordLibResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateKeywordLibResponseBody) Validate() error {
	return dara.Validate(s)
}
