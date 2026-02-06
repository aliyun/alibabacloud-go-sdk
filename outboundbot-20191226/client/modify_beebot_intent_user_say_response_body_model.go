// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentUserSayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *ModifyBeebotIntentUserSayResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *ModifyBeebotIntentUserSayResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyBeebotIntentUserSayResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyBeebotIntentUserSayResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBeebotIntentUserSayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBeebotIntentUserSayResponseBody
	GetSuccess() *bool
	SetUserSayId(v int64) *ModifyBeebotIntentUserSayResponseBody
	GetUserSayId() *int64
}

type ModifyBeebotIntentUserSayResponseBody struct {
	// example:
	//
	// A1F21BF2-CB21-1968-8039-C74699E7DDEB
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
	// Success
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

func (s ModifyBeebotIntentUserSayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentUserSayResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentUserSayResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *ModifyBeebotIntentUserSayResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBeebotIntentUserSayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBeebotIntentUserSayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBeebotIntentUserSayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBeebotIntentUserSayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBeebotIntentUserSayResponseBody) GetUserSayId() *int64 {
	return s.UserSayId
}

func (s *ModifyBeebotIntentUserSayResponseBody) SetBeebotRequestId(v string) *ModifyBeebotIntentUserSayResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponseBody) SetCode(v string) *ModifyBeebotIntentUserSayResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponseBody) SetHttpStatusCode(v int32) *ModifyBeebotIntentUserSayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponseBody) SetMessage(v string) *ModifyBeebotIntentUserSayResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponseBody) SetRequestId(v string) *ModifyBeebotIntentUserSayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponseBody) SetSuccess(v bool) *ModifyBeebotIntentUserSayResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponseBody) SetUserSayId(v int64) *ModifyBeebotIntentUserSayResponseBody {
	s.UserSayId = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponseBody) Validate() error {
	return dara.Validate(s)
}
