// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddon(v *DescribeAddonTemplateResponseBodyAddon) *DescribeAddonTemplateResponseBody
	GetAddon() *DescribeAddonTemplateResponseBodyAddon
	SetPageNumber(v int64) *DescribeAddonTemplateResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAddonTemplateResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeAddonTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAddonTemplateResponseBody
	GetTotalCount() *int32
}

type DescribeAddonTemplateResponseBody struct {
	// The details of the addon template.
	Addon *DescribeAddonTemplateResponseBodyAddon `json:"Addon,omitempty" xml:"Addon,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAddonTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBody) GetAddon() *DescribeAddonTemplateResponseBodyAddon {
	return s.Addon
}

func (s *DescribeAddonTemplateResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAddonTemplateResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAddonTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAddonTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAddonTemplateResponseBody) SetAddon(v *DescribeAddonTemplateResponseBodyAddon) *DescribeAddonTemplateResponseBody {
	s.Addon = v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetPageNumber(v int64) *DescribeAddonTemplateResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetPageSize(v int64) *DescribeAddonTemplateResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetRequestId(v string) *DescribeAddonTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddonTemplateResponseBody) SetTotalCount(v int32) *DescribeAddonTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAddonTemplateResponseBody) Validate() error {
	if s.Addon != nil {
		if err := s.Addon.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAddonTemplateResponseBodyAddon struct {
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon icon.
	//
	// example:
	//
	// /assets/icons/your_icon.svg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The date when the addon template was last updated.
	//
	// example:
	//
	// 2024-08-22 18:11:17
	LastUpdate *string `json:"LastUpdate,omitempty" xml:"LastUpdate,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	ResourcesSpec *DescribeAddonTemplateResponseBodyAddonResourcesSpec `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty" type:"Struct"`
	// The addon configurations.
	ServicesSpec []*DescribeAddonTemplateResponseBodyAddonServicesSpec `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty" type:"Repeated"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddon) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddon) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetDescription() *string {
	return s.Description
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetIcon() *string {
	return s.Icon
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetLabel() *string {
	return s.Label
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetLastUpdate() *string {
	return s.LastUpdate
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetName() *string {
	return s.Name
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetResourcesSpec() *DescribeAddonTemplateResponseBodyAddonResourcesSpec {
	return s.ResourcesSpec
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetServicesSpec() []*DescribeAddonTemplateResponseBodyAddonServicesSpec {
	return s.ServicesSpec
}

func (s *DescribeAddonTemplateResponseBodyAddon) GetVersion() *string {
	return s.Version
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetDescription(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Description = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetIcon(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Icon = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetLabel(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Label = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetLastUpdate(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.LastUpdate = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetName(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Name = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetResourcesSpec(v *DescribeAddonTemplateResponseBodyAddonResourcesSpec) *DescribeAddonTemplateResponseBodyAddon {
	s.ResourcesSpec = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetServicesSpec(v []*DescribeAddonTemplateResponseBodyAddonServicesSpec) *DescribeAddonTemplateResponseBodyAddon {
	s.ServicesSpec = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) SetVersion(v string) *DescribeAddonTemplateResponseBodyAddon {
	s.Version = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddon) Validate() error {
	if s.ResourcesSpec != nil {
		if err := s.ResourcesSpec.Validate(); err != nil {
			return err
		}
	}
	if s.ServicesSpec != nil {
		for _, item := range s.ServicesSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAddonTemplateResponseBodyAddonResourcesSpec struct {
	// The Elastic Compute Service (ECS) resource configurations of the addon.
	EcsResources []*AddonNodeTemplate `json:"EcsResources,omitempty" xml:"EcsResources,omitempty" type:"Repeated"`
	// The Elastic IP Address (EIP) configurations of the service.
	EipResource *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource `json:"EipResource,omitempty" xml:"EipResource,omitempty" type:"Struct"`
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpec) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpec) GetEcsResources() []*AddonNodeTemplate {
	return s.EcsResources
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpec) GetEipResource() *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	return s.EipResource
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpec) SetEcsResources(v []*AddonNodeTemplate) *DescribeAddonTemplateResponseBodyAddonResourcesSpec {
	s.EcsResources = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpec) SetEipResource(v *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) *DescribeAddonTemplateResponseBodyAddonResourcesSpec {
	s.EipResource = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpec) Validate() error {
	if s.EcsResources != nil {
		for _, item := range s.EcsResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EipResource != nil {
		if err := s.EipResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource struct {
	// Indicates whether the EIP is automatically created.
	//
	// example:
	//
	// True
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// The maximum bandwidth of the EIP. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The EIP ID.
	//
	// example:
	//
	// eip-bp1jwtsuoiif2qf90****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	// The billing method of the EIP. Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// Default value: PostPaid
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the EIP. Valid values:
	//
	// 	- PayByBandwidth: pay by bandwidth.
	//
	// 	- PayByTraffic: pay by data transfer.
	//
	// Valid values of N: 1 to 10.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) GetAutoCreate() *bool {
	return s.AutoCreate
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) GetEipInstanceId() *string {
	return s.EipInstanceId
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetAutoCreate(v bool) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.AutoCreate = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetBandwidth(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.Bandwidth = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetEipInstanceId(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.EipInstanceId = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetInstanceChargeType(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) SetInternetChargeType(v string) *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonResourcesSpecEipResource) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonTemplateResponseBodyAddonServicesSpec struct {
	// The parameter configurations of the service.
	InputParams []*DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams `json:"InputParams,omitempty" xml:"InputParams,omitempty" type:"Repeated"`
	// The security group configurations of the service.
	NetworkACL []*DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL `json:"NetworkACL,omitempty" xml:"NetworkACL,omitempty" type:"Repeated"`
	// The service access type.
	//
	// example:
	//
	// URL
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The service access URL.
	//
	// example:
	//
	// https://47.96.xx.xx:12008
	ServiceAccessUrl *string `json:"ServiceAccessUrl,omitempty" xml:"ServiceAccessUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Web Portal
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpec) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) GetInputParams() []*DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	return s.InputParams
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) GetNetworkACL() []*DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL {
	return s.NetworkACL
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) GetServiceAccessType() *string {
	return s.ServiceAccessType
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) GetServiceAccessUrl() *string {
	return s.ServiceAccessUrl
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetInputParams(v []*DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.InputParams = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetNetworkACL(v []*DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.NetworkACL = v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetServiceAccessType(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.ServiceAccessType = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetServiceAccessUrl(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.ServiceAccessUrl = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) SetServiceName(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpec {
	s.ServiceName = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpec) Validate() error {
	if s.InputParams != nil {
		for _, item := range s.InputParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkACL != nil {
		for _, item := range s.NetworkACL {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams struct {
	// The help information of the parameter.
	HelpText *string `json:"HelpText,omitempty" xml:"HelpText,omitempty"`
	// The parameter label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The parameter name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL_HOME
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter type.
	//
	// This parameter is required.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The parameter value.
	//
	// This parameter is required.
	//
	// example:
	//
	// usr/local/mysql
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) GetHelpText() *string {
	return s.HelpText
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) GetLabel() *string {
	return s.Label
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) GetName() *string {
	return s.Name
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) GetType() *string {
	return s.Type
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) GetValue() *string {
	return s.Value
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetHelpText(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.HelpText = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetLabel(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Label = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetName(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Name = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetType(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Type = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) SetValue(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams {
	s.Value = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecInputParams) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL struct {
	// The protocol type. Valid values:
	//
	// 	- **TCP**: forwards TCP packets.
	//
	// 	- **UDP**: forwards UDP packets.
	//
	// 	- **Any**: forwards all packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The port number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *float32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The source CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/12
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) GoString() string {
	return s.String()
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) GetPort() *float32 {
	return s.Port
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) SetIpProtocol(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL {
	s.IpProtocol = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) SetPort(v float32) *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL {
	s.Port = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) SetSourceCidrIp(v string) *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeAddonTemplateResponseBodyAddonServicesSpecNetworkACL) Validate() error {
	return dara.Validate(s)
}
