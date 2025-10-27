// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInclinedNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeInclinedNodesRequest
	GetDBClusterId() *string
	SetLang(v string) *DescribeInclinedNodesRequest
	GetLang() *string
	SetOwnerAccount(v string) *DescribeInclinedNodesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInclinedNodesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInclinedNodesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInclinedNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInclinedNodesRequest
	GetResourceOwnerId() *int64
}

type DescribeInclinedNodesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-ufxxxxxxxxxx3q1x1
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The language of file titles and error messages. Valid values:
	//
	// 	- **zh**: simplified Chinese.
	//
	// 	- **en**: English.
	//
	// 	- **ja**: Japanese.
	//
	// 	- **zh-tw**: traditional Chinese.
	//
	// example:
	//
	// zh
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInclinedNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInclinedNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInclinedNodesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeInclinedNodesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInclinedNodesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInclinedNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInclinedNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInclinedNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInclinedNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInclinedNodesRequest) SetDBClusterId(v string) *DescribeInclinedNodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeInclinedNodesRequest) SetLang(v string) *DescribeInclinedNodesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInclinedNodesRequest) SetOwnerAccount(v string) *DescribeInclinedNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInclinedNodesRequest) SetOwnerId(v int64) *DescribeInclinedNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInclinedNodesRequest) SetRegionId(v string) *DescribeInclinedNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInclinedNodesRequest) SetResourceOwnerAccount(v string) *DescribeInclinedNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInclinedNodesRequest) SetResourceOwnerId(v int64) *DescribeInclinedNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInclinedNodesRequest) Validate() error {
	return dara.Validate(s)
}
