// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnSLSRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataCenter(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest
	GetDataCenter() *string
	SetDomainName(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetProjectName(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest
	GetProjectName() *string
	SetSLSLogStore(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest
	GetSLSLogStore() *string
	SetSLSProject(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest
	GetSLSProject() *string
	SetSLSRegion(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest
	GetSLSRegion() *string
	SetSamplingRate(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest
	GetSamplingRate() *string
}

type UpdateDcdnSLSRealtimeLogDeliveryRequest struct {
	// The region from which logs are collected.
	//
	// 	- **cn**: Chinese mainland
	//
	// 	- **sg**: Singapore
	//
	// 	- **in**: India
	//
	// 	- **eu**: Europe
	//
	// 	- **us**: United States
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The domain names from which logs were collected. You can specify one or more domain names. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-cn
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	// The name of the log file.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-cn
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region to which logs were delivered.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SLSRegion *string `json:"SLSRegion,omitempty" xml:"SLSRegion,omitempty"`
	// The sampling rate.
	//
	// example:
	//
	// 1.0
	SamplingRate *string `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
}

func (s UpdateDcdnSLSRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSLSRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) GetSLSLogStore() *string {
	return s.SLSLogStore
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) GetSLSProject() *string {
	return s.SLSProject
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) GetSamplingRate() *string {
	return s.SamplingRate
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetDataCenter(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.DataCenter = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetDomainName(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetProjectName(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSLSLogStore(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SLSLogStore = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSLSProject(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SLSProject = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSLSRegion(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SLSRegion = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) SetSamplingRate(v string) *UpdateDcdnSLSRealtimeLogDeliveryRequest {
	s.SamplingRate = &v
	return s
}

func (s *UpdateDcdnSLSRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
