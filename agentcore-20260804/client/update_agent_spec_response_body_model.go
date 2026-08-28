// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateAgentSpecResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateAgentSpecResponseBody
	GetRequestId() *string
}

type UpdateAgentSpecResponseBody struct {
	// The response data.
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAgentSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentSpecResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateAgentSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentSpecResponseBody) SetData(v bool) *UpdateAgentSpecResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAgentSpecResponseBody) SetRequestId(v string) *UpdateAgentSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
