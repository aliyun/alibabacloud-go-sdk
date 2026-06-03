// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *DeleteBeebotIntentUserSayResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *DeleteBeebotIntentUserSayResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteBeebotIntentUserSayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteBeebotIntentUserSayResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBeebotIntentUserSayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBeebotIntentUserSayResponseBody
	GetSuccess() *bool
	SetUserSayId(v int64) *DeleteBeebotIntentUserSayResponseBody
	GetUserSayId() *int64
}

type DeleteBeebotIntentUserSayResponseBody struct {
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

func (s DeleteBeebotIntentUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentUserSayResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *DeleteBeebotIntentUserSayResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBeebotIntentUserSayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBeebotIntentUserSayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBeebotIntentUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBeebotIntentUserSayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBeebotIntentUserSayResponseBody) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *DeleteBeebotIntentUserSayResponseBody) SetBeebotRequestId(v string) *DeleteBeebotIntentUserSayResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponseBody) SetCode(v string) *DeleteBeebotIntentUserSayResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponseBody) SetHttpStatusCode(v int32) *DeleteBeebotIntentUserSayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponseBody) SetMessage(v string) *DeleteBeebotIntentUserSayResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponseBody) SetRequestId(v string) *DeleteBeebotIntentUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponseBody) SetSuccess(v bool) *DeleteBeebotIntentUserSayResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponseBody) SetUserSayId(v int64) *DeleteBeebotIntentUserSayResponseBody {
	s.UserSayId = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponseBody) Validate() error {
	return dara.Validate(s)
}
