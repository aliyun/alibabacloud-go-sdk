// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAirEcsInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcsInstanceId(v string) *DeleteAirEcsInstanceShrinkRequest
	GetEcsInstanceId() *string
	SetUninstallClientSourceTypesShrink(v string) *DeleteAirEcsInstanceShrinkRequest
	GetUninstallClientSourceTypesShrink() *string
}

type DeleteAirEcsInstanceShrinkRequest struct {
	// The ID of the Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// i-uf6ir9y******hvisj
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The data sources for which the client needs to be uninstalled.
	UninstallClientSourceTypesShrink *string `json:"UninstallClientSourceTypes,omitempty" xml:"UninstallClientSourceTypes,omitempty"`
}

func (s DeleteAirEcsInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirEcsInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAirEcsInstanceShrinkRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DeleteAirEcsInstanceShrinkRequest) GetUninstallClientSourceTypesShrink() *string {
	return s.UninstallClientSourceTypesShrink
}

func (s *DeleteAirEcsInstanceShrinkRequest) SetEcsInstanceId(v string) *DeleteAirEcsInstanceShrinkRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *DeleteAirEcsInstanceShrinkRequest) SetUninstallClientSourceTypesShrink(v string) *DeleteAirEcsInstanceShrinkRequest {
	s.UninstallClientSourceTypesShrink = &v
	return s
}

func (s *DeleteAirEcsInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
