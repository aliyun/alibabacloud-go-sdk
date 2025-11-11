// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneGlobalReplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInterveneGlobalReplyResponseBody
	GetCode() *string
	SetData(v *GetInterveneGlobalReplyResponseBodyData) *GetInterveneGlobalReplyResponseBody
	GetData() *GetInterveneGlobalReplyResponseBodyData
	SetHttpStatusCode(v int32) *GetInterveneGlobalReplyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInterveneGlobalReplyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInterveneGlobalReplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInterveneGlobalReplyResponseBody
	GetSuccess() *bool
}

type GetInterveneGlobalReplyResponseBody struct {
	// example:
	//
	// 0
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInterveneGlobalReplyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneGlobalReplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInterveneGlobalReplyResponseBody) GetData() *GetInterveneGlobalReplyResponseBodyData {
	return s.Data
}

func (s *GetInterveneGlobalReplyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInterveneGlobalReplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInterveneGlobalReplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInterveneGlobalReplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInterveneGlobalReplyResponseBody) SetCode(v string) *GetInterveneGlobalReplyResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetData(v *GetInterveneGlobalReplyResponseBodyData) *GetInterveneGlobalReplyResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetHttpStatusCode(v int32) *GetInterveneGlobalReplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetMessage(v string) *GetInterveneGlobalReplyResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetRequestId(v string) *GetInterveneGlobalReplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetSuccess(v bool) *GetInterveneGlobalReplyResponseBody {
	s.Success = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterveneGlobalReplyResponseBodyData struct {
	Code            *int32                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	ReplyMessagList []*GetInterveneGlobalReplyResponseBodyDataReplyMessagList `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty" type:"Repeated"`
}

func (s GetInterveneGlobalReplyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *GetInterveneGlobalReplyResponseBodyData) GetReplyMessagList() []*GetInterveneGlobalReplyResponseBodyDataReplyMessagList {
	return s.ReplyMessagList
}

func (s *GetInterveneGlobalReplyResponseBodyData) SetCode(v int32) *GetInterveneGlobalReplyResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBodyData) SetReplyMessagList(v []*GetInterveneGlobalReplyResponseBodyDataReplyMessagList) *GetInterveneGlobalReplyResponseBodyData {
	s.ReplyMessagList = v
	return s
}

func (s *GetInterveneGlobalReplyResponseBodyData) Validate() error {
	if s.ReplyMessagList != nil {
		for _, item := range s.ReplyMessagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInterveneGlobalReplyResponseBodyDataReplyMessagList struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// namespace_qa_query
	ReplyType *string `json:"ReplyType,omitempty" xml:"ReplyType,omitempty"`
}

func (s GetInterveneGlobalReplyResponseBodyDataReplyMessagList) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBodyDataReplyMessagList) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) GetMessage() *string {
	return s.Message
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) GetReplyType() *string {
	return s.ReplyType
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) SetMessage(v string) *GetInterveneGlobalReplyResponseBodyDataReplyMessagList {
	s.Message = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) SetReplyType(v string) *GetInterveneGlobalReplyResponseBodyDataReplyMessagList {
	s.ReplyType = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) Validate() error {
	return dara.Validate(s)
}
