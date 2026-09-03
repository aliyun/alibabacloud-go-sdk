// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMsgStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMsgId(v string) *QueryMsgStatRequest
	GetMsgId() *string
}

type QueryMsgStatRequest struct {
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s QueryMsgStatRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMsgStatRequest) GoString() string {
	return s.String()
}

func (s *QueryMsgStatRequest) GetMsgId() *string {
	return s.MsgId
}

func (s *QueryMsgStatRequest) SetMsgId(v string) *QueryMsgStatRequest {
	s.MsgId = &v
	return s
}

func (s *QueryMsgStatRequest) Validate() error {
	return dara.Validate(s)
}
