// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessageResendByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsDLQMessageResendByIdResponseBodyData) *OnsDLQMessageResendByIdResponseBody
	GetData() *OnsDLQMessageResendByIdResponseBodyData
	SetRequestId(v string) *OnsDLQMessageResendByIdResponseBody
	GetRequestId() *string
}

type OnsDLQMessageResendByIdResponseBody struct {
	Data *OnsDLQMessageResendByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// D94CC769-4DC3-4690-A868-9D0631B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsDLQMessageResendByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageResendByIdResponseBody) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdResponseBody) GetData() *OnsDLQMessageResendByIdResponseBodyData {
	return s.Data
}

func (s *OnsDLQMessageResendByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsDLQMessageResendByIdResponseBody) SetData(v *OnsDLQMessageResendByIdResponseBodyData) *OnsDLQMessageResendByIdResponseBody {
	s.Data = v
	return s
}

func (s *OnsDLQMessageResendByIdResponseBody) SetRequestId(v string) *OnsDLQMessageResendByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsDLQMessageResendByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsDLQMessageResendByIdResponseBodyData struct {
	MsgId []*string `json:"MsgId,omitempty" xml:"MsgId,omitempty" type:"Repeated"`
}

func (s OnsDLQMessageResendByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessageResendByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsDLQMessageResendByIdResponseBodyData) GetMsgId() []*string {
	return s.MsgId
}

func (s *OnsDLQMessageResendByIdResponseBodyData) SetMsgId(v []*string) *OnsDLQMessageResendByIdResponseBodyData {
	s.MsgId = v
	return s
}

func (s *OnsDLQMessageResendByIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
