// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSmartAccessGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSmartAccessGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSmartAccessGatewaysResponseBody
	GetRequestId() *string
	SetSmartAccessGateways(v *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways) *DescribeSmartAccessGatewaysResponseBody
	GetSmartAccessGateways() *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways
	SetTotalCount(v int32) *DescribeSmartAccessGatewaysResponseBody
	GetTotalCount() *int32
}

type DescribeSmartAccessGatewaysResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// ADE68CEE-8E4F-4D0B-9EE9-2C2FFAABF41F
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmartAccessGateways *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways `json:"SmartAccessGateways,omitempty" xml:"SmartAccessGateways,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSmartAccessGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSmartAccessGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSmartAccessGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSmartAccessGatewaysResponseBody) GetSmartAccessGateways() *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways {
	return s.SmartAccessGateways
}

func (s *DescribeSmartAccessGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSmartAccessGatewaysResponseBody) SetPageNumber(v int32) *DescribeSmartAccessGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBody) SetPageSize(v int32) *DescribeSmartAccessGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBody) SetRequestId(v string) *DescribeSmartAccessGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBody) SetSmartAccessGateways(v *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways) *DescribeSmartAccessGatewaysResponseBody {
	s.SmartAccessGateways = v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBody) SetTotalCount(v int32) *DescribeSmartAccessGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBody) Validate() error {
	if s.SmartAccessGateways != nil {
		if err := s.SmartAccessGateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSmartAccessGatewaysResponseBodySmartAccessGateways struct {
	SmartAccessGateway []*DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway `json:"SmartAccessGateway,omitempty" xml:"SmartAccessGateway,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGateways) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways) GetSmartAccessGateway() []*DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	return s.SmartAccessGateway
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways) SetSmartAccessGateway(v []*DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways {
	s.SmartAccessGateway = v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGateways) Validate() error {
	if s.SmartAccessGateway != nil {
		for _, item := range s.SmartAccessGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway struct {
	AccelerateBandwidth           *int64                                                                             `json:"AccelerateBandwidth,omitempty" xml:"AccelerateBandwidth,omitempty"`
	AccessPointId                 *string                                                                            `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	AclIds                        *string                                                                            `json:"AclIds,omitempty" xml:"AclIds,omitempty"`
	ApplicationBandwidthPackageId *string                                                                            `json:"ApplicationBandwidthPackageId,omitempty" xml:"ApplicationBandwidthPackageId,omitempty"`
	AssociatedCcnId               *string                                                                            `json:"AssociatedCcnId,omitempty" xml:"AssociatedCcnId,omitempty"`
	AssociatedCcnName             *string                                                                            `json:"AssociatedCcnName,omitempty" xml:"AssociatedCcnName,omitempty"`
	BackupSoftwareVersion         *string                                                                            `json:"BackupSoftwareVersion,omitempty" xml:"BackupSoftwareVersion,omitempty"`
	BackupStatus                  *string                                                                            `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	CidrBlock                     *string                                                                            `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	City                          *string                                                                            `json:"City,omitempty" xml:"City,omitempty"`
	CreateTime                    *int64                                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataPlan                      *int64                                                                             `json:"DataPlan,omitempty" xml:"DataPlan,omitempty"`
	Description                   *string                                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	DpiMonitorStatus              *string                                                                            `json:"DpiMonitorStatus,omitempty" xml:"DpiMonitorStatus,omitempty"`
	DpiStatus                     *string                                                                            `json:"DpiStatus,omitempty" xml:"DpiStatus,omitempty"`
	EnableAdvancedMonitor         *bool                                                                              `json:"EnableAdvancedMonitor,omitempty" xml:"EnableAdvancedMonitor,omitempty"`
	EnableSoftwareConnectionAudit *bool                                                                              `json:"EnableSoftwareConnectionAudit,omitempty" xml:"EnableSoftwareConnectionAudit,omitempty"`
	EndTime                       *int64                                                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnterpriseCode                *string                                                                            `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	HardwareVersion               *string                                                                            `json:"HardwareVersion,omitempty" xml:"HardwareVersion,omitempty"`
	IRIds                         *string                                                                            `json:"IRIds,omitempty" xml:"IRIds,omitempty"`
	IdaasApplicationId            *string                                                                            `json:"IdaasApplicationId,omitempty" xml:"IdaasApplicationId,omitempty"`
	IdaasId                       *string                                                                            `json:"IdaasId,omitempty" xml:"IdaasId,omitempty"`
	IpsecStatus                   *string                                                                            `json:"IpsecStatus,omitempty" xml:"IpsecStatus,omitempty"`
	Isp                           *string                                                                            `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Links                         *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks `json:"Links,omitempty" xml:"Links,omitempty" type:"Struct"`
	MaxBandwidth                  *string                                                                            `json:"MaxBandwidth,omitempty" xml:"MaxBandwidth,omitempty"`
	Name                          *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Position                      *string                                                                            `json:"Position,omitempty" xml:"Position,omitempty"`
	QosIds                        *string                                                                            `json:"QosIds,omitempty" xml:"QosIds,omitempty"`
	ResellerInstanceId            *string                                                                            `json:"ResellerInstanceId,omitempty" xml:"ResellerInstanceId,omitempty"`
	ResellerUid                   *string                                                                            `json:"ResellerUid,omitempty" xml:"ResellerUid,omitempty"`
	ResourceGroupId               *string                                                                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RoutingStrategy               *string                                                                            `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	SecurityLockThreshold         *int32                                                                             `json:"SecurityLockThreshold,omitempty" xml:"SecurityLockThreshold,omitempty"`
	SerialNumber                  *string                                                                            `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SmartAGId                     *string                                                                            `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	SmartAGUid                    *int64                                                                             `json:"SmartAGUid,omitempty" xml:"SmartAGUid,omitempty"`
	SoftwareVersion               *string                                                                            `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
	Status                        *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	UpBandwidth4G                 *int32                                                                             `json:"UpBandwidth4G,omitempty" xml:"UpBandwidth4G,omitempty"`
	UpBandwidthWan                *int32                                                                             `json:"UpBandwidthWan,omitempty" xml:"UpBandwidthWan,omitempty"`
	UserCount                     *int32                                                                             `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	VpnStatus                     *string                                                                            `json:"VpnStatus,omitempty" xml:"VpnStatus,omitempty"`
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetAccelerateBandwidth() *int64 {
	return s.AccelerateBandwidth
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetAclIds() *string {
	return s.AclIds
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetApplicationBandwidthPackageId() *string {
	return s.ApplicationBandwidthPackageId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetAssociatedCcnId() *string {
	return s.AssociatedCcnId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetAssociatedCcnName() *string {
	return s.AssociatedCcnName
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetBackupSoftwareVersion() *string {
	return s.BackupSoftwareVersion
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetCity() *string {
	return s.City
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetDataPlan() *int64 {
	return s.DataPlan
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetDescription() *string {
	return s.Description
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetDpiMonitorStatus() *string {
	return s.DpiMonitorStatus
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetDpiStatus() *string {
	return s.DpiStatus
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetEnableAdvancedMonitor() *bool {
	return s.EnableAdvancedMonitor
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetEnableSoftwareConnectionAudit() *bool {
	return s.EnableSoftwareConnectionAudit
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetEnterpriseCode() *string {
	return s.EnterpriseCode
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetHardwareVersion() *string {
	return s.HardwareVersion
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetIRIds() *string {
	return s.IRIds
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetIdaasApplicationId() *string {
	return s.IdaasApplicationId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetIdaasId() *string {
	return s.IdaasId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetIpsecStatus() *string {
	return s.IpsecStatus
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetLinks() *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks {
	return s.Links
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetMaxBandwidth() *string {
	return s.MaxBandwidth
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetName() *string {
	return s.Name
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetPosition() *string {
	return s.Position
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetQosIds() *string {
	return s.QosIds
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetResellerInstanceId() *string {
	return s.ResellerInstanceId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetResellerUid() *string {
	return s.ResellerUid
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetSecurityLockThreshold() *int32 {
	return s.SecurityLockThreshold
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetSmartAGUid() *int64 {
	return s.SmartAGUid
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetSoftwareVersion() *string {
	return s.SoftwareVersion
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetStatus() *string {
	return s.Status
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetUpBandwidth4G() *int32 {
	return s.UpBandwidth4G
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetUpBandwidthWan() *int32 {
	return s.UpBandwidthWan
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) GetVpnStatus() *string {
	return s.VpnStatus
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetAccelerateBandwidth(v int64) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.AccelerateBandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetAccessPointId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.AccessPointId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetAclIds(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.AclIds = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetApplicationBandwidthPackageId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.ApplicationBandwidthPackageId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetAssociatedCcnId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.AssociatedCcnId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetAssociatedCcnName(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.AssociatedCcnName = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetBackupSoftwareVersion(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.BackupSoftwareVersion = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetBackupStatus(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.BackupStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetCidrBlock(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.CidrBlock = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetCity(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.City = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetCreateTime(v int64) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.CreateTime = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetDataPlan(v int64) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.DataPlan = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetDescription(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.Description = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetDpiMonitorStatus(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.DpiMonitorStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetDpiStatus(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.DpiStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetEnableAdvancedMonitor(v bool) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.EnableAdvancedMonitor = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetEnableSoftwareConnectionAudit(v bool) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.EnableSoftwareConnectionAudit = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetEndTime(v int64) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.EndTime = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetEnterpriseCode(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.EnterpriseCode = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetHardwareVersion(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.HardwareVersion = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetIRIds(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.IRIds = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetIdaasApplicationId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.IdaasApplicationId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetIdaasId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.IdaasId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetIpsecStatus(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.IpsecStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetIsp(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.Isp = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetLinks(v *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.Links = v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetMaxBandwidth(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.MaxBandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetName(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.Name = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetPosition(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.Position = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetQosIds(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.QosIds = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetResellerInstanceId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.ResellerInstanceId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetResellerUid(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.ResellerUid = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetResourceGroupId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetRoutingStrategy(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.RoutingStrategy = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetSecurityLockThreshold(v int32) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.SecurityLockThreshold = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetSerialNumber(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.SerialNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetSmartAGId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetSmartAGUid(v int64) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.SmartAGUid = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetSoftwareVersion(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.SoftwareVersion = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetStatus(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.Status = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetUpBandwidth4G(v int32) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.UpBandwidth4G = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetUpBandwidthWan(v int32) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.UpBandwidthWan = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetUserCount(v int32) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.UserCount = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) SetVpnStatus(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway {
	s.VpnStatus = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGateway) Validate() error {
	if s.Links != nil {
		if err := s.Links.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks struct {
	Link []*DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink `json:"Link,omitempty" xml:"Link,omitempty" type:"Repeated"`
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks) GetLink() []*DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	return s.Link
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks) SetLink(v []*DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks {
	s.Link = v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinks) Validate() error {
	if s.Link != nil {
		for _, item := range s.Link {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink struct {
	Bandwidth              *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CommodityType          *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	EndTime                *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RelateInstanceId       *string `json:"RelateInstanceId,omitempty" xml:"RelateInstanceId,omitempty"`
	RelateInstanceRegionId *string `json:"RelateInstanceRegionId,omitempty" xml:"RelateInstanceRegionId,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetCommodityType() *string {
	return s.CommodityType
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetRelateInstanceId() *string {
	return s.RelateInstanceId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetRelateInstanceRegionId() *string {
	return s.RelateInstanceRegionId
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetStatus() *string {
	return s.Status
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) GetType() *string {
	return s.Type
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetBandwidth(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.Bandwidth = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetCommodityType(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.CommodityType = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetEndTime(v int64) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.EndTime = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetInstanceId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.InstanceId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetRelateInstanceId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.RelateInstanceId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetRelateInstanceRegionId(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.RelateInstanceRegionId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetStatus(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.Status = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) SetType(v string) *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink {
	s.Type = &v
	return s
}

func (s *DescribeSmartAccessGatewaysResponseBodySmartAccessGatewaysSmartAccessGatewayLinksLink) Validate() error {
	return dara.Validate(s)
}
