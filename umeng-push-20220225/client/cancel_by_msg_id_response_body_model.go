// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelByMsgIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelByMsgIdResponseBody
	GetCode() *string
	SetData(v *CancelByMsgIdResponseBodyData) *CancelByMsgIdResponseBody
	GetData() *CancelByMsgIdResponseBodyData
	SetHttpStatusCode(v int32) *CancelByMsgIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CancelByMsgIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelByMsgIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelByMsgIdResponseBody
	GetSuccess() *bool
}

type CancelByMsgIdResponseBody struct {
	// example:
	//
	// 0
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CancelByMsgIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CancelByMsgIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelByMsgIdResponseBody) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelByMsgIdResponseBody) GetData() *CancelByMsgIdResponseBodyData {
	return s.Data
}

func (s *CancelByMsgIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CancelByMsgIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelByMsgIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelByMsgIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelByMsgIdResponseBody) SetCode(v string) *CancelByMsgIdResponseBody {
	s.Code = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetData(v *CancelByMsgIdResponseBodyData) *CancelByMsgIdResponseBody {
	s.Data = v
	return s
}

func (s *CancelByMsgIdResponseBody) SetHttpStatusCode(v int32) *CancelByMsgIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetMessage(v string) *CancelByMsgIdResponseBody {
	s.Message = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetRequestId(v string) *CancelByMsgIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelByMsgIdResponseBody) SetSuccess(v bool) *CancelByMsgIdResponseBody {
	s.Success = &v
	return s
}

func (s *CancelByMsgIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelByMsgIdResponseBodyData struct {
	// example:
	//
	// ucj0242167047014687101
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s CancelByMsgIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelByMsgIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdResponseBodyData) GetMsgId() *string {
	return s.MsgId
}

func (s *CancelByMsgIdResponseBodyData) SetMsgId(v string) *CancelByMsgIdResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *CancelByMsgIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
