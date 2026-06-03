// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *CreateBeebotIntentUserSayResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *CreateBeebotIntentUserSayResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateBeebotIntentUserSayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBeebotIntentUserSayResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBeebotIntentUserSayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBeebotIntentUserSayResponseBody
	GetSuccess() *bool
	SetUserSayId(v int64) *CreateBeebotIntentUserSayResponseBody
	GetUserSayId() *int64
}

type CreateBeebotIntentUserSayResponseBody struct {
	// example:
	//
	// 0B219FCB-EC71-1F08-BB1B-0E87C20158C8
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 17448458
	UserSayId *int64 `json:"UserSayId,omitempty" xml:"UserSayId,omitempty"`
}

func (s CreateBeebotIntentUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentUserSayResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *CreateBeebotIntentUserSayResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBeebotIntentUserSayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBeebotIntentUserSayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBeebotIntentUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBeebotIntentUserSayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBeebotIntentUserSayResponseBody) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *CreateBeebotIntentUserSayResponseBody) SetBeebotRequestId(v string) *CreateBeebotIntentUserSayResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponseBody) SetCode(v string) *CreateBeebotIntentUserSayResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponseBody) SetHttpStatusCode(v int32) *CreateBeebotIntentUserSayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponseBody) SetMessage(v string) *CreateBeebotIntentUserSayResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponseBody) SetRequestId(v string) *CreateBeebotIntentUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponseBody) SetSuccess(v bool) *CreateBeebotIntentUserSayResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponseBody) SetUserSayId(v int64) *CreateBeebotIntentUserSayResponseBody {
	s.UserSayId = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponseBody) Validate() error {
	return dara.Validate(s)
}
