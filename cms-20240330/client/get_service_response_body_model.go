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
	// example:
	//
	// 264C3E89-BE6E-5F82-A484-CE9C2196C7DC
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Service   *GetServiceResponseBodyService `json:"service,omitempty" xml:"service,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type GetServiceResponseBodyService struct {
	// example:
	//
	// {"language":"java"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// 2025-05-13T03:32:55Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// by6rjzro2j@0fe8dfa799e5906
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// cwzxvuc6uo@4bc6b15ad81f166174ffb
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// demo-app
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// example:
	//
	// Running
	ServiceStatus *string `json:"serviceStatus,omitempty" xml:"serviceStatus,omitempty"`
	// example:
	//
	// TRACE
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
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

func (s *GetServiceResponseBodyService) SetWorkspace(v string) *GetServiceResponseBodyService {
	s.Workspace = &v
	return s
}

func (s *GetServiceResponseBodyService) Validate() error {
	return dara.Validate(s)
}
