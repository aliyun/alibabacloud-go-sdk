// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSdkStreamMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonStreamMessage(v string) *SendSdkStreamMessageResponseBody
	GetCommonStreamMessage() *string
}

type SendSdkStreamMessageResponseBody struct {
	// example:
	//
	// {"id":"123"}
	CommonStreamMessage *string `json:"commonStreamMessage,omitempty" xml:"commonStreamMessage,omitempty"`
}

func (s SendSdkStreamMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendSdkStreamMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendSdkStreamMessageResponseBody) GetCommonStreamMessage() *string {
	return s.CommonStreamMessage
}

func (s *SendSdkStreamMessageResponseBody) SetCommonStreamMessage(v string) *SendSdkStreamMessageResponseBody {
	s.CommonStreamMessage = &v
	return s
}

func (s *SendSdkStreamMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
