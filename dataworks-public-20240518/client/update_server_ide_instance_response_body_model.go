// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerIdeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateServerIdeInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *UpdateServerIdeInstanceResponseBody
	GetRequestId() *string
}

type UpdateServerIdeInstanceResponseBody struct {
	// The personal development environment instance ID.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServerIdeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateServerIdeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServerIdeInstanceResponseBody) SetInstanceId(v string) *UpdateServerIdeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateServerIdeInstanceResponseBody) SetRequestId(v string) *UpdateServerIdeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServerIdeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
