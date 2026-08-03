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
	SetStatus(v string) *DeleteAgentResponseBody
	GetStatus() *string
}

type DeleteAgentResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *DeleteAgentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetStatus(v string) *DeleteAgentResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
