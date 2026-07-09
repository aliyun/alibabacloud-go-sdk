// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgentSpaceResponseBody
	GetRequestId() *string
}

type DeleteAgentSpaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D0173835-9E0F-508F-8BFA-9F556E59C302
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAgentSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentSpaceResponseBody) SetRequestId(v string) *DeleteAgentSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
