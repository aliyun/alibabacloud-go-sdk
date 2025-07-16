// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateResourceInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *UpdateResourceInstanceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *UpdateResourceInstanceResponseBody
	GetResourceId() *string
}

type UpdateResourceInstanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// eas-i-asdasdasd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// eas-r-asdasdasd
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s UpdateResourceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateResourceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceInstanceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateResourceInstanceResponseBody) SetInstanceId(v string) *UpdateResourceInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceInstanceResponseBody) SetRequestId(v string) *UpdateResourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceInstanceResponseBody) SetResourceId(v string) *UpdateResourceInstanceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
