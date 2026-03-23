// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAgentPlatformRequest
	GetDBClusterId() *string
	SetName(v string) *DeleteAgentPlatformRequest
	GetName() *string
	SetRegionId(v string) *DeleteAgentPlatformRequest
	GetRegionId() *string
}

type DeleteAgentPlatformRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_platform
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAgentPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentPlatformRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAgentPlatformRequest) GetName() *string {
	return s.Name
}

func (s *DeleteAgentPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAgentPlatformRequest) SetDBClusterId(v string) *DeleteAgentPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAgentPlatformRequest) SetName(v string) *DeleteAgentPlatformRequest {
	s.Name = &v
	return s
}

func (s *DeleteAgentPlatformRequest) SetRegionId(v string) *DeleteAgentPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAgentPlatformRequest) Validate() error {
	return dara.Validate(s)
}
