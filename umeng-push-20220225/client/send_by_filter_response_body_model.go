// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendByFilterResponseBody
	GetCode() *string
	SetData(v *SendByFilterResponseBodyData) *SendByFilterResponseBody
	GetData() *SendByFilterResponseBodyData
	SetHttpStatusCode(v int32) *SendByFilterResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendByFilterResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendByFilterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendByFilterResponseBody
	GetSuccess() *bool
}

type SendByFilterResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByFilterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 86C4236B-D6C2-1E31-8370-2FAEC5CFE012
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendByFilterResponseBody) GoString() string {
	return s.String()
}

func (s *SendByFilterResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendByFilterResponseBody) GetData() *SendByFilterResponseBodyData {
	return s.Data
}

func (s *SendByFilterResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendByFilterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendByFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendByFilterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendByFilterResponseBody) SetCode(v string) *SendByFilterResponseBody {
	s.Code = &v
	return s
}

func (s *SendByFilterResponseBody) SetData(v *SendByFilterResponseBodyData) *SendByFilterResponseBody {
	s.Data = v
	return s
}

func (s *SendByFilterResponseBody) SetHttpStatusCode(v int32) *SendByFilterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByFilterResponseBody) SetMessage(v string) *SendByFilterResponseBody {
	s.Message = &v
	return s
}

func (s *SendByFilterResponseBody) SetRequestId(v string) *SendByFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByFilterResponseBody) SetSuccess(v bool) *SendByFilterResponseBody {
	s.Success = &v
	return s
}

func (s *SendByFilterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendByFilterResponseBodyData struct {
	// example:
	//
	// usouag1167056659161101
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByFilterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendByFilterResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByFilterResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *SendByFilterResponseBodyData) SetMsgId(v string) *SendByFilterResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *SendByFilterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
