// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorComputeSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorComputeSummaryRequest
	GetClusterId() *string
	SetComponentInfo(v *GetDoctorComputeSummaryRequestComponentInfo) *GetDoctorComputeSummaryRequest
	GetComponentInfo() *GetDoctorComputeSummaryRequestComponentInfo
	SetDateTime(v string) *GetDoctorComputeSummaryRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorComputeSummaryRequest
	GetRegionId() *string
}

type GetDoctorComputeSummaryRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The resource information, which is used to filter the results.
	ComponentInfo *GetDoctorComputeSummaryRequestComponentInfo `json:"ComponentInfo,omitempty" xml:"ComponentInfo,omitempty" type:"Struct"`
	// Specify the date in the ISO 8601 standard. For example, 2023-01-01 represents January 1, 2023.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDoctorComputeSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorComputeSummaryRequest) GetComponentInfo() *GetDoctorComputeSummaryRequestComponentInfo {
	return s.ComponentInfo
}

func (s *GetDoctorComputeSummaryRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorComputeSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorComputeSummaryRequest) SetClusterId(v string) *GetDoctorComputeSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorComputeSummaryRequest) SetComponentInfo(v *GetDoctorComputeSummaryRequestComponentInfo) *GetDoctorComputeSummaryRequest {
	s.ComponentInfo = v
	return s
}

func (s *GetDoctorComputeSummaryRequest) SetDateTime(v string) *GetDoctorComputeSummaryRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorComputeSummaryRequest) SetRegionId(v string) *GetDoctorComputeSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorComputeSummaryRequest) Validate() error {
	if s.ComponentInfo != nil {
		if err := s.ComponentInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDoctorComputeSummaryRequestComponentInfo struct {
	// Set the filter condition name based on the value of ComponentType. For example, if you set ComponentType to queue, you can specify a specific queue name to obtain the resource usage of a specific queue.
	//
	// example:
	//
	// MAPREDUCE
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The resource type for filtering. Valid values:
	//
	// 	- engine: filters results by engine.
	//
	// 	- queue: filters results by queue.
	//
	// 	- cluster: displays the results at the cluster level.
	//
	// If you do not specify this parameter, the information at the cluster level is displayed by default.
	//
	// example:
	//
	// engine
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
}

func (s GetDoctorComputeSummaryRequestComponentInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryRequestComponentInfo) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryRequestComponentInfo) GetComponentName() *string {
	return s.ComponentName
}

func (s *GetDoctorComputeSummaryRequestComponentInfo) GetComponentType() *string {
	return s.ComponentType
}

func (s *GetDoctorComputeSummaryRequestComponentInfo) SetComponentName(v string) *GetDoctorComputeSummaryRequestComponentInfo {
	s.ComponentName = &v
	return s
}

func (s *GetDoctorComputeSummaryRequestComponentInfo) SetComponentType(v string) *GetDoctorComputeSummaryRequestComponentInfo {
	s.ComponentType = &v
	return s
}

func (s *GetDoctorComputeSummaryRequestComponentInfo) Validate() error {
	return dara.Validate(s)
}
