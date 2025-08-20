// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChangeSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetId(v string) *DeleteChangeSetRequest
	GetChangeSetId() *string
	SetRegionId(v string) *DeleteChangeSetRequest
	GetRegionId() *string
}

type DeleteChangeSetRequest struct {
	// The ID of the change set.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1f6521a4-05af-4975-afe9-bc4b45ad****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteChangeSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChangeSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetRequest) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *DeleteChangeSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteChangeSetRequest) SetChangeSetId(v string) *DeleteChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

func (s *DeleteChangeSetRequest) SetRegionId(v string) *DeleteChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteChangeSetRequest) Validate() error {
	return dara.Validate(s)
}
