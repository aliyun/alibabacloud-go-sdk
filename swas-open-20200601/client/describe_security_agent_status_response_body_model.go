// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientStatus(v string) *DescribeSecurityAgentStatusResponseBody
	GetClientStatus() *string
	SetRequestId(v string) *DescribeSecurityAgentStatusResponseBody
	GetRequestId() *string
}

type DescribeSecurityAgentStatusResponseBody struct {
	// The status of the Security Center agent. Valid values:
	//
	// - pause: The Security Center agent suspends protection for your server.
	//
	// - online: The Security Center agent is protecting your server.
	//
	// - offline: The Security Center agent does not protect your server.
	//
	// example:
	//
	// online
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityAgentStatusResponseBody) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeSecurityAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityAgentStatusResponseBody) SetClientStatus(v string) *DescribeSecurityAgentStatusResponseBody {
	s.ClientStatus = &v
	return s
}

func (s *DescribeSecurityAgentStatusResponseBody) SetRequestId(v string) *DescribeSecurityAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityAgentStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
