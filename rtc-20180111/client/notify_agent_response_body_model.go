// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *NotifyAgentResponseBody
	GetRequestId() *string
}

type NotifyAgentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NotifyAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s NotifyAgentResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *NotifyAgentResponseBody) SetRequestId(v string) *NotifyAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *NotifyAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
