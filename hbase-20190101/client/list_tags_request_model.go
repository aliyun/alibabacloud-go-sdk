// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListTagsRequest
	GetRegionId() *string
}

type ListTagsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagsRequest) SetRegionId(v string) *ListTagsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagsRequest) Validate() error {
	return dara.Validate(s)
}
