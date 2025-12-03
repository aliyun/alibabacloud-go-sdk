// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendByAliasResponseBody
	GetCode() *string
	SetData(v *SendByAliasResponseBodyData) *SendByAliasResponseBody
	GetData() *SendByAliasResponseBodyData
	SetHttpStatusCode(v int32) *SendByAliasResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendByAliasResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendByAliasResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendByAliasResponseBody
	GetSuccess() *bool
}

type SendByAliasResponseBody struct {
	// example:
	//
	// 0
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SendByAliasResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SendByAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasResponseBody) GoString() string {
	return s.String()
}

func (s *SendByAliasResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendByAliasResponseBody) GetData() *SendByAliasResponseBodyData {
	return s.Data
}

func (s *SendByAliasResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendByAliasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendByAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendByAliasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendByAliasResponseBody) SetCode(v string) *SendByAliasResponseBody {
	s.Code = &v
	return s
}

func (s *SendByAliasResponseBody) SetData(v *SendByAliasResponseBodyData) *SendByAliasResponseBody {
	s.Data = v
	return s
}

func (s *SendByAliasResponseBody) SetHttpStatusCode(v int32) *SendByAliasResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByAliasResponseBody) SetMessage(v string) *SendByAliasResponseBody {
	s.Message = &v
	return s
}

func (s *SendByAliasResponseBody) SetRequestId(v string) *SendByAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByAliasResponseBody) SetSuccess(v bool) *SendByAliasResponseBody {
	s.Success = &v
	return s
}

func (s *SendByAliasResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendByAliasResponseBodyData struct {
	// example:
	//
	// uacxo27167041814609201
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByAliasResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByAliasResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *SendByAliasResponseBodyData) SetMsgId(v string) *SendByAliasResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *SendByAliasResponseBodyData) Validate() error {
	return dara.Validate(s)
}
