// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *AddressDetail
	GetAccessId() *string
	SetAccessSecret(v string) *AddressDetail
	GetAccessSecret() *string
	SetAddressType(v string) *AddressDetail
	GetAddressType() *string
	SetAgentList(v string) *AddressDetail
	GetAgentList() *string
	SetBucket(v string) *AddressDetail
	GetBucket() *string
	SetDataType(v string) *AddressDetail
	GetDataType() *string
	SetDomain(v string) *AddressDetail
	GetDomain() *string
	SetInvAccessId(v string) *AddressDetail
	GetInvAccessId() *string
	SetInvAccessSecret(v string) *AddressDetail
	GetInvAccessSecret() *string
	SetInvBucket(v string) *AddressDetail
	GetInvBucket() *string
	SetInvDomain(v string) *AddressDetail
	GetInvDomain() *string
	SetInvLocation(v string) *AddressDetail
	GetInvLocation() *string
	SetInvPath(v string) *AddressDetail
	GetInvPath() *string
	SetInvRegionId(v string) *AddressDetail
	GetInvRegionId() *string
	SetInvRole(v string) *AddressDetail
	GetInvRole() *string
	SetPrefix(v string) *AddressDetail
	GetPrefix() *string
	SetRegionId(v string) *AddressDetail
	GetRegionId() *string
	SetRole(v string) *AddressDetail
	GetRole() *string
}

type AddressDetail struct {
	// This parameter is required.
	//
	// example:
	//
	// *******************
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *******************
	AccessSecret *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ossinv
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// example:
	//
	// agnet1,agent2,agent3
	AgentList *string `json:"AgentList,omitempty" xml:"AgentList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <your-bucket-name>
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// oss
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// *******************
	InvAccessId *string `json:"InvAccessId,omitempty" xml:"InvAccessId,omitempty"`
	// example:
	//
	// *******************
	InvAccessSecret *string `json:"InvAccessSecret,omitempty" xml:"InvAccessSecret,omitempty"`
	// example:
	//
	// <your-inv-bucket-name>
	InvBucket *string `json:"InvBucket,omitempty" xml:"InvBucket,omitempty"`
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	InvDomain *string `json:"InvDomain,omitempty" xml:"InvDomain,omitempty"`
	// example:
	//
	// oss
	InvLocation *string `json:"InvLocation,omitempty" xml:"InvLocation,omitempty"`
	// example:
	//
	// dir/manifest.json
	InvPath *string `json:"InvPath,omitempty" xml:"InvPath,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	InvRegionId *string `json:"InvRegionId,omitempty" xml:"InvRegionId,omitempty"`
	// example:
	//
	// <your-role-name>
	InvRole *string `json:"InvRole,omitempty" xml:"InvRole,omitempty"`
	// example:
	//
	// dir1/dir2/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// oss-cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// <your-role-name>
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AddressDetail) String() string {
	return dara.Prettify(s)
}

func (s AddressDetail) GoString() string {
	return s.String()
}

func (s *AddressDetail) GetAccessId() *string {
	return s.AccessId
}

func (s *AddressDetail) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *AddressDetail) GetAddressType() *string {
	return s.AddressType
}

func (s *AddressDetail) GetAgentList() *string {
	return s.AgentList
}

func (s *AddressDetail) GetBucket() *string {
	return s.Bucket
}

func (s *AddressDetail) GetDataType() *string {
	return s.DataType
}

func (s *AddressDetail) GetDomain() *string {
	return s.Domain
}

func (s *AddressDetail) GetInvAccessId() *string {
	return s.InvAccessId
}

func (s *AddressDetail) GetInvAccessSecret() *string {
	return s.InvAccessSecret
}

func (s *AddressDetail) GetInvBucket() *string {
	return s.InvBucket
}

func (s *AddressDetail) GetInvDomain() *string {
	return s.InvDomain
}

func (s *AddressDetail) GetInvLocation() *string {
	return s.InvLocation
}

func (s *AddressDetail) GetInvPath() *string {
	return s.InvPath
}

func (s *AddressDetail) GetInvRegionId() *string {
	return s.InvRegionId
}

func (s *AddressDetail) GetInvRole() *string {
	return s.InvRole
}

func (s *AddressDetail) GetPrefix() *string {
	return s.Prefix
}

func (s *AddressDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *AddressDetail) GetRole() *string {
	return s.Role
}

func (s *AddressDetail) SetAccessId(v string) *AddressDetail {
	s.AccessId = &v
	return s
}

func (s *AddressDetail) SetAccessSecret(v string) *AddressDetail {
	s.AccessSecret = &v
	return s
}

func (s *AddressDetail) SetAddressType(v string) *AddressDetail {
	s.AddressType = &v
	return s
}

func (s *AddressDetail) SetAgentList(v string) *AddressDetail {
	s.AgentList = &v
	return s
}

func (s *AddressDetail) SetBucket(v string) *AddressDetail {
	s.Bucket = &v
	return s
}

func (s *AddressDetail) SetDataType(v string) *AddressDetail {
	s.DataType = &v
	return s
}

func (s *AddressDetail) SetDomain(v string) *AddressDetail {
	s.Domain = &v
	return s
}

func (s *AddressDetail) SetInvAccessId(v string) *AddressDetail {
	s.InvAccessId = &v
	return s
}

func (s *AddressDetail) SetInvAccessSecret(v string) *AddressDetail {
	s.InvAccessSecret = &v
	return s
}

func (s *AddressDetail) SetInvBucket(v string) *AddressDetail {
	s.InvBucket = &v
	return s
}

func (s *AddressDetail) SetInvDomain(v string) *AddressDetail {
	s.InvDomain = &v
	return s
}

func (s *AddressDetail) SetInvLocation(v string) *AddressDetail {
	s.InvLocation = &v
	return s
}

func (s *AddressDetail) SetInvPath(v string) *AddressDetail {
	s.InvPath = &v
	return s
}

func (s *AddressDetail) SetInvRegionId(v string) *AddressDetail {
	s.InvRegionId = &v
	return s
}

func (s *AddressDetail) SetInvRole(v string) *AddressDetail {
	s.InvRole = &v
	return s
}

func (s *AddressDetail) SetPrefix(v string) *AddressDetail {
	s.Prefix = &v
	return s
}

func (s *AddressDetail) SetRegionId(v string) *AddressDetail {
	s.RegionId = &v
	return s
}

func (s *AddressDetail) SetRole(v string) *AddressDetail {
	s.Role = &v
	return s
}

func (s *AddressDetail) Validate() error {
	return dara.Validate(s)
}
