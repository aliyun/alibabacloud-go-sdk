// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddImageLibResponseBody
	GetCode() *int32
	SetData(v bool) *AddImageLibResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AddImageLibResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *AddImageLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *AddImageLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddImageLibResponseBody
	GetSuccess() *bool
}

type AddImageLibResponseBody struct {
	// Return code. A return of 200 represents success.
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
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddImageLibResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddImageLibResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddImageLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *AddImageLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImageLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddImageLibResponseBody) SetCode(v int32) *AddImageLibResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageLibResponseBody) SetData(v bool) *AddImageLibResponseBody {
	s.Data = &v
	return s
}

func (s *AddImageLibResponseBody) SetHttpStatusCode(v int32) *AddImageLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddImageLibResponseBody) SetMsg(v string) *AddImageLibResponseBody {
	s.Msg = &v
	return s
}

func (s *AddImageLibResponseBody) SetRequestId(v string) *AddImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageLibResponseBody) SetSuccess(v bool) *AddImageLibResponseBody {
	s.Success = &v
	return s
}

func (s *AddImageLibResponseBody) Validate() error {
	return dara.Validate(s)
}
