// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCACertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCACertificate(v string) *UploadCACertificateRequest
	GetCACertificate() *string
	SetCACertificateName(v string) *UploadCACertificateRequest
	GetCACertificateName() *string
	SetOwnerAccount(v string) *UploadCACertificateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UploadCACertificateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UploadCACertificateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UploadCACertificateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *UploadCACertificateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UploadCACertificateRequest
	GetResourceOwnerId() *int64
	SetTag(v []*UploadCACertificateRequestTag) *UploadCACertificateRequest
	GetTag() []*UploadCACertificateRequestTag
}

type UploadCACertificateRequest struct {
	// The information about the CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	CACertificate *string `json:"CACertificate,omitempty" xml:"CACertificate,omitempty"`
	// The CA certificate name.
	//
	// example:
	//
	// mycacert01
	CACertificateName *string `json:"CACertificateName,omitempty" xml:"CACertificateName,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the CA certificates.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2401682.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rto*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	//
	// example:
	//
	// UploadCACertificate
	Tag []*UploadCACertificateRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UploadCACertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadCACertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadCACertificateRequest) GetCACertificate() *string {
	return s.CACertificate
}

func (s *UploadCACertificateRequest) GetCACertificateName() *string {
	return s.CACertificateName
}

func (s *UploadCACertificateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UploadCACertificateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UploadCACertificateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadCACertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UploadCACertificateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UploadCACertificateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UploadCACertificateRequest) GetTag() []*UploadCACertificateRequestTag {
	return s.Tag
}

func (s *UploadCACertificateRequest) SetCACertificate(v string) *UploadCACertificateRequest {
	s.CACertificate = &v
	return s
}

func (s *UploadCACertificateRequest) SetCACertificateName(v string) *UploadCACertificateRequest {
	s.CACertificateName = &v
	return s
}

func (s *UploadCACertificateRequest) SetOwnerAccount(v string) *UploadCACertificateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UploadCACertificateRequest) SetOwnerId(v int64) *UploadCACertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadCACertificateRequest) SetRegionId(v string) *UploadCACertificateRequest {
	s.RegionId = &v
	return s
}

func (s *UploadCACertificateRequest) SetResourceGroupId(v string) *UploadCACertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UploadCACertificateRequest) SetResourceOwnerAccount(v string) *UploadCACertificateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UploadCACertificateRequest) SetResourceOwnerId(v int64) *UploadCACertificateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UploadCACertificateRequest) SetTag(v []*UploadCACertificateRequestTag) *UploadCACertificateRequest {
	s.Tag = v
	return s
}

func (s *UploadCACertificateRequest) Validate() error {
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

type UploadCACertificateRequestTag struct {
	// The key of tag N. Valid values of N: **1*	- to **20**. The tag key cannot be an empty string. The tag key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: **1 to 20**. The tag value can be an empty string. The tag value must be 1 to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UploadCACertificateRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UploadCACertificateRequestTag) GoString() string {
	return s.String()
}

func (s *UploadCACertificateRequestTag) GetKey() *string {
	return s.Key
}

func (s *UploadCACertificateRequestTag) GetValue() *string {
	return s.Value
}

func (s *UploadCACertificateRequestTag) SetKey(v string) *UploadCACertificateRequestTag {
	s.Key = &v
	return s
}

func (s *UploadCACertificateRequestTag) SetValue(v string) *UploadCACertificateRequestTag {
	s.Value = &v
	return s
}

func (s *UploadCACertificateRequestTag) Validate() error {
	return dara.Validate(s)
}
