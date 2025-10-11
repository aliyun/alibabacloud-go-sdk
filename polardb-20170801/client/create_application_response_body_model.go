// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationResponseBody
	GetApplicationId() *string
	SetComponents(v []*CreateApplicationResponseBodyComponents) *CreateApplicationResponseBody
	GetComponents() []*CreateApplicationResponseBodyComponents
	SetDescription(v string) *CreateApplicationResponseBody
	GetDescription() *string
	SetOrderId(v string) *CreateApplicationResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateApplicationResponseBody
	GetRequestId() *string
	SetResourceAvailable(v bool) *CreateApplicationResponseBody
	GetResourceAvailable() *bool
	SetResourceGroupId(v string) *CreateApplicationResponseBody
	GetResourceGroupId() *string
}

type CreateApplicationResponseBody struct {
	// example:
	//
	// pa-********************
	ApplicationId *string                                    `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Components    []*CreateApplicationResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// myapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2148126708*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ResourceAvailable *bool `json:"ResourceAvailable,omitempty" xml:"ResourceAvailable,omitempty"`
	// example:
	//
	// rg-********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationResponseBody) GetComponents() []*CreateApplicationResponseBodyComponents {
	return s.Components
}

func (s *CreateApplicationResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationResponseBody) GetResourceAvailable() *bool {
	return s.ResourceAvailable
}

func (s *CreateApplicationResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationResponseBody) SetApplicationId(v string) *CreateApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetComponents(v []*CreateApplicationResponseBodyComponents) *CreateApplicationResponseBody {
	s.Components = v
	return s
}

func (s *CreateApplicationResponseBody) SetDescription(v string) *CreateApplicationResponseBody {
	s.Description = &v
	return s
}

func (s *CreateApplicationResponseBody) SetOrderId(v string) *CreateApplicationResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetResourceAvailable(v bool) *CreateApplicationResponseBody {
	s.ResourceAvailable = &v
	return s
}

func (s *CreateApplicationResponseBody) SetResourceGroupId(v string) *CreateApplicationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationResponseBodyComponents struct {
	// example:
	//
	// pac-********************
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
}

func (s CreateApplicationResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyComponents) GetComponentId() *string {
	return s.ComponentId
}

func (s *CreateApplicationResponseBodyComponents) SetComponentId(v string) *CreateApplicationResponseBodyComponents {
	s.ComponentId = &v
	return s
}

func (s *CreateApplicationResponseBodyComponents) Validate() error {
	return dara.Validate(s)
}
