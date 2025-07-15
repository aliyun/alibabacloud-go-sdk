// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPackagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPackageId(v []*string) *DeleteNetworkPackagesRequest
	GetNetworkPackageId() []*string
	SetRegionId(v string) *DeleteNetworkPackagesRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *DeleteNetworkPackagesRequest
	GetResellerOwnerUid() *int64
}

type DeleteNetworkPackagesRequest struct {
	// The IDs of premium bandwidth plans. You can specify one or more IDs.
	//
	// This parameter is required.
	NetworkPackageId []*string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s DeleteNetworkPackagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPackagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesRequest) GetNetworkPackageId() []*string {
	return s.NetworkPackageId
}

func (s *DeleteNetworkPackagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNetworkPackagesRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DeleteNetworkPackagesRequest) SetNetworkPackageId(v []*string) *DeleteNetworkPackagesRequest {
	s.NetworkPackageId = v
	return s
}

func (s *DeleteNetworkPackagesRequest) SetRegionId(v string) *DeleteNetworkPackagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNetworkPackagesRequest) SetResellerOwnerUid(v int64) *DeleteNetworkPackagesRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DeleteNetworkPackagesRequest) Validate() error {
	return dara.Validate(s)
}
