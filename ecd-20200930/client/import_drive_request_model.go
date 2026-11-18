// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *ImportDriveRequest
	GetDomainId() *string
	SetEndUserId(v string) *ImportDriveRequest
	GetEndUserId() *string
	SetForceUpdate(v bool) *ImportDriveRequest
	GetForceUpdate() *bool
	SetMountUrl(v string) *ImportDriveRequest
	GetMountUrl() *string
	SetPlatform(v string) *ImportDriveRequest
	GetPlatform() *string
	SetRegionId(v string) *ImportDriveRequest
	GetRegionId() *string
}

type ImportDriveRequest struct {
	// The target UPM domain ID. For workspace scenarios, use the UpmDomainId instead of the NAS file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dom-aaaa****
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The target ECD end user ID. The caller account is obtained from the authentication context.
	//
	// This parameter is required.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Specifies whether to force an update. Default value: false. If the file already exists and this parameter is set to false, the request fails. If this parameter is set to true, the original DriveId is reused for the update.
	//
	// example:
	//
	// false
	ForceUpdate *bool `json:"ForceUpdate,omitempty" xml:"ForceUpdate,omitempty"`
	// The relative path of the cloud application UPM. It must be the same as the path specified during upload authorization. For cloud desktops, do not pass this parameter or pass an empty value.
	//
	// example:
	//
	// appv1/ab****
	MountUrl *string `json:"MountUrl,omitempty" xml:"MountUrl,omitempty"`
	// The target platform. Valid values: WINDOWS_10, WINDOWS_11, WINDOWS_SERVER_2019, WINDOWS_SERVER_2022, and WINDOWS_SERVER_2025.
	//
	// This parameter is required.
	//
	// example:
	//
	// WINDOWS_SERVER_2025
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region of the target UPM.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ImportDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportDriveRequest) GoString() string {
	return s.String()
}

func (s *ImportDriveRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *ImportDriveRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ImportDriveRequest) GetForceUpdate() *bool {
	return s.ForceUpdate
}

func (s *ImportDriveRequest) GetMountUrl() *string {
	return s.MountUrl
}

func (s *ImportDriveRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ImportDriveRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportDriveRequest) SetDomainId(v string) *ImportDriveRequest {
	s.DomainId = &v
	return s
}

func (s *ImportDriveRequest) SetEndUserId(v string) *ImportDriveRequest {
	s.EndUserId = &v
	return s
}

func (s *ImportDriveRequest) SetForceUpdate(v bool) *ImportDriveRequest {
	s.ForceUpdate = &v
	return s
}

func (s *ImportDriveRequest) SetMountUrl(v string) *ImportDriveRequest {
	s.MountUrl = &v
	return s
}

func (s *ImportDriveRequest) SetPlatform(v string) *ImportDriveRequest {
	s.Platform = &v
	return s
}

func (s *ImportDriveRequest) SetRegionId(v string) *ImportDriveRequest {
	s.RegionId = &v
	return s
}

func (s *ImportDriveRequest) Validate() error {
	return dara.Validate(s)
}
