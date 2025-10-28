// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDefineRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteUserDefineRegionRequest
	GetId() *int64
	SetRegionTag(v string) *DeleteUserDefineRegionRequest
	GetRegionTag() *string
}

type DeleteUserDefineRegionRequest struct {
	// The unique ID of the custom namespace. You can call the ListUserDefineRegion operation to query the ID. For more information, see [ListUserDefineRegion](https://help.aliyun.com/document_detail/149377.html).
	//
	// example:
	//
	// 2564
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tag of the custom namespace.
	//
	// example:
	//
	// regiontag
	RegionTag *string `json:"RegionTag,omitempty" xml:"RegionTag,omitempty"`
}

func (s DeleteUserDefineRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDefineRegionRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDefineRegionRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteUserDefineRegionRequest) GetRegionTag() *string {
	return s.RegionTag
}

func (s *DeleteUserDefineRegionRequest) SetId(v int64) *DeleteUserDefineRegionRequest {
	s.Id = &v
	return s
}

func (s *DeleteUserDefineRegionRequest) SetRegionTag(v string) *DeleteUserDefineRegionRequest {
	s.RegionTag = &v
	return s
}

func (s *DeleteUserDefineRegionRequest) Validate() error {
	return dara.Validate(s)
}
