// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTextMsgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendTextMsgResponseBody
	GetRequestId() *string
	SetStatus(v string) *SendTextMsgResponseBody
	GetStatus() *string
}

type SendTextMsgResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SendTextMsgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendTextMsgResponseBody) GoString() string {
	return s.String()
}

func (s *SendTextMsgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendTextMsgResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SendTextMsgResponseBody) SetRequestId(v string) *SendTextMsgResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendTextMsgResponseBody) SetStatus(v string) *SendTextMsgResponseBody {
	s.Status = &v
	return s
}

func (s *SendTextMsgResponseBody) Validate() error {
	return dara.Validate(s)
}
