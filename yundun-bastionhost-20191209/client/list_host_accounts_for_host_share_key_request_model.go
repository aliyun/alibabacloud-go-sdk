// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForHostShareKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostShareKeyId(v string) *ListHostAccountsForHostShareKeyRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *ListHostAccountsForHostShareKeyRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListHostAccountsForHostShareKeyRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostAccountsForHostShareKeyRequest
	GetPageSize() *string
	SetRegionId(v string) *ListHostAccountsForHostShareKeyRequest
	GetRegionId() *string
}

type ListHostAccountsForHostShareKeyRequest struct {
	// The shared key ID.
	//
	// >  You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to query the shared key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1124
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostAccountsForHostShareKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *ListHostAccountsForHostShareKeyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostAccountsForHostShareKeyRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostAccountsForHostShareKeyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostAccountsForHostShareKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostAccountsForHostShareKeyRequest) SetHostShareKeyId(v string) *ListHostAccountsForHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetInstanceId(v string) *ListHostAccountsForHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetPageNumber(v string) *ListHostAccountsForHostShareKeyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetPageSize(v string) *ListHostAccountsForHostShareKeyRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetRegionId(v string) *ListHostAccountsForHostShareKeyRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) Validate() error {
	return dara.Validate(s)
}
