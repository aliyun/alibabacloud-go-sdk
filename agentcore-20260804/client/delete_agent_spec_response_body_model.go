// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteAgentSpecResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteAgentSpecResponseBody
	GetRequestId() *string
}

type DeleteAgentSpecResponseBody struct {
	// The response data.
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAgentSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpecResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAgentSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentSpecResponseBody) SetData(v bool) *DeleteAgentSpecResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAgentSpecResponseBody) SetRequestId(v string) *DeleteAgentSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
