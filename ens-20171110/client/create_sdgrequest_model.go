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
	SetSize(v string) *CreateSDGRequest
	GetSize() *string
}

type CreateSDGRequest struct {
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The description of the SDG.
	//
	// >  We recommend that you specify this parameter in details for subsequent queries.
	//
	// example:
	//
	// Testing SDGs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskType    *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the SDG from which you want to create an SDG.
	//
	// >
	//
	// 	- The first time you create an SDG, the **FromSDGId*	- parameter is empty.
	//
	// 	- If the value of the **FromSDGId*	- parameter is invalid or does not correspond to an original disk, an error is reported.
	//
	// 	- If the value of the **FromSDGId*	- parameter is not empty, you have created an SDG, and the operation is performed on the existing SDG.
	//
	// example:
	//
	// sdg-xxxxx
	FromSDGId *string `json:"FromSDGId,omitempty" xml:"FromSDGId,omitempty"`
	// The ID of the AIC instance. You can call the [DescribeARMServerInstances](~~DescribeARMServerInstances~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aic-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum capacity of the SDG. Unit: GB.
	//
	// >
	//
	// 	- To save costs, we recommend that you specify this parameter based on your business requirements.
	//
	// 	- The first time that you create an SDG, the **Size*	- parameter is required.
	//
	// 	- When the amount of data increases, you can pass a new **Size*	- parameter for resizing. If the value of the new **Size*	- parameter is greater than the value of the old **Size*	- parameter, the disk size of the SDG is increased to the size that is specified by the new **Size*	- parameter. If the value of the new **Size*	- parameter is empty or smaller than that of the old **Size*	- parameter, no operation is performed.
	//
	// example:
	//
	// 20
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

func (s *CreateSDGRequest) SetSize(v string) *CreateSDGRequest {
	s.Size = &v
	return s
}

func (s *CreateSDGRequest) Validate() error {
	return dara.Validate(s)
}
