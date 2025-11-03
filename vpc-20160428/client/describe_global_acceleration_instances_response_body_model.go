// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalAccelerationInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalAccelerationInstances(v *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) *DescribeGlobalAccelerationInstancesResponseBody
	GetGlobalAccelerationInstances() *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances
	SetPageNumber(v int32) *DescribeGlobalAccelerationInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGlobalAccelerationInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGlobalAccelerationInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGlobalAccelerationInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeGlobalAccelerationInstancesResponseBody struct {
	// The details of the GA instances.
	GlobalAccelerationInstances *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances `json:"GlobalAccelerationInstances,omitempty" xml:"GlobalAccelerationInstances,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6B4EE38D-C75B-4E1F-844E-863A94430676
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGlobalAccelerationInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) GetGlobalAccelerationInstances() *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances {
	return s.GlobalAccelerationInstances
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) SetGlobalAccelerationInstances(v *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) *DescribeGlobalAccelerationInstancesResponseBody {
	s.GlobalAccelerationInstances = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) SetPageNumber(v int32) *DescribeGlobalAccelerationInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) SetPageSize(v int32) *DescribeGlobalAccelerationInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) SetRequestId(v string) *DescribeGlobalAccelerationInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) SetTotalCount(v int32) *DescribeGlobalAccelerationInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBody) Validate() error {
	if s.GlobalAccelerationInstances != nil {
		if err := s.GlobalAccelerationInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances struct {
	GlobalAccelerationInstance []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance `json:"GlobalAccelerationInstance,omitempty" xml:"GlobalAccelerationInstance,omitempty" type:"Repeated"`
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) GetGlobalAccelerationInstance() []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	return s.GlobalAccelerationInstance
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) SetGlobalAccelerationInstance(v []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances {
	s.GlobalAccelerationInstance = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstances) Validate() error {
	if s.GlobalAccelerationInstance != nil {
		for _, item := range s.GlobalAccelerationInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	// The acceleration area of the GA instance.
	//
	// example:
	//
	// china-mainland
	AccelerationLocation *string `json:"AccelerationLocation,omitempty" xml:"AccelerationLocation,omitempty"`
	// The details about the backend servers of the GA instance.
	BackendServers *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	// The maximum bandwidth of the GA instance.
	//
	// example:
	//
	// 10
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth type of the GA instance.
	//
	// 	- **Sharing**
	//
	// 	- **Exclusive*	- (default)
	//
	// example:
	//
	// Exclusive
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The billing method of the GA instance.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the GA instance was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-07-05T03:39:31Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the GA instance.
	//
	// example:
	//
	// apiDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2018-08-05T16:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1x99kj7kl1ziw5x****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	// The billing method of the GA instance.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The public IP address of the dedicated GA instance.
	//
	// example:
	//
	// 47.xx.xx.99
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the GA instance.
	//
	// example:
	//
	// instanceName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The public IP address.
	PublicIpAddresses *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty" type:"Struct"`
	// The region ID of the GA instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service area of the GA instance.
	//
	// example:
	//
	// china-mainland
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// The status of the GA instance.
	//
	// 	- **Available**
	//
	// 	- **Inuse**
	//
	// 	- **Associating**
	//
	// 	- **Unassociating**
	//
	// example:
	//
	// InUse
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetAccelerationLocation() *string {
	return s.AccelerationLocation
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetBackendServers() *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers {
	return s.BackendServers
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetName() *string {
	return s.Name
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetPublicIpAddresses() *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses {
	return s.PublicIpAddresses
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetAccelerationLocation(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.AccelerationLocation = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetBackendServers(v *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.BackendServers = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetBandwidth(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.Bandwidth = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetBandwidthType(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.BandwidthType = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetChargeType(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetCreationTime(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetDescription(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.Description = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetExpiredTime(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetGlobalAccelerationInstanceId(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetInternetChargeType(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetIpAddress(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.IpAddress = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetName(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.Name = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetPublicIpAddresses(v *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.PublicIpAddresses = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetRegionId(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetServiceLocation(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.ServiceLocation = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) SetStatus(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance {
	s.Status = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstance) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	if s.PublicIpAddresses != nil {
		if err := s.PublicIpAddresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers struct {
	BackendServer []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers) GetBackendServer() []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer {
	return s.BackendServer
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers) SetBackendServer(v []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers {
	s.BackendServer = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServers) Validate() error {
	if s.BackendServer != nil {
		for _, item := range s.BackendServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer struct {
	// The region where the backend servers are deployed.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the backend server.
	//
	// example:
	//
	// i-2zeg83zvn5d4ed4y****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The IP address of the backend server.
	//
	// example:
	//
	// 172.xx.xx.109
	ServerIpAddress *string `json:"ServerIpAddress,omitempty" xml:"ServerIpAddress,omitempty"`
	// The type of the backend server.
	//
	// 	- **EcsInstance**: Elastic Compute Service (ECS) instance
	//
	// 	- **SlbInstance**: Server Load Balancer (SLB) instance
	//
	// example:
	//
	// EcsInstance
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) GetServerIpAddress() *string {
	return s.ServerIpAddress
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) SetRegionId(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) SetServerId(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) SetServerIpAddress(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer {
	s.ServerIpAddress = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) SetServerType(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer {
	s.ServerType = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstanceBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses struct {
	PublicIpAddress []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty" type:"Repeated"`
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses) GetPublicIpAddress() []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress {
	return s.PublicIpAddress
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses) SetPublicIpAddress(v []*DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses {
	s.PublicIpAddress = v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddresses) Validate() error {
	if s.PublicIpAddress != nil {
		for _, item := range s.PublicIpAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress struct {
	// The ID of the public IP address of the GA instance.
	//
	// example:
	//
	// eip-bp19yqraac4w3y0jd****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The public IP address of the GA instance.
	//
	// example:
	//
	// 12.xx.xx.78
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) SetAllocationId(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) SetIpAddress(v string) *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesResponseBodyGlobalAccelerationInstancesGlobalAccelerationInstancePublicIpAddressesPublicIpAddress) Validate() error {
	return dara.Validate(s)
}
