// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTurnCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTurnCredentialsResponseBody
	GetCode() *string
	SetData(v *GetTurnCredentialsResponseBodyData) *GetTurnCredentialsResponseBody
	GetData() *GetTurnCredentialsResponseBodyData
	SetHttpStatusCode(v int32) *GetTurnCredentialsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTurnCredentialsResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetTurnCredentialsResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetTurnCredentialsResponseBody
	GetRequestId() *string
}

type GetTurnCredentialsResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTurnCredentialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTurnCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTurnCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTurnCredentialsResponseBody) GetData() *GetTurnCredentialsResponseBodyData {
	return s.Data
}

func (s *GetTurnCredentialsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTurnCredentialsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTurnCredentialsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetTurnCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTurnCredentialsResponseBody) SetCode(v string) *GetTurnCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetData(v *GetTurnCredentialsResponseBodyData) *GetTurnCredentialsResponseBody {
	s.Data = v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetHttpStatusCode(v int32) *GetTurnCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetMessage(v string) *GetTurnCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetParams(v []*string) *GetTurnCredentialsResponseBody {
	s.Params = v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetRequestId(v string) *GetTurnCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTurnCredentialsResponseBodyData struct {
	// example:
	//
	// M0NQNG/uRUrfIxW7er/S9gKX****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 1602585817:****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetTurnCredentialsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTurnCredentialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *GetTurnCredentialsResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetTurnCredentialsResponseBodyData) SetPassword(v string) *GetTurnCredentialsResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetTurnCredentialsResponseBodyData) SetUserName(v string) *GetTurnCredentialsResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetTurnCredentialsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
