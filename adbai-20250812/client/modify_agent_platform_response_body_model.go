// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAgentPlatformResponseBody
	GetRequestId() *string
}

type ModifyAgentPlatformResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B47EED99-BFA5-529D-8D85-A6642421D390
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAgentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAgentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAgentPlatformResponseBody) SetRequestId(v string) *ModifyAgentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAgentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
