// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBundlesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v []*string) *DeleteBundlesRequest
	GetBundleId() []*string
	SetRegionId(v string) *DeleteBundlesRequest
	GetRegionId() *string
}

type DeleteBundlesRequest struct {
	// The IDs of the cloud computer templates. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// b-cezrnfgecbich****
	BundleId []*string `json:"BundleId,omitempty" xml:"BundleId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBundlesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBundlesRequest) GoString() string {
	return s.String()
}

func (s *DeleteBundlesRequest) GetBundleId() []*string {
	return s.BundleId
}

func (s *DeleteBundlesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBundlesRequest) SetBundleId(v []*string) *DeleteBundlesRequest {
	s.BundleId = v
	return s
}

func (s *DeleteBundlesRequest) SetRegionId(v string) *DeleteBundlesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBundlesRequest) Validate() error {
	return dara.Validate(s)
}
