// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopAgentResponseBody
	GetRequestId() *string
}

type StopAgentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAgentResponseBody) GoString() string {
	return s.String()
}

func (s *StopAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAgentResponseBody) SetRequestId(v string) *StopAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
