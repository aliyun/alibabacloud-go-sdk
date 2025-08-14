// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentOutboundCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartAIAgentOutboundCallResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StartAIAgentOutboundCallResponseBody
	GetRequestId() *string
}

type StartAIAgentOutboundCallResponseBody struct {
	// example:
	//
	// *********296014bb58670940*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ***********-4417-BDB2************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAIAgentOutboundCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentOutboundCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartAIAgentOutboundCallResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartAIAgentOutboundCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAIAgentOutboundCallResponseBody) SetInstanceId(v string) *StartAIAgentOutboundCallResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartAIAgentOutboundCallResponseBody) SetRequestId(v string) *StartAIAgentOutboundCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAIAgentOutboundCallResponseBody) Validate() error {
	return dara.Validate(s)
}
