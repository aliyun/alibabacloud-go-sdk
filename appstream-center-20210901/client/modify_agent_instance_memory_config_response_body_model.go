// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentInstanceMemoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAgentInstanceMemoryConfigResponseBody
	GetRequestId() *string
}

type ModifyAgentInstanceMemoryConfigResponseBody struct {
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 01A1CEEB-CDFF-18F9-B2BF-9997600D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAgentInstanceMemoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentInstanceMemoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAgentInstanceMemoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAgentInstanceMemoryConfigResponseBody) SetRequestId(v string) *ModifyAgentInstanceMemoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAgentInstanceMemoryConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
