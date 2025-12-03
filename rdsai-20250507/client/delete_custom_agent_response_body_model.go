// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomAgentResponseBody
	GetRequestId() *string
	SetResult(v string) *DeleteCustomAgentResponseBody
	GetResult() *string
}

type DeleteCustomAgentResponseBody struct {
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomAgentResponseBody) GetResult() *string {
	return s.Result
}

func (s *DeleteCustomAgentResponseBody) SetRequestId(v string) *DeleteCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) SetResult(v string) *DeleteCustomAgentResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteCustomAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
