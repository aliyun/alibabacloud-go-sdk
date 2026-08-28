// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentIMChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgentIMChannelResponseBody
	GetRequestId() *string
}

type DeleteAgentIMChannelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAgentIMChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentIMChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentIMChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentIMChannelResponseBody) SetRequestId(v string) *DeleteAgentIMChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentIMChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
