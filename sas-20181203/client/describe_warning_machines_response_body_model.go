// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarningMachinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeWarningMachinesResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeWarningMachinesResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWarningMachinesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWarningMachinesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWarningMachinesResponseBody
	GetTotalCount() *int32
	SetWarningMachines(v []*DescribeWarningMachinesResponseBodyWarningMachines) *DescribeWarningMachinesResponseBody
	GetWarningMachines() []*DescribeWarningMachinesResponseBodyWarningMachines
}

type DescribeWarningMachinesResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6E4CBE87-E887-5AF1-86D2-78EB4AA3E771
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the servers.
	WarningMachines []*DescribeWarningMachinesResponseBodyWarningMachines `json:"WarningMachines,omitempty" xml:"WarningMachines,omitempty" type:"Repeated"`
}

func (s DescribeWarningMachinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarningMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWarningMachinesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWarningMachinesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWarningMachinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWarningMachinesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWarningMachinesResponseBody) GetWarningMachines() []*DescribeWarningMachinesResponseBodyWarningMachines {
	return s.WarningMachines
}

func (s *DescribeWarningMachinesResponseBody) SetCount(v int32) *DescribeWarningMachinesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetCurrentPage(v int32) *DescribeWarningMachinesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetPageSize(v int32) *DescribeWarningMachinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetRequestId(v string) *DescribeWarningMachinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetTotalCount(v int32) *DescribeWarningMachinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetWarningMachines(v []*DescribeWarningMachinesResponseBodyWarningMachines) *DescribeWarningMachinesResponseBody {
	s.WarningMachines = v
	return s
}

func (s *DescribeWarningMachinesResponseBody) Validate() error {
	if s.WarningMachines != nil {
		for _, item := range s.WarningMachines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWarningMachinesResponseBodyWarningMachines struct {
	// The type of cloud product assets.
	//
	// > You can call [ListCloudAssetInstances](~~ListCloudAssetInstances~~) to query the types of cloud product assets.
	//
	// example:
	//
	// 0
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The edition of Security Center that is authorized to protect the asset. Valid values:
	//
	// 	- **1**: Basic edition
	//
	// 	- **6**: Anti-virus edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **3**: Enterprise edition
	//
	// 	- **7**: Ultimate edition
	//
	// 	- **10**: Value-added Plan edition
	//
	// example:
	//
	// 7
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Indicates whether Security Center is authorized to protect the asset. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// The container ID.
	//
	// example:
	//
	// 48a6d9a92435a13ad573372c3f3c63b7e04d106458141df9f92155709d******
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// redis
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The number of **high-risk*	- items on the server.
	//
	// example:
	//
	// 10
	HighWarningCount *int32 `json:"HighWarningCount,omitempty" xml:"HighWarningCount,omitempty"`
	// The server ID.
	//
	// example:
	//
	// i-m5ee8b9nr3h30mhb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// oracle-win-00****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 1.2.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The number of **low-risk*	- items on the server.
	//
	// example:
	//
	// 3
	LowWarningCount *int32 `json:"LowWarningCount,omitempty" xml:"LowWarningCount,omitempty"`
	// The number of **medium-risk*	- items on the server.
	//
	// example:
	//
	// 2
	MediumWarningCount *int32 `json:"MediumWarningCount,omitempty" xml:"MediumWarningCount,omitempty"`
	// Indicates whether the agent is online. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// The number of the check items that passed the baseline check on the server.
	//
	// example:
	//
	// 10
	PassCount *int32 `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	// Indicates whether a port on the server is accessible over the Internet. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PortOpen *bool `json:"PortOpen,omitempty" xml:"PortOpen,omitempty"`
	// The ID of the region in which the server is deployed.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The verification status of the risk item after the risk item is fixed. Valid values:
	//
	// 	- **1**: complete
	//
	// 	- **2**: verifying
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server on which the baseline check is performed.
	//
	// example:
	//
	// 9888955c-0076-49da-bd9c-34f5492b****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWarningMachinesResponseBodyWarningMachines) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarningMachinesResponseBodyWarningMachines) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetBind() *bool {
	return s.Bind
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetHighWarningCount() *int32 {
	return s.HighWarningCount
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetLowWarningCount() *int32 {
	return s.LowWarningCount
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetMediumWarningCount() *int32 {
	return s.MediumWarningCount
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetOnline() *bool {
	return s.Online
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetPassCount() *int32 {
	return s.PassCount
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetPortOpen() *bool {
	return s.PortOpen
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetAssetType(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.AssetType = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetAuthVersion(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.AuthVersion = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetBind(v bool) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.Bind = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetContainerId(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.ContainerId = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetContainerName(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.ContainerName = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetHighWarningCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.HighWarningCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetInstanceId(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.InstanceId = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetInstanceName(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.InstanceName = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetInternetIp(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.InternetIp = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetIntranetIp(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.IntranetIp = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetLowWarningCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.LowWarningCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetMediumWarningCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.MediumWarningCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetOnline(v bool) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.Online = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetPassCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.PassCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetPortOpen(v bool) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.PortOpen = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetRegionId(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.RegionId = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetStatus(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.Status = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetUuid(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.Uuid = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) Validate() error {
	return dara.Validate(s)
}
