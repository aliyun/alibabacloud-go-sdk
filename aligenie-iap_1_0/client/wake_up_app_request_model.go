// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeUpAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsDebug(v bool) *WakeUpAppRequest
	GetIsDebug() *bool
	SetPath(v string) *WakeUpAppRequest
	GetPath() *string
	SetTargetInfo(v *WakeUpAppRequestTargetInfo) *WakeUpAppRequest
	GetTargetInfo() *WakeUpAppRequestTargetInfo
}

type WakeUpAppRequest struct {
	// The debug identity. When the skill application is in the debugging phase, this setting determines whether the retrieved skill application information is from the editing state or the online state.
	//
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// The path to launch the application, similar to the Link you enter in the skill application console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 应用拉起路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Device Information of the target to be launched.
	//
	// This parameter is required.
	TargetInfo *WakeUpAppRequestTargetInfo `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty" type:"Struct"`
}

func (s WakeUpAppRequest) String() string {
	return dara.Prettify(s)
}

func (s WakeUpAppRequest) GoString() string {
	return s.String()
}

func (s *WakeUpAppRequest) GetIsDebug() *bool {
	return s.IsDebug
}

func (s *WakeUpAppRequest) GetPath() *string {
	return s.Path
}

func (s *WakeUpAppRequest) GetTargetInfo() *WakeUpAppRequestTargetInfo {
	return s.TargetInfo
}

func (s *WakeUpAppRequest) SetIsDebug(v bool) *WakeUpAppRequest {
	s.IsDebug = &v
	return s
}

func (s *WakeUpAppRequest) SetPath(v string) *WakeUpAppRequest {
	s.Path = &v
	return s
}

func (s *WakeUpAppRequest) SetTargetInfo(v *WakeUpAppRequestTargetInfo) *WakeUpAppRequest {
	s.TargetInfo = v
	return s
}

func (s *WakeUpAppRequest) Validate() error {
	if s.TargetInfo != nil {
		if err := s.TargetInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WakeUpAppRequestTargetInfo struct {
	// The value corresponding to the encoding type.
	//
	// - When the encoding type is SKILLID, the value is the skill ID obtained from the webhook service.
	//
	// - When the encoding type is PACKAGENAME, the value is the packageName of the corresponding client application.
	//
	// This parameter is required.
	//
	// example:
	//
	// apk包名
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// Encoding type. There are multiple ways to obtain the MeowGenie device ID, and each method corresponds to a different encoding type:
	//
	// - PACKAGE_NAME: APK package name
	//
	// - SKILL_ID: Skill ID
	//
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// Organization ID. This is required only when the push type is `XX_UNION_XX`. If you have multiple methods to obtain MeowGenie device IDs and need interoperability among them, you must request an organization from the platform. After approval, the platform will assign an organization ID.
	//
	// example:
	//
	// 11
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The identity value corresponding to the Push Target Type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2VpiDQ6aMjxz******Eo7r6e08oIVZ3fKrm5TyEfY=
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
	// Push Target Type, which indicates the type of the obtained device identity:
	//
	// - `DEVICE_UNION_ID`: Device unionId
	//
	// - `DEVICE_OPEN_ID`: Device openId
	//
	// This parameter is required.
	//
	// example:
	//
	// DEVICE_OPEN_ID
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s WakeUpAppRequestTargetInfo) String() string {
	return dara.Prettify(s)
}

func (s WakeUpAppRequestTargetInfo) GoString() string {
	return s.String()
}

func (s *WakeUpAppRequestTargetInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *WakeUpAppRequestTargetInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *WakeUpAppRequestTargetInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *WakeUpAppRequestTargetInfo) GetTargetIdentity() *string {
	return s.TargetIdentity
}

func (s *WakeUpAppRequestTargetInfo) GetTargetType() *string {
	return s.TargetType
}

func (s *WakeUpAppRequestTargetInfo) SetEncodeKey(v string) *WakeUpAppRequestTargetInfo {
	s.EncodeKey = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetEncodeType(v string) *WakeUpAppRequestTargetInfo {
	s.EncodeType = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetOrganizationId(v string) *WakeUpAppRequestTargetInfo {
	s.OrganizationId = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetTargetIdentity(v string) *WakeUpAppRequestTargetInfo {
	s.TargetIdentity = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetTargetType(v string) *WakeUpAppRequestTargetInfo {
	s.TargetType = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) Validate() error {
	return dara.Validate(s)
}
