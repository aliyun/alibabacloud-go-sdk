// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadBalancerLeaveSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *LoadBalancerLeaveSecurityGroupResponseBody
	GetJobId() *string
	SetRequestId(v string) *LoadBalancerLeaveSecurityGroupResponseBody
	GetRequestId() *string
}

type LoadBalancerLeaveSecurityGroupResponseBody struct {
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 51c5b627-3500-487c-b17d-5cc583f0****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC0C96E4-7CCB-599C-9329-3A5DB6FF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoadBalancerLeaveSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoadBalancerLeaveSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBalancerLeaveSecurityGroupResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *LoadBalancerLeaveSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoadBalancerLeaveSecurityGroupResponseBody) SetJobId(v string) *LoadBalancerLeaveSecurityGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupResponseBody) SetRequestId(v string) *LoadBalancerLeaveSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoadBalancerLeaveSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
