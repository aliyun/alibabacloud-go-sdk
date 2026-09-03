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
	SetHdfsAuthConfig(v *HdfsAuthConfig) *AddressDetail
	GetHdfsAuthConfig() *HdfsAuthConfig
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
	// The AccessKey ID used to access the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_access_id
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The AccessKey Secret used to access the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_secret_key
	AccessSecret *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	// The address type. Valid values:
	//
	// Object storage: `oss`, `s3`, `qiniu`, `cos`, `bos`, `obs`, `ks3`, `ucloud`, `upyun`, `azure`, `gcp`, and `tos`.
	//
	// File storage: `local`, `obsfs`, and `ftp`.
	//
	// Inventory: `inv`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ossinv
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// A list of agents required to access data over a leased line or a VPN connection.
	//
	// example:
	//
	// agent1,agent2
	AgentList *string `json:"AgentList,omitempty" xml:"AgentList,omitempty"`
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Specifies the underlying storage type, such as `oss`, `s3`, or `obs`, when `AddressType` is `inv`.
	//
	// example:
	//
	// oss
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The domain of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_domain
	Domain         *string         `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HdfsAuthConfig *HdfsAuthConfig `json:"HdfsAuthConfig,omitempty" xml:"HdfsAuthConfig,omitempty"`
	// The AccessKey ID used to access the inventory bucket.
	//
	// example:
	//
	// test_inv_access_id
	InvAccessId *string `json:"InvAccessId,omitempty" xml:"InvAccessId,omitempty"`
	// The AccessKey Secret used to access the inventory bucket.
	//
	// example:
	//
	// test_inv_secret_key
	InvAccessSecret *string `json:"InvAccessSecret,omitempty" xml:"InvAccessSecret,omitempty"`
	// The name of the inventory bucket.
	//
	// example:
	//
	// test_inv_bucket
	InvBucket *string `json:"InvBucket,omitempty" xml:"InvBucket,omitempty"`
	// The domain of the inventory bucket.
	//
	// example:
	//
	// test_inv_domain
	InvDomain *string `json:"InvDomain,omitempty" xml:"InvDomain,omitempty"`
	// The type of the inventory bucket.<br>
	//
	// Valid values: `oss` and `s3`.<br>
	//
	// example:
	//
	// oss
	InvLocation *string `json:"InvLocation,omitempty" xml:"InvLocation,omitempty"`
	// The path to the manifest file, which lists the objects and their format.
	//
	// example:
	//
	// manifest.json
	InvPath *string `json:"InvPath,omitempty" xml:"InvPath,omitempty"`
	// The region of the inventory bucket. This parameter is required for OSS inventory buckets and eliminates the need to specify a domain. Use the OSS-specific region ID format.
	//
	// example:
	//
	// test_inv_region_id
	InvRegionId *string `json:"InvRegionId,omitempty" xml:"InvRegionId,omitempty"`
	// The authorization role for the inventory bucket. If the inventory bucket is an OSS bucket, you can use this role to grant access instead of an AccessKey.
	//
	// example:
	//
	// test_inv_role
	InvRole *string `json:"InvRole,omitempty" xml:"InvRole,omitempty"`
	// The object prefix in the bucket.
	//
	// example:
	//
	// test_prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The region of the bucket. This parameter is required for OSS buckets and eliminates the need to specify a domain. Use the OSS-specific region ID format.
	//
	// example:
	//
	// test_region_id
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The authorization role. If the bucket is an OSS bucket, you can use this role to grant access instead of an AccessKey.
	//
	// example:
	//
	// test_role
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

func (s *AddressDetail) GetHdfsAuthConfig() *HdfsAuthConfig {
	return s.HdfsAuthConfig
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

func (s *AddressDetail) SetHdfsAuthConfig(v *HdfsAuthConfig) *AddressDetail {
	s.HdfsAuthConfig = v
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
	if s.HdfsAuthConfig != nil {
		if err := s.HdfsAuthConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
