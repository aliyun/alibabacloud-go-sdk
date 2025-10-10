// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadBalancerJoinSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *LoadBalancerJoinSecurityGroupResponseBody
	GetJobId() *string
	SetRequestId(v string) *LoadBalancerJoinSecurityGroupResponseBody
	GetRequestId() *string
}

type LoadBalancerJoinSecurityGroupResponseBody struct {
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 8fe81f25-79a0-4fa0-9036-f2601fda****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D3B9AE45-F5DB-58E3-A4B5-EE58F1EC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoadBalancerJoinSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoadBalancerJoinSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBalancerJoinSecurityGroupResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *LoadBalancerJoinSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoadBalancerJoinSecurityGroupResponseBody) SetJobId(v string) *LoadBalancerJoinSecurityGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupResponseBody) SetRequestId(v string) *LoadBalancerJoinSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoadBalancerJoinSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
