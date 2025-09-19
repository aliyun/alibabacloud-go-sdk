// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClusterInterceptionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v string) *SetClusterInterceptionConfigRequest
	GetClusterIds() *string
	SetSwitchOn(v int32) *SetClusterInterceptionConfigRequest
	GetSwitchOn() *int32
	SetSwitchType(v int32) *SetClusterInterceptionConfigRequest
	GetSwitchType() *int32
}

type SetClusterInterceptionConfigRequest struct {
	// The ID of the cluster. Separate multiple cluster IDs with commas (,).
	//
	// > You can call the [ListClusterInterceptionConfig](~~ListClusterInterceptionConfig~~) operation to query the IDs of clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// c60b77fe62093480db6164a3c2fa****
	ClusterIds *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	// Specifies whether to turn on the switch. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SwitchOn *int32 `json:"SwitchOn,omitempty" xml:"SwitchOn,omitempty"`
	// The type of the switch that you want to configure. Valid values:
	//
	// 	- **0**: the interception switch
	//
	// 	- **1**: the interception type switch
	//
	// 	- **2**: the interception history switch
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SwitchType *int32 `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s SetClusterInterceptionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetClusterInterceptionConfigRequest) GoString() string {
	return s.String()
}

func (s *SetClusterInterceptionConfigRequest) GetClusterIds() *string {
	return s.ClusterIds
}

func (s *SetClusterInterceptionConfigRequest) GetSwitchOn() *int32 {
	return s.SwitchOn
}

func (s *SetClusterInterceptionConfigRequest) GetSwitchType() *int32 {
	return s.SwitchType
}

func (s *SetClusterInterceptionConfigRequest) SetClusterIds(v string) *SetClusterInterceptionConfigRequest {
	s.ClusterIds = &v
	return s
}

func (s *SetClusterInterceptionConfigRequest) SetSwitchOn(v int32) *SetClusterInterceptionConfigRequest {
	s.SwitchOn = &v
	return s
}

func (s *SetClusterInterceptionConfigRequest) SetSwitchType(v int32) *SetClusterInterceptionConfigRequest {
	s.SwitchType = &v
	return s
}

func (s *SetClusterInterceptionConfigRequest) Validate() error {
	return dara.Validate(s)
}
