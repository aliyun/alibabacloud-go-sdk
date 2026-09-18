// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgentResponseBody
	GetSuccess() *bool
}

type DeleteAgentResponseBody struct {
	// The request ID, which is used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the Agent is deleted. The value `true` is returned for a successful response. An error response is returned for a failed request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetSuccess(v bool) *DeleteAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
