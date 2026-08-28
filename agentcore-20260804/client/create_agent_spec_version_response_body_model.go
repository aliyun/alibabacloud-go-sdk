// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateAgentSpecVersionResponseBody
	GetData() *string
	SetRequestId(v string) *CreateAgentSpecVersionResponseBody
	GetRequestId() *string
}

type CreateAgentSpecVersionResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// 1.0.0
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAgentSpecVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateAgentSpecVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentSpecVersionResponseBody) SetData(v string) *CreateAgentSpecVersionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAgentSpecVersionResponseBody) SetRequestId(v string) *CreateAgentSpecVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentSpecVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
