// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelDeploymentProfilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v string) *ListModelDeploymentProfilesResponseBody
	GetModelId() *string
	SetModelVersion(v string) *ListModelDeploymentProfilesResponseBody
	GetModelVersion() *string
	SetProfiles(v []*ListModelDeploymentProfilesResponseBodyProfiles) *ListModelDeploymentProfilesResponseBody
	GetProfiles() []*ListModelDeploymentProfilesResponseBodyProfiles
	SetRequestId(v string) *ListModelDeploymentProfilesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListModelDeploymentProfilesResponseBody
	GetTotalCount() *int64
}

type ListModelDeploymentProfilesResponseBody struct {
	// example:
	//
	// model-6wiou4ta20tgtq9lda
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// 1.0.0
	ModelVersion *string                                            `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	Profiles     []*ListModelDeploymentProfilesResponseBodyProfiles `json:"Profiles,omitempty" xml:"Profiles,omitempty" type:"Repeated"`
	// example:
	//
	// B6B54325-C98C-5937-87A3-2F96C07652EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelDeploymentProfilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelDeploymentProfilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelDeploymentProfilesResponseBody) GetModelId() *string {
	return s.ModelId
}

func (s *ListModelDeploymentProfilesResponseBody) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *ListModelDeploymentProfilesResponseBody) GetProfiles() []*ListModelDeploymentProfilesResponseBodyProfiles {
	return s.Profiles
}

func (s *ListModelDeploymentProfilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelDeploymentProfilesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelDeploymentProfilesResponseBody) SetModelId(v string) *ListModelDeploymentProfilesResponseBody {
	s.ModelId = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBody) SetModelVersion(v string) *ListModelDeploymentProfilesResponseBody {
	s.ModelVersion = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBody) SetProfiles(v []*ListModelDeploymentProfilesResponseBodyProfiles) *ListModelDeploymentProfilesResponseBody {
	s.Profiles = v
	return s
}

func (s *ListModelDeploymentProfilesResponseBody) SetRequestId(v string) *ListModelDeploymentProfilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBody) SetTotalCount(v int64) *ListModelDeploymentProfilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBody) Validate() error {
	if s.Profiles != nil {
		for _, item := range s.Profiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelDeploymentProfilesResponseBodyProfiles struct {
	// example:
	//
	// singlenode
	Category *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	Devices  []*ListModelDeploymentProfilesResponseBodyProfilesDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// sglang
	Framework     *string                                                   `json:"Framework,omitempty" xml:"Framework,omitempty"`
	Labels        map[string]*string                                        `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Members       []*ListModelDeploymentProfilesResponseBodyProfilesMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Optimizations map[string]*string                                        `json:"Optimizations,omitempty" xml:"Optimizations,omitempty"`
	// example:
	//
	// prf_4f73b31ae****23f9a04c6e83ebc78f
	ProfileId *string `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
	// example:
	//
	// singlenode-balanced
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
}

func (s ListModelDeploymentProfilesResponseBodyProfiles) String() string {
	return dara.Prettify(s)
}

func (s ListModelDeploymentProfilesResponseBodyProfiles) GoString() string {
	return s.String()
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetCategory() *string {
	return s.Category
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetDevices() []*ListModelDeploymentProfilesResponseBodyProfilesDevices {
	return s.Devices
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetFramework() *string {
	return s.Framework
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetMembers() []*ListModelDeploymentProfilesResponseBodyProfilesMembers {
	return s.Members
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetOptimizations() map[string]*string {
	return s.Optimizations
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetProfileId() *string {
	return s.ProfileId
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) GetScenario() *string {
	return s.Scenario
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetCategory(v string) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.Category = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetDevices(v []*ListModelDeploymentProfilesResponseBodyProfilesDevices) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.Devices = v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetFramework(v string) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.Framework = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetLabels(v map[string]*string) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.Labels = v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetMembers(v []*ListModelDeploymentProfilesResponseBodyProfilesMembers) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.Members = v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetOptimizations(v map[string]*string) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.Optimizations = v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetProfileId(v string) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.ProfileId = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) SetScenario(v string) *ListModelDeploymentProfilesResponseBodyProfiles {
	s.Scenario = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfiles) Validate() error {
	if s.Devices != nil {
		for _, item := range s.Devices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelDeploymentProfilesResponseBodyProfilesDevices struct {
	// example:
	//
	// NVIDIA
	DeviceCategory *string `json:"DeviceCategory,omitempty" xml:"DeviceCategory,omitempty"`
	// example:
	//
	// NVIDIA_L20C
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// L20C
	DisplayName   *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
}

func (s ListModelDeploymentProfilesResponseBodyProfilesDevices) String() string {
	return dara.Prettify(s)
}

func (s ListModelDeploymentProfilesResponseBodyProfilesDevices) GoString() string {
	return s.String()
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) GetDeviceCategory() *string {
	return s.DeviceCategory
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) SetDeviceCategory(v string) *ListModelDeploymentProfilesResponseBodyProfilesDevices {
	s.DeviceCategory = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) SetDeviceType(v string) *ListModelDeploymentProfilesResponseBodyProfilesDevices {
	s.DeviceType = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) SetDisplayName(v string) *ListModelDeploymentProfilesResponseBodyProfilesDevices {
	s.DisplayName = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) SetInstanceTypes(v []*string) *ListModelDeploymentProfilesResponseBodyProfilesDevices {
	s.InstanceTypes = v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesDevices) Validate() error {
	return dara.Validate(s)
}

type ListModelDeploymentProfilesResponseBodyProfilesMembers struct {
	// example:
	//
	// Default
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// example:
	//
	// { "script": "sglang serve --disaggregation-mode decode" }
	Meta map[string]interface{} `json:"Meta,omitempty" xml:"Meta,omitempty"`
}

func (s ListModelDeploymentProfilesResponseBodyProfilesMembers) String() string {
	return dara.Prettify(s)
}

func (s ListModelDeploymentProfilesResponseBodyProfilesMembers) GoString() string {
	return s.String()
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesMembers) GetMemberType() *string {
	return s.MemberType
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesMembers) GetMeta() map[string]interface{} {
	return s.Meta
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesMembers) SetMemberType(v string) *ListModelDeploymentProfilesResponseBodyProfilesMembers {
	s.MemberType = &v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesMembers) SetMeta(v map[string]interface{}) *ListModelDeploymentProfilesResponseBodyProfilesMembers {
	s.Meta = v
	return s
}

func (s *ListModelDeploymentProfilesResponseBodyProfilesMembers) Validate() error {
	return dara.Validate(s)
}
