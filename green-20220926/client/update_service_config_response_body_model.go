// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateServiceConfigResponseBody
	GetCode() *int32
	SetData(v bool) *UpdateServiceConfigResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateServiceConfigResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *UpdateServiceConfigResponseBody
	GetMsg() *string
	SetRequestId(v string) *UpdateServiceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateServiceConfigResponseBody
	GetSuccess() *bool
}

type UpdateServiceConfigResponseBody struct {
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
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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

func (s UpdateServiceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateServiceConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateServiceConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateServiceConfigResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *UpdateServiceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateServiceConfigResponseBody) SetCode(v int32) *UpdateServiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetData(v bool) *UpdateServiceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetHttpStatusCode(v int32) *UpdateServiceConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetMsg(v string) *UpdateServiceConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetRequestId(v string) *UpdateServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) SetSuccess(v bool) *UpdateServiceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateServiceConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
