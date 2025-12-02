// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeywordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteKeywordResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteKeywordResponseBody
	GetData() *bool
	SetMsg(v string) *DeleteKeywordResponseBody
	GetMsg() *string
	SetRequestId(v string) *DeleteKeywordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteKeywordResponseBody
	GetSuccess() *bool
}

type DeleteKeywordResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Response message for this request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteKeywordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeywordResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteKeywordResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteKeywordResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DeleteKeywordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKeywordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteKeywordResponseBody) SetCode(v int32) *DeleteKeywordResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetData(v bool) *DeleteKeywordResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetMsg(v string) *DeleteKeywordResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetRequestId(v string) *DeleteKeywordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKeywordResponseBody) SetSuccess(v bool) *DeleteKeywordResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteKeywordResponseBody) Validate() error {
	return dara.Validate(s)
}
