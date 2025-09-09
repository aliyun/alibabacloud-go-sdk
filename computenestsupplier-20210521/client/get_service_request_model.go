// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterAliUid(v bool) *GetServiceRequest
	GetFilterAliUid() *bool
	SetRegionId(v string) *GetServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *GetServiceRequest
	GetServiceInstanceId() *string
	SetServiceName(v string) *GetServiceRequest
	GetServiceName() *string
	SetServiceVersion(v string) *GetServiceRequest
	GetServiceVersion() *string
	SetSharedAccountType(v string) *GetServiceRequest
	GetSharedAccountType() *string
	SetShowDetail(v []*string) *GetServiceRequest
	GetShowDetail() []*string
}

type GetServiceRequest struct {
	// Specifies whether to filter information based on Alibaba Cloud account IDs.
	//
	// example:
	//
	// false
	FilterAliUid *bool `json:"FilterAliUid,omitempty" xml:"FilterAliUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-4ee86df83fd948******
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The Service Instance Id.
	//
	// example:
	//
	// si-85b1exxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The Service Name.
	//
	// example:
	//
	// Wordpress社区版
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The share type of the service. Default value: SharedAccount. Valid values:
	//
	// 	- SharedAccount: The service is shared by multiple accounts.
	//
	// 	- Resell: The service is distributed.
	//
	// example:
	//
	// SharedAccount
	SharedAccountType *string `json:"SharedAccountType,omitempty" xml:"SharedAccountType,omitempty"`
	// The information that you want to query.
	ShowDetail []*string `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty" type:"Repeated"`
}

func (s GetServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) GetFilterAliUid() *bool {
	return s.FilterAliUid
}

func (s *GetServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *GetServiceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceRequest) GetSharedAccountType() *string {
	return s.SharedAccountType
}

func (s *GetServiceRequest) GetShowDetail() []*string {
	return s.ShowDetail
}

func (s *GetServiceRequest) SetFilterAliUid(v bool) *GetServiceRequest {
	s.FilterAliUid = &v
	return s
}

func (s *GetServiceRequest) SetRegionId(v string) *GetServiceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceRequest) SetServiceId(v string) *GetServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceInstanceId(v string) *GetServiceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceName(v string) *GetServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceRequest) SetServiceVersion(v string) *GetServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceRequest) SetSharedAccountType(v string) *GetServiceRequest {
	s.SharedAccountType = &v
	return s
}

func (s *GetServiceRequest) SetShowDetail(v []*string) *GetServiceRequest {
	s.ShowDetail = v
	return s
}

func (s *GetServiceRequest) Validate() error {
	return dara.Validate(s)
}
