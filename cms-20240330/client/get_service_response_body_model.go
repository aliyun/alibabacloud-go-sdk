// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceResponseBody
	GetRequestId() *string
	SetService(v *GetServiceResponseBodyService) *GetServiceResponseBody
	GetService() *GetServiceResponseBodyService
}

type GetServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 264C3E89-BE6E-5F82-A484-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The service object.
	Service *GetServiceResponseBodyService `json:"service,omitempty" xml:"service,omitempty" type:"Struct"`
}

func (s GetServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceResponseBody) GetService() *GetServiceResponseBodyService {
	return s.Service
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceResponseBody) SetService(v *GetServiceResponseBodyService) *GetServiceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceResponseBody) Validate() error {
	if s.Service != nil {
		if err := s.Service.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceResponseBodyService struct {
	// The extended information.
	//
	// example:
	//
	// {"language":"java"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// The time when the service was created.
	//
	// example:
	//
	// 2025-05-13T03:32:55Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// demo应用
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The ID of the ARMS application that is compatible with earlier versions.
	//
	// example:
	//
	// by6rjzro2j@0fe8dfa799e5906
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekxxzuad5zzzz
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// cwzxvuc6uo@4bc6b15ad81f166174ffb
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// demo-app
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The service status. This parameter is valid only when serviceType is set to RUM.
	//
	// example:
	//
	// Running
	ServiceStatus *string `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	// The service type.
	//
	// example:
	//
	// TRACE
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The array of tags.
	Tags []*GetServiceResponseBodyServiceTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The workspace name.
	//
	// example:
	//
	// default-cms-1106439496876715-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetServiceResponseBodyService) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyService) GetAttributes() *string {
	return s.Attributes
}

func (s *GetServiceResponseBodyService) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetServiceResponseBodyService) GetDescription() *string {
	return s.Description
}

func (s *GetServiceResponseBodyService) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetServiceResponseBodyService) GetPid() *string {
	return s.Pid
}

func (s *GetServiceResponseBodyService) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceResponseBodyService) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetServiceResponseBodyService) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceResponseBodyService) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceResponseBodyService) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *GetServiceResponseBodyService) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceResponseBodyService) GetTags() []*GetServiceResponseBodyServiceTags {
	return s.Tags
}

func (s *GetServiceResponseBodyService) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetServiceResponseBodyService) SetAttributes(v string) *GetServiceResponseBodyService {
	s.Attributes = &v
	return s
}

func (s *GetServiceResponseBodyService) SetCreateTime(v string) *GetServiceResponseBodyService {
	s.CreateTime = &v
	return s
}

func (s *GetServiceResponseBodyService) SetDescription(v string) *GetServiceResponseBodyService {
	s.Description = &v
	return s
}

func (s *GetServiceResponseBodyService) SetDisplayName(v string) *GetServiceResponseBodyService {
	s.DisplayName = &v
	return s
}

func (s *GetServiceResponseBodyService) SetPid(v string) *GetServiceResponseBodyService {
	s.Pid = &v
	return s
}

func (s *GetServiceResponseBodyService) SetRegionId(v string) *GetServiceResponseBodyService {
	s.RegionId = &v
	return s
}

func (s *GetServiceResponseBodyService) SetResourceGroupId(v string) *GetServiceResponseBodyService {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceResponseBodyService) SetServiceId(v string) *GetServiceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceResponseBodyService) SetServiceName(v string) *GetServiceResponseBodyService {
	s.ServiceName = &v
	return s
}

func (s *GetServiceResponseBodyService) SetServiceStatus(v string) *GetServiceResponseBodyService {
	s.ServiceStatus = &v
	return s
}

func (s *GetServiceResponseBodyService) SetServiceType(v string) *GetServiceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceResponseBodyService) SetTags(v []*GetServiceResponseBodyServiceTags) *GetServiceResponseBodyService {
	s.Tags = v
	return s
}

func (s *GetServiceResponseBodyService) SetWorkspace(v string) *GetServiceResponseBodyService {
	s.Workspace = &v
	return s
}

func (s *GetServiceResponseBodyService) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceResponseBodyServiceTags struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prod
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetServiceResponseBodyServiceTags) String() string {
	return dara.Prettify(s)
}

func (s GetServiceResponseBodyServiceTags) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceTags) GetKey() *string {
	return s.Key
}

func (s *GetServiceResponseBodyServiceTags) GetValue() *string {
	return s.Value
}

func (s *GetServiceResponseBodyServiceTags) SetKey(v string) *GetServiceResponseBodyServiceTags {
	s.Key = &v
	return s
}

func (s *GetServiceResponseBodyServiceTags) SetValue(v string) *GetServiceResponseBodyServiceTags {
	s.Value = &v
	return s
}

func (s *GetServiceResponseBodyServiceTags) Validate() error {
	return dara.Validate(s)
}
