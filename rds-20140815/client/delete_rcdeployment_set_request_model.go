// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCDeploymentSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *DeleteRCDeploymentSetRequest
	GetDeploymentSetId() *string
	SetRegionId(v string) *DeleteRCDeploymentSetRequest
	GetRegionId() *string
}

type DeleteRCDeploymentSetRequest struct {
	// The deployment set ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds-uf6c8qerk019bj1l****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRCDeploymentSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCDeploymentSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCDeploymentSetRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DeleteRCDeploymentSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCDeploymentSetRequest) SetDeploymentSetId(v string) *DeleteRCDeploymentSetRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *DeleteRCDeploymentSetRequest) SetRegionId(v string) *DeleteRCDeploymentSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCDeploymentSetRequest) Validate() error {
	return dara.Validate(s)
}
