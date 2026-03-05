// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetUserResponseBody
	GetAppId() *string
	SetCode(v string) *GetUserResponseBody
	GetCode() *string
	SetHost(v string) *GetUserResponseBody
	GetHost() *string
	SetMessage(v string) *GetUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetToken(v string) *GetUserResponseBody
	GetToken() *string
}

type GetUserResponseBody struct {
	// example:
	//
	// 20200425******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ai-service.******.alicontainer.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 897658D5-1FB8-5CFA-A030-727CCAE59EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Njk0Njk******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserResponseBody) GetHost() *string {
	return s.Host
}

func (s *GetUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetUserResponseBody) SetAppId(v string) *GetUserResponseBody {
	s.AppId = &v
	return s
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetHost(v string) *GetUserResponseBody {
	s.Host = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetToken(v string) *GetUserResponseBody {
	s.Token = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	return dara.Validate(s)
}
