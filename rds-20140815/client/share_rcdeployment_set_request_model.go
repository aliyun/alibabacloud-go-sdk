// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareRCDeploymentSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *ShareRCDeploymentSetRequest
	GetDeploymentSetId() *string
	SetRegionId(v string) *ShareRCDeploymentSetRequest
	GetRegionId() *string
}

type ShareRCDeploymentSetRequest struct {
	// This parameter is required.
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ShareRCDeploymentSetRequest) String() string {
	return dara.Prettify(s)
}

func (s ShareRCDeploymentSetRequest) GoString() string {
	return s.String()
}

func (s *ShareRCDeploymentSetRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ShareRCDeploymentSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ShareRCDeploymentSetRequest) SetDeploymentSetId(v string) *ShareRCDeploymentSetRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ShareRCDeploymentSetRequest) SetRegionId(v string) *ShareRCDeploymentSetRequest {
	s.RegionId = &v
	return s
}

func (s *ShareRCDeploymentSetRequest) Validate() error {
	return dara.Validate(s)
}
