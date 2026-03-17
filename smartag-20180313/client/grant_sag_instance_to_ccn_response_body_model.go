// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSagInstanceToCcnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GrantSagInstanceToCcnResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GrantSagInstanceToCcnResponseBody
	GetRequestId() *string
}

type GrantSagInstanceToCcnResponseBody struct {
	// The ID of the permission policy.
	//
	// example:
	//
	// sgc-6z21oj0vjjrx6s****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6E1674AC-083C-4031-B047-7A66E418E0C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantSagInstanceToCcnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantSagInstanceToCcnResponseBody) GoString() string {
	return s.String()
}

func (s *GrantSagInstanceToCcnResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GrantSagInstanceToCcnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantSagInstanceToCcnResponseBody) SetInstanceId(v string) *GrantSagInstanceToCcnResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GrantSagInstanceToCcnResponseBody) SetRequestId(v string) *GrantSagInstanceToCcnResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantSagInstanceToCcnResponseBody) Validate() error {
	return dara.Validate(s)
}
