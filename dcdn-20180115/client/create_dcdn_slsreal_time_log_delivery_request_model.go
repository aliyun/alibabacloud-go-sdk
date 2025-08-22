// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnSLSRealTimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetBusinessType() *string
	SetDataCenter(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetDataCenter() *string
	SetDomainName(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetDomainName() *string
	SetProjectName(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetProjectName() *string
	SetSLSLogStore(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetSLSLogStore() *string
	SetSLSProject(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetSLSProject() *string
	SetSLSRegion(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetSLSRegion() *string
	SetSamplingRate(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest
	GetSamplingRate() *string
}

type CreateDcdnSLSRealTimeLogDeliveryRequest struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of Dynamic Content Delivery Network (DCDN) points of presence (POPs)
	//
	// 	- **cdn_log_origin**: back-to-origin logs
	//
	// 	- **cdn_log_er**: EdgeRoutine logs
	//
	// example:
	//
	// cdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The data center. Valid values:
	//
	// 	- cn: China
	//
	// 	- sg: Singapore
	//
	// 	- eu: Europe
	//
	// 	- us: United States
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
	// The name of a real-time log delivery project.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the Log Service Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	SLSLogStore *string `json:"SLSLogStore,omitempty" xml:"SLSLogStore,omitempty"`
	// The name of the Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyundoc
	SLSProject *string `json:"SLSProject,omitempty" xml:"SLSProject,omitempty"`
	// The region to which real-time logs are delivered.
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

func (s CreateDcdnSLSRealTimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetSLSLogStore() *string {
	return s.SLSLogStore
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetSLSProject() *string {
	return s.SLSProject
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetSLSRegion() *string {
	return s.SLSRegion
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) GetSamplingRate() *string {
	return s.SamplingRate
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetBusinessType(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetDataCenter(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.DataCenter = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetDomainName(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetProjectName(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSLSLogStore(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SLSLogStore = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSLSProject(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SLSProject = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSLSRegion(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SLSRegion = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) SetSamplingRate(v string) *CreateDcdnSLSRealTimeLogDeliveryRequest {
	s.SamplingRate = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
