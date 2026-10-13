// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVoiceCloneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyVoiceCloneResponseBody
	GetCode() *string
	SetData(v string) *ModifyVoiceCloneResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ModifyVoiceCloneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyVoiceCloneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyVoiceCloneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyVoiceCloneResponseBody
	GetSuccess() *bool
}

type ModifyVoiceCloneResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// cosyvoice-v2-067dc4b2e9d34dbeb5de20b5b093d1e3
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 7A573837-3AD3-54CF-930A-07A3287042C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyVoiceCloneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVoiceCloneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVoiceCloneResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyVoiceCloneResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyVoiceCloneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyVoiceCloneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyVoiceCloneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVoiceCloneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyVoiceCloneResponseBody) SetCode(v string) *ModifyVoiceCloneResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyVoiceCloneResponseBody) SetData(v string) *ModifyVoiceCloneResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyVoiceCloneResponseBody) SetHttpStatusCode(v int32) *ModifyVoiceCloneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyVoiceCloneResponseBody) SetMessage(v string) *ModifyVoiceCloneResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyVoiceCloneResponseBody) SetRequestId(v string) *ModifyVoiceCloneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVoiceCloneResponseBody) SetSuccess(v bool) *ModifyVoiceCloneResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyVoiceCloneResponseBody) Validate() error {
	return dara.Validate(s)
}
