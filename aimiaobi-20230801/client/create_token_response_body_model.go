// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTokenResponseBody
	GetCode() *string
	SetData(v *CreateTokenResponseBodyData) *CreateTokenResponseBody
	GetData() *CreateTokenResponseBodyData
	SetHttpStatusCode(v int32) *CreateTokenResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTokenResponseBody
	GetSuccess() *bool
}

type CreateTokenResponseBody struct {
	// example:
	//
	// NoData
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTokenResponseBody) GetData() *CreateTokenResponseBodyData {
	return s.Data
}

func (s *CreateTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTokenResponseBody) SetCode(v string) *CreateTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTokenResponseBody) SetData(v *CreateTokenResponseBodyData) *CreateTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateTokenResponseBody) SetHttpStatusCode(v int32) *CreateTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTokenResponseBody) SetMessage(v string) *CreateTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTokenResponseBody) SetRequestId(v string) *CreateTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResponseBody) SetSuccess(v bool) *CreateTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateTokenResponseBodyData struct {
	// example:
	//
	// 1705388704855
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBodyData) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *CreateTokenResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *CreateTokenResponseBodyData) SetExpiredTime(v int64) *CreateTokenResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *CreateTokenResponseBodyData) SetToken(v string) *CreateTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *CreateTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
