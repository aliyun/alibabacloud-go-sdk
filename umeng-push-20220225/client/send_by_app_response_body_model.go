// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendByAppResponseBody
	GetCode() *string
	SetData(v *SendByAppResponseBodyData) *SendByAppResponseBody
	GetData() *SendByAppResponseBodyData
	SetHttpStatusCode(v int32) *SendByAppResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendByAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendByAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendByAppResponseBody
	GetSuccess() *bool
}

type SendByAppResponseBody struct {
	// example:
	//
	// 0
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
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

func (s SendByAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendByAppResponseBody) GoString() string {
	return s.String()
}

func (s *SendByAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendByAppResponseBody) GetData() *SendByAppResponseBodyData {
	return s.Data
}

func (s *SendByAppResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendByAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendByAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendByAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendByAppResponseBody) SetCode(v string) *SendByAppResponseBody {
	s.Code = &v
	return s
}

func (s *SendByAppResponseBody) SetData(v *SendByAppResponseBodyData) *SendByAppResponseBody {
	s.Data = v
	return s
}

func (s *SendByAppResponseBody) SetHttpStatusCode(v int32) *SendByAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByAppResponseBody) SetMessage(v string) *SendByAppResponseBody {
	s.Message = &v
	return s
}

func (s *SendByAppResponseBody) SetRequestId(v string) *SendByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByAppResponseBody) SetSuccess(v bool) *SendByAppResponseBody {
	s.Success = &v
	return s
}

func (s *SendByAppResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendByAppResponseBodyData struct {
	// example:
	//
	// um3zlgb166876370784300
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByAppResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendByAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByAppResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *SendByAppResponseBodyData) SetMsgId(v string) *SendByAppResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *SendByAppResponseBodyData) Validate() error {
	return dara.Validate(s)
}
