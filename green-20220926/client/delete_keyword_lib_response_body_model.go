// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeywordLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteKeywordLibResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteKeywordLibResponseBody
	GetData() *bool
	SetMsg(v string) *DeleteKeywordLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *DeleteKeywordLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKeywordLibResponseBody
	GetSuccess() *bool
}

type DeleteKeywordLibResponseBody struct {
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

func (s DeleteKeywordLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteKeywordLibResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteKeywordLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DeleteKeywordLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKeywordLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKeywordLibResponseBody) SetCode(v int32) *DeleteKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetData(v bool) *DeleteKeywordLibResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetMsg(v string) *DeleteKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetRequestId(v string) *DeleteKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) SetSuccess(v bool) *DeleteKeywordLibResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKeywordLibResponseBody) Validate() error {
	return dara.Validate(s)
}
