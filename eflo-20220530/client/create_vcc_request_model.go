// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessCouldService(v bool) *CreateVccRequest
	GetAccessCouldService() *bool
	SetBandwidth(v int32) *CreateVccRequest
	GetBandwidth() *int32
	SetBgpAsn(v int64) *CreateVccRequest
	GetBgpAsn() *int64
	SetBgpCidr(v string) *CreateVccRequest
	GetBgpCidr() *string
	SetCenId(v string) *CreateVccRequest
	GetCenId() *string
	SetCenOwnerId(v string) *CreateVccRequest
	GetCenOwnerId() *string
	SetConnectionType(v string) *CreateVccRequest
	GetConnectionType() *string
	SetDescription(v string) *CreateVccRequest
	GetDescription() *string
	SetRegionId(v string) *CreateVccRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVccRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateVccRequestTag) *CreateVccRequest
	GetTag() []*CreateVccRequestTag
	SetVSwitchId(v string) *CreateVccRequest
	GetVSwitchId() *string
	SetVccId(v string) *CreateVccRequest
	GetVccId() *string
	SetVccName(v string) *CreateVccRequest
	GetVccName() *string
	SetVpcId(v string) *CreateVccRequest
	GetVpcId() *string
	SetVpdId(v string) *CreateVccRequest
	GetVpdId() *string
	SetZoneId(v string) *CreateVccRequest
	GetZoneId() *string
}

type CreateVccRequest struct {
	// Enabled access to cloud services. Optional values:
	//
	// 	- **true**: Enable access to cloud services
	//
	// 	- **false**: Do not enable access to cloud services
	//
	// example:
	//
	// true
	AccessCouldService *bool `json:"AccessCouldService,omitempty" xml:"AccessCouldService,omitempty"`
	// The bandwidth. Unit: Mbit /s. The minimum value is 1000, representing 1Gbps bandwidth; the maximum value is 400000, representing 400Gbps bandwidth.
	//
	// >  1Gbps = 1000Mbps
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// bgp as number
	//
	// example:
	//
	// bgpAsn
	BgpAsn *int64 `json:"BgpAsn,omitempty" xml:"BgpAsn,omitempty"`
	// Internet segment, on-premises input, off-premises default
	//
	// example:
	//
	// 10.0.0.0/24
	BgpCidr *string `json:"BgpCidr,omitempty" xml:"BgpCidr,omitempty"`
	// CEN Instance ID
	//
	// example:
	//
	// cen-bkiw0x1347roekr7f2
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Account to which cen belongs
	//
	// example:
	//
	// 1511928242963727
	CenOwnerId *string `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The connection mode. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **CEN (CENTR)**
	//
	// example:
	//
	// CENTR
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The description of the document.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.htm?spm=a2c4g.11186623.0.0.29e15d7akXhpuu).
	//
	// example:
	//
	// rg-aeky5f3qx6ceapq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags.
	Tag []*CreateVccRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch. [Virtual Private Cloud VSwitch](https://help.aliyun.com/document_detail/100380.html).
	//
	// You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query created vSwitches.
	//
	// example:
	//
	// vsw-t4nahb0pxckgktx1kot8q
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// Lingjun Connection Name
	//
	// example:
	//
	// test
	VccName *string `json:"VccName,omitempty" xml:"VccName,omitempty"`
	// Virtual Private Cloud IDs; [What is Virtual Private Cloud](https://help.aliyun.com/document_detail/34217.html)
	//
	// You can call the [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html#demo-0) operation to query the specified VPC.
	//
	// example:
	//
	// vpc-uf6aa4ddo97frj22tgp52
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Lingjun CIDR block instance ID
	//
	// example:
	//
	// vpd-t2jseldp
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID of the disk.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVccRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVccRequest) GoString() string {
	return s.String()
}

func (s *CreateVccRequest) GetAccessCouldService() *bool {
	return s.AccessCouldService
}

func (s *CreateVccRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateVccRequest) GetBgpAsn() *int64 {
	return s.BgpAsn
}

func (s *CreateVccRequest) GetBgpCidr() *string {
	return s.BgpCidr
}

func (s *CreateVccRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateVccRequest) GetCenOwnerId() *string {
	return s.CenOwnerId
}

func (s *CreateVccRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *CreateVccRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVccRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVccRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVccRequest) GetTag() []*CreateVccRequestTag {
	return s.Tag
}

func (s *CreateVccRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateVccRequest) GetVccId() *string {
	return s.VccId
}

func (s *CreateVccRequest) GetVccName() *string {
	return s.VccName
}

func (s *CreateVccRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVccRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *CreateVccRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateVccRequest) SetAccessCouldService(v bool) *CreateVccRequest {
	s.AccessCouldService = &v
	return s
}

func (s *CreateVccRequest) SetBandwidth(v int32) *CreateVccRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateVccRequest) SetBgpAsn(v int64) *CreateVccRequest {
	s.BgpAsn = &v
	return s
}

func (s *CreateVccRequest) SetBgpCidr(v string) *CreateVccRequest {
	s.BgpCidr = &v
	return s
}

func (s *CreateVccRequest) SetCenId(v string) *CreateVccRequest {
	s.CenId = &v
	return s
}

func (s *CreateVccRequest) SetCenOwnerId(v string) *CreateVccRequest {
	s.CenOwnerId = &v
	return s
}

func (s *CreateVccRequest) SetConnectionType(v string) *CreateVccRequest {
	s.ConnectionType = &v
	return s
}

func (s *CreateVccRequest) SetDescription(v string) *CreateVccRequest {
	s.Description = &v
	return s
}

func (s *CreateVccRequest) SetRegionId(v string) *CreateVccRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVccRequest) SetResourceGroupId(v string) *CreateVccRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVccRequest) SetTag(v []*CreateVccRequestTag) *CreateVccRequest {
	s.Tag = v
	return s
}

func (s *CreateVccRequest) SetVSwitchId(v string) *CreateVccRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVccRequest) SetVccId(v string) *CreateVccRequest {
	s.VccId = &v
	return s
}

func (s *CreateVccRequest) SetVccName(v string) *CreateVccRequest {
	s.VccName = &v
	return s
}

func (s *CreateVccRequest) SetVpcId(v string) *CreateVccRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVccRequest) SetVpdId(v string) *CreateVccRequest {
	s.VpdId = &v
	return s
}

func (s *CreateVccRequest) SetZoneId(v string) *CreateVccRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateVccRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVccRequestTag struct {
	// The tag key of the VPN attachment.
	//
	// You cannot specify an empty string as a tag key. It can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tag-vcc
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the VPN connection.
	//
	// The tag value can be empty or a string of up to 128 characters. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
	//
	// example:
	//
	// vcc-group-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVccRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVccRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVccRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVccRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVccRequestTag) SetKey(v string) *CreateVccRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVccRequestTag) SetValue(v string) *CreateVccRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVccRequestTag) Validate() error {
	return dara.Validate(s)
}
