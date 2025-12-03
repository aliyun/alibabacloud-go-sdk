// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendByDeviceResponseBody
	GetCode() *string
	SetData(v *SendByDeviceResponseBodyData) *SendByDeviceResponseBody
	GetData() *SendByDeviceResponseBodyData
	SetHttpStatusCode(v int32) *SendByDeviceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendByDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendByDeviceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendByDeviceResponseBody
	GetSuccess() *bool
}

type SendByDeviceResponseBody struct {
	// example:
	//
	// 0
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 内部错误
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 74808AA4-A044-102F-8F5F-AFE4D97A0F26
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *SendByDeviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendByDeviceResponseBody) GetData() *SendByDeviceResponseBodyData {
	return s.Data
}

func (s *SendByDeviceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendByDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendByDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendByDeviceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendByDeviceResponseBody) SetCode(v string) *SendByDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *SendByDeviceResponseBody) SetData(v *SendByDeviceResponseBodyData) *SendByDeviceResponseBody {
	s.Data = v
	return s
}

func (s *SendByDeviceResponseBody) SetHttpStatusCode(v int32) *SendByDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByDeviceResponseBody) SetMessage(v string) *SendByDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *SendByDeviceResponseBody) SetRequestId(v string) *SendByDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByDeviceResponseBody) SetSuccess(v bool) *SendByDeviceResponseBody {
	s.Success = &v
	return s
}

func (s *SendByDeviceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendByDeviceResponseBodyData struct {
	// example:
	//
	// ula4wbu166876119986400
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByDeviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendByDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByDeviceResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *SendByDeviceResponseBodyData) SetMsgId(v string) *SendByDeviceResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *SendByDeviceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
