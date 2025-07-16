// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResourceResponseBody
	GetRequestId() *string
	SetResourceId(v string) *UpdateResourceResponseBody
	GetResourceId() *string
	SetResourceName(v string) *UpdateResourceResponseBody
	GetResourceName() *string
}

type UpdateResourceResponseBody struct {
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
	// The name of the resource group.
	//
	// example:
	//
	// iot
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourceResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateResourceResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetResourceId(v string) *UpdateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetResourceName(v string) *UpdateResourceResponseBody {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
