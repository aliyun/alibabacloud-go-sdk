// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *StartServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartServiceResponseBody
	GetRequestId() *string
}

type StartServiceResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Service [foo] in region [cn-shanghai] is starting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartServiceResponseBody) SetMessage(v string) *StartServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartServiceResponseBody) SetRequestId(v string) *StartServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
