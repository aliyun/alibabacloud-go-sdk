// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAgentSpaceResponseBody
	GetRequestId() *string
}

type CreateAgentSpaceResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAgentSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentSpaceResponseBody) SetRequestId(v string) *CreateAgentSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
