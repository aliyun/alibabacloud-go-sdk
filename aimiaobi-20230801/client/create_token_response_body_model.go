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
	// The status code.
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *CreateTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. \\`true\\` indicates success and \\`false\\` indicates failure.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTokenResponseBodyData struct {
	// The expiration time.
	//
	// example:
	//
	// 1705388704855
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The access token.
	//
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
