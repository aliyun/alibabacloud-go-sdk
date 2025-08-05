// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAirEcsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcsInstanceId(v string) *DeleteAirEcsInstanceRequest
	GetEcsInstanceId() *string
	SetUninstallClientSourceTypes(v []*string) *DeleteAirEcsInstanceRequest
	GetUninstallClientSourceTypes() []*string
}

type DeleteAirEcsInstanceRequest struct {
	// The ID of the Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// i-uf6ir9y******hvisj
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The data sources for which the client needs to be uninstalled.
	UninstallClientSourceTypes []*string `json:"UninstallClientSourceTypes,omitempty" xml:"UninstallClientSourceTypes,omitempty" type:"Repeated"`
}

func (s DeleteAirEcsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirEcsInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAirEcsInstanceRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DeleteAirEcsInstanceRequest) GetUninstallClientSourceTypes() []*string {
	return s.UninstallClientSourceTypes
}

func (s *DeleteAirEcsInstanceRequest) SetEcsInstanceId(v string) *DeleteAirEcsInstanceRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *DeleteAirEcsInstanceRequest) SetUninstallClientSourceTypes(v []*string) *DeleteAirEcsInstanceRequest {
	s.UninstallClientSourceTypes = v
	return s
}

func (s *DeleteAirEcsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
