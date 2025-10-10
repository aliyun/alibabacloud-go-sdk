// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateLoadBalancerAttributeResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody
	GetRequestId() *string
}

type UpdateLoadBalancerAttributeResponseBody struct {
	// The ID of the synchronous task.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerAttributeResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetJobId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) SetRequestId(v string) *UpdateLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLoadBalancerAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
