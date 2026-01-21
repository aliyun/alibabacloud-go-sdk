// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAgentPlatformResponseBody
	GetRequestId() *string
}

type CreateAgentPlatformResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentPlatformResponseBody) SetRequestId(v string) *CreateAgentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
