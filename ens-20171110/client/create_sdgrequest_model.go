// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingCycle(v string) *CreateSDGRequest
	GetBillingCycle() *string
	SetDescription(v string) *CreateSDGRequest
	GetDescription() *string
	SetDiskType(v string) *CreateSDGRequest
	GetDiskType() *string
	SetFromSDGId(v string) *CreateSDGRequest
	GetFromSDGId() *string
	SetInstanceId(v string) *CreateSDGRequest
	GetInstanceId() *string
	SetPerformanceLevel(v int64) *CreateSDGRequest
	GetPerformanceLevel() *int64
	SetSize(v string) *CreateSDGRequest
	GetSize() *string
}

type CreateSDGRequest struct {
	// The billing logic for SDG child disks. This parameter specifies the billing method for the RBD child disks produced by this SDG. This parameter does not apply to standard cloud disks.
	//
	// Optional parameter values:
	//
	// - Hour.
	//
	// -  Day. This is the default value.
	//
	//
	// > - Parameter description: The longer the cycle, the lower the cost. If the disk is released before the billing cycle ends, a service fee is charged.
	//
	// example:
	//
	// Day
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The description of the SDG.
	//
	// > Provide a detailed description for easier future queries.
	//
	// example:
	//
	// Test SDG
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the source disk.
	//
	// Optional parameter values:
	//
	// -  standard: standard cloud disk. This is the default value.
	//
	// -  rbd: RBD disk.
	//
	// > - Limits: RBD disk SDGs and standard disk SDGs cannot inherit from each other.
	//
	// example:
	//
	// standard
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of an existing SDG.
	//
	// > - If you are creating an SDG for the first time, leave **FromSDGId*	- empty.
	//
	// >- If **FromSDGId*	- is invalid or the corresponding raw cloud disk cannot be found, an error is returned.
	//
	// >- If **FromSDGId*	- is not empty, this indicates that you are operating on an existing SDG rather than creating one for the first time.
	//
	// example:
	//
	// sdg-xxxxx
	FromSDGId *string `json:"FromSDGId,omitempty" xml:"FromSDGId,omitempty"`
	// The AIC instance ID. You can call the [DescribeARMServerInstances](~~DescribeARMServerInstances~~) operation to obtain the InstanceId value.
	//
	// This parameter is required.
	//
	// example:
	//
	// aic-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The SDG sharing ratio parameter. This parameter supports cloud disks with different maximum sharing ratios.
	//
	// example:
	//
	// 100
	PerformanceLevel *int64 `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The maximum capacity of the SDG. Unit: GB.
	//
	// >- To reduce costs, evaluate the required capacity as needed.
	//
	// >- If you are creating an SDG for the first time, **Size*	- is required.
	//
	// >- If the data volume increases later, pass a new **Size*	- value to expand the capacity. If **Size*	- is not empty and is greater than the original **Size*	- value of the SDG, the raw cloud disk of the SDG is expanded to the specified **Size**. If **Size*	- is empty or the specified **Size*	- is less than the original **Size*	- value of the SDG, no action is taken.
	//
	// The minimum tested usable value is 150 GB.
	//
	// example:
	//
	// 150
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSDGRequest) GoString() string {
	return s.String()
}

func (s *CreateSDGRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *CreateSDGRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSDGRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateSDGRequest) GetFromSDGId() *string {
	return s.FromSDGId
}

func (s *CreateSDGRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSDGRequest) GetPerformanceLevel() *int64 {
	return s.PerformanceLevel
}

func (s *CreateSDGRequest) GetSize() *string {
	return s.Size
}

func (s *CreateSDGRequest) SetBillingCycle(v string) *CreateSDGRequest {
	s.BillingCycle = &v
	return s
}

func (s *CreateSDGRequest) SetDescription(v string) *CreateSDGRequest {
	s.Description = &v
	return s
}

func (s *CreateSDGRequest) SetDiskType(v string) *CreateSDGRequest {
	s.DiskType = &v
	return s
}

func (s *CreateSDGRequest) SetFromSDGId(v string) *CreateSDGRequest {
	s.FromSDGId = &v
	return s
}

func (s *CreateSDGRequest) SetInstanceId(v string) *CreateSDGRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSDGRequest) SetPerformanceLevel(v int64) *CreateSDGRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateSDGRequest) SetSize(v string) *CreateSDGRequest {
	s.Size = &v
	return s
}

func (s *CreateSDGRequest) Validate() error {
	return dara.Validate(s)
}
