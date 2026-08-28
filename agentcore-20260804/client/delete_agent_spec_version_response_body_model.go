// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpecVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteAgentSpecVersionResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteAgentSpecVersionResponseBody
	GetRequestId() *string
}

type DeleteAgentSpecVersionResponseBody struct {
	// The response data.
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAgentSpecVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpecVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpecVersionResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAgentSpecVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentSpecVersionResponseBody) SetData(v bool) *DeleteAgentSpecVersionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAgentSpecVersionResponseBody) SetRequestId(v string) *DeleteAgentSpecVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentSpecVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
