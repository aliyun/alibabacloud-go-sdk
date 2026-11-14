// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAgentInstanceMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearAgentInstanceMemoryResponseBody
	GetRequestId() *string
}

type ClearAgentInstanceMemoryResponseBody struct {
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// 01A1CEEB-CDFF-18F9-B2BF-9997600D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearAgentInstanceMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearAgentInstanceMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ClearAgentInstanceMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearAgentInstanceMemoryResponseBody) SetRequestId(v string) *ClearAgentInstanceMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearAgentInstanceMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
