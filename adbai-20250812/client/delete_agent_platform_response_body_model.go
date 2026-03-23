// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgentPlatformResponseBody
	GetRequestId() *string
}

type DeleteAgentPlatformResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentPlatformResponseBody) SetRequestId(v string) *DeleteAgentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
