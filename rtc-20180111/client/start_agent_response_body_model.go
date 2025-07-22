// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartAgentResponseBody
	GetRequestId() *string
}

type StartAgentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAgentResponseBody) GoString() string {
	return s.String()
}

func (s *StartAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAgentResponseBody) SetRequestId(v string) *StartAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
