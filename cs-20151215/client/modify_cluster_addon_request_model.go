// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyClusterAddonRequest
	GetConfig() *string
}

type ModifyClusterAddonRequest struct {
	// The custom parameter settings that you want to use.
	//
	// example:
	//
	// {"CpuRequest":"800m"}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s ModifyClusterAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAddonRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAddonRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyClusterAddonRequest) SetConfig(v string) *ModifyClusterAddonRequest {
	s.Config = &v
	return s
}

func (s *ModifyClusterAddonRequest) Validate() error {
	return dara.Validate(s)
}
