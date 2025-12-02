// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOnlineTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteOnlineTestRequest
	GetRegionId() *string
	SetResourceType(v string) *DeleteOnlineTestRequest
	GetResourceType() *string
}

type DeleteOnlineTestRequest struct {
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// image
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteOnlineTestRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOnlineTestRequest) GoString() string {
	return s.String()
}

func (s *DeleteOnlineTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOnlineTestRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteOnlineTestRequest) SetRegionId(v string) *DeleteOnlineTestRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOnlineTestRequest) SetResourceType(v string) *DeleteOnlineTestRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteOnlineTestRequest) Validate() error {
	return dara.Validate(s)
}
