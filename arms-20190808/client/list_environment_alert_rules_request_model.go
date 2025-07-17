// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListEnvironmentAlertRulesRequest
	GetAddonName() *string
	SetEnvironmentId(v string) *ListEnvironmentAlertRulesRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListEnvironmentAlertRulesRequest
	GetRegionId() *string
	SetScene(v string) *ListEnvironmentAlertRulesRequest
	GetScene() *string
}

type ListEnvironmentAlertRulesRequest struct {
	// The name of the add-on. You must specify AddonName or Scene.
	//
	// example:
	//
	// mysql
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scenario of the add-on. You must specify AddonName or Scene.
	//
	// example:
	//
	// database
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListEnvironmentAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAlertRulesRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListEnvironmentAlertRulesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentAlertRulesRequest) GetScene() *string {
	return s.Scene
}

func (s *ListEnvironmentAlertRulesRequest) SetAddonName(v string) *ListEnvironmentAlertRulesRequest {
	s.AddonName = &v
	return s
}

func (s *ListEnvironmentAlertRulesRequest) SetEnvironmentId(v string) *ListEnvironmentAlertRulesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentAlertRulesRequest) SetRegionId(v string) *ListEnvironmentAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentAlertRulesRequest) SetScene(v string) *ListEnvironmentAlertRulesRequest {
	s.Scene = &v
	return s
}

func (s *ListEnvironmentAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}
