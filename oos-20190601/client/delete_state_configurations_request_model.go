// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStateConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteStateConfigurationsRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteStateConfigurationsRequest
	GetRegionId() *string
	SetStateConfigurationIds(v string) *DeleteStateConfigurationsRequest
	GetStateConfigurationIds() *string
}

type DeleteStateConfigurationsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// abcde3OARpx77No54nv6
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of desired-state configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sc-asfgdhj12345"]
	StateConfigurationIds *string `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
}

func (s DeleteStateConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStateConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteStateConfigurationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStateConfigurationsRequest) GetStateConfigurationIds() *string {
	return s.StateConfigurationIds
}

func (s *DeleteStateConfigurationsRequest) SetClientToken(v string) *DeleteStateConfigurationsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStateConfigurationsRequest) SetRegionId(v string) *DeleteStateConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStateConfigurationsRequest) SetStateConfigurationIds(v string) *DeleteStateConfigurationsRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *DeleteStateConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
