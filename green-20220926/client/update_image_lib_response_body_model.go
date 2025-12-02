// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateImageLibResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateImageLibResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateImageLibResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *UpdateImageLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *UpdateImageLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageLibResponseBody
	GetSuccess() *bool
}

type UpdateImageLibResponseBody struct {
	// Error code, consistent with HTTP status.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
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
	// Further description of the error code.
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

func (s UpdateImageLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateImageLibResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateImageLibResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateImageLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *UpdateImageLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageLibResponseBody) SetCode(v int32) *UpdateImageLibResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetData(v bool) *UpdateImageLibResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetHttpStatusCode(v int32) *UpdateImageLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetMsg(v string) *UpdateImageLibResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetRequestId(v string) *UpdateImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageLibResponseBody) SetSuccess(v bool) *UpdateImageLibResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageLibResponseBody) Validate() error {
	return dara.Validate(s)
}
