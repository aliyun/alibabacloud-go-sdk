// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingDataPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateRenderingDataPackageRequest
	GetCategory() *string
	SetDescription(v string) *CreateRenderingDataPackageRequest
	GetDescription() *string
	SetInstanceBillingCycle(v string) *CreateRenderingDataPackageRequest
	GetInstanceBillingCycle() *string
	SetRenderingInstanceId(v string) *CreateRenderingDataPackageRequest
	GetRenderingInstanceId() *string
}

type CreateRenderingDataPackageRequest struct {
	// The data type. Valid values:
	//
	// - Package (default)
	//
	// - Baseline
	//
	// example:
	//
	// Package
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description.
	//
	// example:
	//
	// testdescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing type. Valid values:
	//
	// - Hour: hourly.
	//
	// example:
	//
	// Hour
	InstanceBillingCycle *string `json:"InstanceBillingCycle,omitempty" xml:"InstanceBillingCycle,omitempty"`
	// The instance ID of the Cloud Application Service. Note: This operation applies only to crs.cp.	- series instance types.
	//
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s CreateRenderingDataPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingDataPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingDataPackageRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateRenderingDataPackageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRenderingDataPackageRequest) GetInstanceBillingCycle() *string {
	return s.InstanceBillingCycle
}

func (s *CreateRenderingDataPackageRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *CreateRenderingDataPackageRequest) SetCategory(v string) *CreateRenderingDataPackageRequest {
	s.Category = &v
	return s
}

func (s *CreateRenderingDataPackageRequest) SetDescription(v string) *CreateRenderingDataPackageRequest {
	s.Description = &v
	return s
}

func (s *CreateRenderingDataPackageRequest) SetInstanceBillingCycle(v string) *CreateRenderingDataPackageRequest {
	s.InstanceBillingCycle = &v
	return s
}

func (s *CreateRenderingDataPackageRequest) SetRenderingInstanceId(v string) *CreateRenderingDataPackageRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *CreateRenderingDataPackageRequest) Validate() error {
	return dara.Validate(s)
}
