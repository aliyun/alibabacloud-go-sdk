// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateThirdPartyAppLoginStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfo(v *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) *InvalidateThirdPartyAppLoginStateRequest
	GetDeviceInfo() *InvalidateThirdPartyAppLoginStateRequestDeviceInfo
	SetThirdPartyAppId(v string) *InvalidateThirdPartyAppLoginStateRequest
	GetThirdPartyAppId() *string
}

type InvalidateThirdPartyAppLoginStateRequest struct {
	// Device identification information
	//
	// This parameter is required.
	DeviceInfo *InvalidateThirdPartyAppLoginStateRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	// Third-party application identity
	//
	// This parameter is required.
	//
	// example:
	//
	// com.*.*.*
	ThirdPartyAppId *string `json:"ThirdPartyAppId,omitempty" xml:"ThirdPartyAppId,omitempty"`
}

func (s InvalidateThirdPartyAppLoginStateRequest) String() string {
	return dara.Prettify(s)
}

func (s InvalidateThirdPartyAppLoginStateRequest) GoString() string {
	return s.String()
}

func (s *InvalidateThirdPartyAppLoginStateRequest) GetDeviceInfo() *InvalidateThirdPartyAppLoginStateRequestDeviceInfo {
	return s.DeviceInfo
}

func (s *InvalidateThirdPartyAppLoginStateRequest) GetThirdPartyAppId() *string {
	return s.ThirdPartyAppId
}

func (s *InvalidateThirdPartyAppLoginStateRequest) SetDeviceInfo(v *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) *InvalidateThirdPartyAppLoginStateRequest {
	s.DeviceInfo = v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateRequest) SetThirdPartyAppId(v string) *InvalidateThirdPartyAppLoginStateRequest {
	s.ThirdPartyAppId = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateRequest) Validate() error {
	if s.DeviceInfo != nil {
		if err := s.DeviceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvalidateThirdPartyAppLoginStateRequestDeviceInfo struct {
	// The value corresponding to the encoding type. When the encoding type is SKILL_ID, the value is the Skill ID of the application. When the encoding type is PACKAGE_NAME, the value is the packageName of the corresponding client app.
	//
	// This parameter is required.
	//
	// example:
	//
	// com.*.*.*
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the Tmall Genie device ID, and each method corresponds to a different encoding type: PACKAGE_NAME refers to the APK package name, used in the Android application customer flow; SKILL_ID refers to the skill ID, used in the cloud-based flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Device ID (deviceOpenId or deviceUnionId)
	//
	// This parameter is required.
	//
	// example:
	//
	// rV/XSgPuxZjx/hN3iw8U+e8ou***lk1r43LWcVW6fvY1Rr4sEPFodpnA==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of Device ID. OPEN_ID: the default device identity. UNION_ID: the device identity at the organization dimension, which is available only after an organization has been requested on the Maojing Skills Application Open Platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// Organization ID. Required if IdType is UNION_ID.
	//
	// example:
	//
	// 123
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s InvalidateThirdPartyAppLoginStateRequestDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s InvalidateThirdPartyAppLoginStateRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) GetId() *string {
	return s.Id
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) SetEncodeKey(v string) *InvalidateThirdPartyAppLoginStateRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) SetEncodeType(v string) *InvalidateThirdPartyAppLoginStateRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) SetId(v string) *InvalidateThirdPartyAppLoginStateRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) SetIdType(v string) *InvalidateThirdPartyAppLoginStateRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) SetOrganizationId(v string) *InvalidateThirdPartyAppLoginStateRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateRequestDeviceInfo) Validate() error {
	return dara.Validate(s)
}
