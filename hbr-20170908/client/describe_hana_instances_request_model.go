// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHanaInstancesRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeHanaInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaInstancesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeHanaInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeHanaInstancesRequestTag) *DescribeHanaInstancesRequest
	GetTag() []*DescribeHanaInstancesRequestTag
	SetVaultId(v string) *DescribeHanaInstancesRequest
	GetVaultId() *string
}

type DescribeHanaInstancesRequest struct {
	// The ID of the SAP HANA instance.
	//
	// example:
	//
	// cl-0001zfc******50pr3
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekz24ikcjyqjkq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the SAP HANA instance.
	Tag []*DescribeHanaInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000b0ov******6zs
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeHanaInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHanaInstancesRequest) GetTag() []*DescribeHanaInstancesRequestTag {
	return s.Tag
}

func (s *DescribeHanaInstancesRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaInstancesRequest) SetClusterId(v string) *DescribeHanaInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetPageNumber(v int32) *DescribeHanaInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetPageSize(v int32) *DescribeHanaInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetResourceGroupId(v string) *DescribeHanaInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaInstancesRequest) SetTag(v []*DescribeHanaInstancesRequestTag) *DescribeHanaInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeHanaInstancesRequest) SetVaultId(v string) *DescribeHanaInstancesRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaInstancesRequest) Validate() error {
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

type DescribeHanaInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// ace:rm:rgld
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// rg-acfmwutpyat2kwy
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHanaInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeHanaInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeHanaInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeHanaInstancesRequestTag) SetKey(v string) *DescribeHanaInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeHanaInstancesRequestTag) SetValue(v string) *DescribeHanaInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeHanaInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
