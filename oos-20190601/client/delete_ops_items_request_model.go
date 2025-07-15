// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpsItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpsItemIds(v []*string) *DeleteOpsItemsRequest
	GetOpsItemIds() []*string
	SetRegionId(v string) *DeleteOpsItemsRequest
	GetRegionId() *string
}

type DeleteOpsItemsRequest struct {
	// The IDs of O\\&M items.
	OpsItemIds []*string `json:"OpsItemIds,omitempty" xml:"OpsItemIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteOpsItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpsItemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpsItemsRequest) GetOpsItemIds() []*string {
	return s.OpsItemIds
}

func (s *DeleteOpsItemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOpsItemsRequest) SetOpsItemIds(v []*string) *DeleteOpsItemsRequest {
	s.OpsItemIds = v
	return s
}

func (s *DeleteOpsItemsRequest) SetRegionId(v string) *DeleteOpsItemsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOpsItemsRequest) Validate() error {
	return dara.Validate(s)
}
