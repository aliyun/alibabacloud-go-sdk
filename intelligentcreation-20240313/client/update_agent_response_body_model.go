// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAgentResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateAgentResponseBody
	GetStatus() *string
}

type UpdateAgentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetStatus(v string) *UpdateAgentResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
