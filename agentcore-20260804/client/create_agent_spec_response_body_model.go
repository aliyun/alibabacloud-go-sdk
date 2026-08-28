// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateAgentSpecResponseBody
	GetData() *string
	SetRequestId(v string) *CreateAgentSpecResponseBody
	GetRequestId() *string
}

type CreateAgentSpecResponseBody struct {
	// The response data.
	//
	// example:
	//
	// agentspec-1234567890abcdef
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAgentSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateAgentSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentSpecResponseBody) SetData(v string) *CreateAgentSpecResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAgentSpecResponseBody) SetRequestId(v string) *CreateAgentSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
