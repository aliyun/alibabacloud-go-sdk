// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetShowDetails(v []*string) *GetServiceRequest
	GetShowDetails() []*string
}

type GetServiceRequest struct {
	// Region Id.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance id.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// Wordpress
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// Whether to disclose service details.
	ShowDetails []*string `json:"ShowDetails,omitempty" xml:"ShowDetails,omitempty" type:"Repeated"`
}

func (s GetServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
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

func (s *GetServiceRequest) GetShowDetails() []*string {
	return s.ShowDetails
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

func (s *GetServiceRequest) SetShowDetails(v []*string) *GetServiceRequest {
	s.ShowDetails = v
	return s
}

func (s *GetServiceRequest) Validate() error {
	return dara.Validate(s)
}
