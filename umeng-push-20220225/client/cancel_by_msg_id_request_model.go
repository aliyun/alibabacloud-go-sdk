// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelByMsgIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMsgId(v string) *CancelByMsgIdRequest
	GetMsgId() *string
}

type CancelByMsgIdRequest struct {
	// example:
	//
	// ucj0242167047014687101
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s CancelByMsgIdRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelByMsgIdRequest) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *CancelByMsgIdRequest) SetMsgId(v string) *CancelByMsgIdRequest {
	s.MsgId = &v
	return s
}

func (s *CancelByMsgIdRequest) Validate() error {
	return dara.Validate(s)
}
