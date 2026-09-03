// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendByAliasFileIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendByAliasFileIdResponseBody
	GetCode() *string
	SetData(v *SendByAliasFileIdResponseBodyData) *SendByAliasFileIdResponseBody
	GetData() *SendByAliasFileIdResponseBodyData
	SetHttpStatusCode(v int32) *SendByAliasFileIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SendByAliasFileIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendByAliasFileIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendByAliasFileIdResponseBody
	GetSuccess() *bool
}

type SendByAliasFileIdResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *SendByAliasFileIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendByAliasFileIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasFileIdResponseBody) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendByAliasFileIdResponseBody) GetData() *SendByAliasFileIdResponseBodyData {
	return s.Data
}

func (s *SendByAliasFileIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SendByAliasFileIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendByAliasFileIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendByAliasFileIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendByAliasFileIdResponseBody) SetCode(v string) *SendByAliasFileIdResponseBody {
	s.Code = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetData(v *SendByAliasFileIdResponseBodyData) *SendByAliasFileIdResponseBody {
	s.Data = v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetHttpStatusCode(v int32) *SendByAliasFileIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetMessage(v string) *SendByAliasFileIdResponseBody {
	s.Message = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetRequestId(v string) *SendByAliasFileIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) SetSuccess(v bool) *SendByAliasFileIdResponseBody {
	s.Success = &v
	return s
}

func (s *SendByAliasFileIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendByAliasFileIdResponseBodyData struct {
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s SendByAliasFileIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendByAliasFileIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendByAliasFileIdResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *SendByAliasFileIdResponseBodyData) SetMsgId(v string) *SendByAliasFileIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *SendByAliasFileIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
