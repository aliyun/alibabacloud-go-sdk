// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *ModifyBeebotIntentLgfResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *ModifyBeebotIntentLgfResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyBeebotIntentLgfResponseBody
	GetHttpStatusCode() *int32
	SetLgfId(v int64) *ModifyBeebotIntentLgfResponseBody
	GetLgfId() *int64
	SetMessage(v string) *ModifyBeebotIntentLgfResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBeebotIntentLgfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBeebotIntentLgfResponseBody
	GetSuccess() *bool
}

type ModifyBeebotIntentLgfResponseBody struct {
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
	// 5666117
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 302C67BD-19FF-5B66-A45D-F95544604155
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBeebotIntentLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentLgfResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentLgfResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *ModifyBeebotIntentLgfResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBeebotIntentLgfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBeebotIntentLgfResponseBody) GetLgfId() *int64 {
	return s.LgfId
}

func (s *ModifyBeebotIntentLgfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBeebotIntentLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBeebotIntentLgfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBeebotIntentLgfResponseBody) SetBeebotRequestId(v string) *ModifyBeebotIntentLgfResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponseBody) SetCode(v string) *ModifyBeebotIntentLgfResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponseBody) SetHttpStatusCode(v int32) *ModifyBeebotIntentLgfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponseBody) SetLgfId(v int64) *ModifyBeebotIntentLgfResponseBody {
	s.LgfId = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponseBody) SetMessage(v string) *ModifyBeebotIntentLgfResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponseBody) SetRequestId(v string) *ModifyBeebotIntentLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponseBody) SetSuccess(v bool) *ModifyBeebotIntentLgfResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponseBody) Validate() error {
	return dara.Validate(s)
}
