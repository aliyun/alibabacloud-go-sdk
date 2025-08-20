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
	// example:
	//
	// true
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 应用拉起路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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
	return dara.Validate(s)
}

type WakeUpAppRequestTargetInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// apk包名
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PACKAGE_NAME
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// 11
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2VpiDQ6aMjxz******Eo7r6e08oIVZ3fKrm5TyEfY=
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
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
