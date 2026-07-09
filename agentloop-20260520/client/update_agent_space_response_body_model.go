// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAgentSpaceResponseBody
	GetRequestId() *string
}

type UpdateAgentSpaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AA689779-61AB-5077-BD91-9F7EA1205D68
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAgentSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentSpaceResponseBody) SetRequestId(v string) *UpdateAgentSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
