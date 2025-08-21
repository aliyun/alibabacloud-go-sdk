// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthPassword(v string) *CreateStorageVolumeRequest
	GetAuthPassword() *string
	SetAuthProtocol(v string) *CreateStorageVolumeRequest
	GetAuthProtocol() *string
	SetAuthUser(v string) *CreateStorageVolumeRequest
	GetAuthUser() *string
	SetDescription(v string) *CreateStorageVolumeRequest
	GetDescription() *string
	SetEnsRegionId(v string) *CreateStorageVolumeRequest
	GetEnsRegionId() *string
	SetGatewayId(v string) *CreateStorageVolumeRequest
	GetGatewayId() *string
	SetIsAuth(v string) *CreateStorageVolumeRequest
	GetIsAuth() *string
	SetIsEnable(v string) *CreateStorageVolumeRequest
	GetIsEnable() *string
	SetStorageId(v string) *CreateStorageVolumeRequest
	GetStorageId() *string
	SetVolumeName(v string) *CreateStorageVolumeRequest
	GetVolumeName() *string
}

type CreateStorageVolumeRequest struct {
	// The password of the CHAP protocol.
	//
	// example:
	//
	// Password
	AuthPassword *string `json:"AuthPassword,omitempty" xml:"AuthPassword,omitempty"`
	// The authentication protocol. Set the value to **CHAP**.
	//
	// example:
	//
	// CHAP
	AuthProtocol *string `json:"AuthProtocol,omitempty" xml:"AuthProtocol,omitempty"`
	// The username of the CHAP protocol.
	//
	// example:
	//
	// User
	AuthUser *string `json:"AuthUser,omitempty" xml:"AuthUser,omitempty"`
	// The description of the volume. The description must be 2 to 128 characters in length. The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-3
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgw-****
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// Specifies whether to enable authentication. Valid values:
	//
	// 	- **1**: Authentication is enabled.
	//
	// 	- **0*	- (default): Authentication is disabled.
	//
	// example:
	//
	// 0
	IsAuth *string `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	// Indicates whether the volume is enabled. Valid values:
	//
	// 	- **1*	- (default): The volume is enabled.
	//
	// 	- **0**: The volume is disabled.
	//
	// example:
	//
	// 1
	IsEnable *string `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	// The ID of the storage medium.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-****
	StorageId *string `json:"StorageId,omitempty" xml:"StorageId,omitempty"`
	// The name of the volume. The name must be 2 to 128 characters in length. The name cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testVolumeName
	VolumeName *string `json:"VolumeName,omitempty" xml:"VolumeName,omitempty"`
}

func (s CreateStorageVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageVolumeRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageVolumeRequest) GetAuthPassword() *string {
	return s.AuthPassword
}

func (s *CreateStorageVolumeRequest) GetAuthProtocol() *string {
	return s.AuthProtocol
}

func (s *CreateStorageVolumeRequest) GetAuthUser() *string {
	return s.AuthUser
}

func (s *CreateStorageVolumeRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStorageVolumeRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateStorageVolumeRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateStorageVolumeRequest) GetIsAuth() *string {
	return s.IsAuth
}

func (s *CreateStorageVolumeRequest) GetIsEnable() *string {
	return s.IsEnable
}

func (s *CreateStorageVolumeRequest) GetStorageId() *string {
	return s.StorageId
}

func (s *CreateStorageVolumeRequest) GetVolumeName() *string {
	return s.VolumeName
}

func (s *CreateStorageVolumeRequest) SetAuthPassword(v string) *CreateStorageVolumeRequest {
	s.AuthPassword = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetAuthProtocol(v string) *CreateStorageVolumeRequest {
	s.AuthProtocol = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetAuthUser(v string) *CreateStorageVolumeRequest {
	s.AuthUser = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetDescription(v string) *CreateStorageVolumeRequest {
	s.Description = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetEnsRegionId(v string) *CreateStorageVolumeRequest {
	s.EnsRegionId = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetGatewayId(v string) *CreateStorageVolumeRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetIsAuth(v string) *CreateStorageVolumeRequest {
	s.IsAuth = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetIsEnable(v string) *CreateStorageVolumeRequest {
	s.IsEnable = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetStorageId(v string) *CreateStorageVolumeRequest {
	s.StorageId = &v
	return s
}

func (s *CreateStorageVolumeRequest) SetVolumeName(v string) *CreateStorageVolumeRequest {
	s.VolumeName = &v
	return s
}

func (s *CreateStorageVolumeRequest) Validate() error {
	return dara.Validate(s)
}
