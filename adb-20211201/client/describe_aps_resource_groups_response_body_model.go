// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeApsResourceGroupsResponseBodyData) *DescribeApsResourceGroupsResponseBody
	GetData() *DescribeApsResourceGroupsResponseBodyData
	SetHttpStatusCode(v int64) *DescribeApsResourceGroupsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DescribeApsResourceGroupsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApsResourceGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApsResourceGroupsResponseBody
	GetSuccess() *bool
}

type DescribeApsResourceGroupsResponseBody struct {
	// The queried resource groups.
	Data *DescribeApsResourceGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// 	- If the request was successful, a success message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FC370D7-1D4C-5A8E-805E-F73366382C66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeApsResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponseBody) GetData() *DescribeApsResourceGroupsResponseBodyData {
	return s.Data
}

func (s *DescribeApsResourceGroupsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeApsResourceGroupsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApsResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsResourceGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApsResourceGroupsResponseBody) SetData(v *DescribeApsResourceGroupsResponseBodyData) *DescribeApsResourceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetHttpStatusCode(v int64) *DescribeApsResourceGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetMessage(v string) *DescribeApsResourceGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetRequestId(v string) *DescribeApsResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) SetSuccess(v bool) *DescribeApsResourceGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApsResourceGroupsResponseBodyData struct {
	// The queried resource groups.
	ResourceGroups []*DescribeApsResourceGroupsResponseBodyDataResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// The step size of resources. Unit: AnalyticDB compute units (ACUs).
	//
	// 	- If the value of GroupType is **Interactive**, 16 is returned.
	//
	// 	- If the value of GroupType is **Job**, 8 is returned.
	//
	// example:
	//
	// 8
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApsResourceGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsResourceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponseBodyData) GetResourceGroups() []*DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	return s.ResourceGroups
}

func (s *DescribeApsResourceGroupsResponseBodyData) GetStep() *int64 {
	return s.Step
}

func (s *DescribeApsResourceGroupsResponseBodyData) SetResourceGroups(v []*DescribeApsResourceGroupsResponseBodyDataResourceGroups) *DescribeApsResourceGroupsResponseBodyData {
	s.ResourceGroups = v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyData) SetStep(v int64) *DescribeApsResourceGroupsResponseBodyData {
	s.Step = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyData) Validate() error {
	if s.ResourceGroups != nil {
		for _, item := range s.ResourceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApsResourceGroupsResponseBodyDataResourceGroups struct {
	// Indicates whether the resource group is available. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Available *bool    `json:"Available,omitempty" xml:"Available,omitempty"`
	CuOptions []*int64 `json:"CuOptions,omitempty" xml:"CuOptions,omitempty" type:"Repeated"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- **Interactive**
	//
	// 	- **Job**
	//
	// >  For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/428610.html).
	//
	// example:
	//
	// Job
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The amount of remaining computing resources. Unit: ACUs.
	//
	// example:
	//
	// 512
	LeftComputeResource *int32 `json:"LeftComputeResource,omitempty" xml:"LeftComputeResource,omitempty"`
	// The maximum amount of reserved computing resources. Unit: ACUs.
	//
	// 	- If the value of GroupType is **Interactive**, the amount of reserved computing resources that are not allocated in the cluster is returned in increments of 16 ACUs.
	//
	// 	- If the value of GroupType is **Job**, the amount of reserved computing resources that are not allocated in the cluster is returned in increments of 8 ACUs.
	//
	// example:
	//
	// 512
	MaxComputeResource *int32 `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// The minimum amount of reserved computing resources. Unit: ACUs.
	//
	// 	- If the value of GroupType is **Interactive**, 16 is returned.
	//
	// 	- If the value of GroupType is **Job**, 0 is returned.
	//
	// example:
	//
	// 0
	MinComputeResource *int32 `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
}

func (s DescribeApsResourceGroupsResponseBodyDataResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsResourceGroupsResponseBodyDataResourceGroups) GoString() string {
	return s.String()
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) GetAvailable() *bool {
	return s.Available
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) GetCuOptions() []*int64 {
	return s.CuOptions
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) GetLeftComputeResource() *int32 {
	return s.LeftComputeResource
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) GetMaxComputeResource() *int32 {
	return s.MaxComputeResource
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) GetMinComputeResource() *int32 {
	return s.MinComputeResource
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetAvailable(v bool) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.Available = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetCuOptions(v []*int64) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.CuOptions = v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetGroupName(v string) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetGroupType(v string) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.GroupType = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetLeftComputeResource(v int32) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.LeftComputeResource = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetMaxComputeResource(v int32) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.MaxComputeResource = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) SetMinComputeResource(v int32) *DescribeApsResourceGroupsResponseBodyDataResourceGroups {
	s.MinComputeResource = &v
	return s
}

func (s *DescribeApsResourceGroupsResponseBodyDataResourceGroups) Validate() error {
	return dara.Validate(s)
}
