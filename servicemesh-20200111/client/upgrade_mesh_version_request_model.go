// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMeshVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPreCheck(v bool) *UpgradeMeshVersionRequest
	GetPreCheck() *bool
	SetServiceMeshId(v string) *UpgradeMeshVersionRequest
	GetServiceMeshId() *string
}

type UpgradeMeshVersionRequest struct {
	// Specifies whether to perform a precheck. Default value: false. If this parameter is set to true, this call only checks whether the current ASM instance meets the upgrade conditions and does not actually perform an upgrade.
	//
	// example:
	//
	// false
	PreCheck *bool `json:"PreCheck,omitempty" xml:"PreCheck,omitempty"`
	// The ASM instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s UpgradeMeshVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMeshVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMeshVersionRequest) GetPreCheck() *bool {
	return s.PreCheck
}

func (s *UpgradeMeshVersionRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *UpgradeMeshVersionRequest) SetPreCheck(v bool) *UpgradeMeshVersionRequest {
	s.PreCheck = &v
	return s
}

func (s *UpgradeMeshVersionRequest) SetServiceMeshId(v string) *UpgradeMeshVersionRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *UpgradeMeshVersionRequest) Validate() error {
	return dara.Validate(s)
}
