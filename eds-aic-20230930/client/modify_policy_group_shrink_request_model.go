// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCameraRedirect(v string) *ModifyPolicyGroupShrinkRequest
	GetCameraRedirect() *string
	SetClipboard(v string) *ModifyPolicyGroupShrinkRequest
	GetClipboard() *string
	SetHtml5FileTransfer(v string) *ModifyPolicyGroupShrinkRequest
	GetHtml5FileTransfer() *string
	SetLocalDrive(v string) *ModifyPolicyGroupShrinkRequest
	GetLocalDrive() *string
	SetLockResolution(v string) *ModifyPolicyGroupShrinkRequest
	GetLockResolution() *string
	SetNetRedirectPolicyShrink(v string) *ModifyPolicyGroupShrinkRequest
	GetNetRedirectPolicyShrink() *string
	SetPolicyGroupId(v string) *ModifyPolicyGroupShrinkRequest
	GetPolicyGroupId() *string
	SetPolicyGroupName(v string) *ModifyPolicyGroupShrinkRequest
	GetPolicyGroupName() *string
	SetResolutionHeight(v int32) *ModifyPolicyGroupShrinkRequest
	GetResolutionHeight() *int32
	SetResolutionWidth(v int32) *ModifyPolicyGroupShrinkRequest
	GetResolutionWidth() *int32
	SetWatermarkShrink(v string) *ModifyPolicyGroupShrinkRequest
	GetWatermarkShrink() *string
}

type ModifyPolicyGroupShrinkRequest struct {
	// Specifies whether to enable local camera redirection.
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The clipboard permission.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy for the Wuying web client.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The local disk mapping permission.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// Network redirection.
	NetRedirectPolicyShrink *string `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-4bi18ebi9tfjh****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Default policy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The height of the resolution. Unit: pixels.
	//
	// example:
	//
	// 1280
	ResolutionHeight *int32 `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	// The width of the resolution. Unit: pixels.
	//
	// example:
	//
	// 720
	ResolutionWidth *int32 `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	// The screen watermark.
	WatermarkShrink *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s ModifyPolicyGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyGroupShrinkRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *ModifyPolicyGroupShrinkRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *ModifyPolicyGroupShrinkRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *ModifyPolicyGroupShrinkRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *ModifyPolicyGroupShrinkRequest) GetLockResolution() *string {
	return s.LockResolution
}

func (s *ModifyPolicyGroupShrinkRequest) GetNetRedirectPolicyShrink() *string {
	return s.NetRedirectPolicyShrink
}

func (s *ModifyPolicyGroupShrinkRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyPolicyGroupShrinkRequest) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *ModifyPolicyGroupShrinkRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *ModifyPolicyGroupShrinkRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *ModifyPolicyGroupShrinkRequest) GetWatermarkShrink() *string {
	return s.WatermarkShrink
}

func (s *ModifyPolicyGroupShrinkRequest) SetCameraRedirect(v string) *ModifyPolicyGroupShrinkRequest {
	s.CameraRedirect = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetClipboard(v string) *ModifyPolicyGroupShrinkRequest {
	s.Clipboard = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetHtml5FileTransfer(v string) *ModifyPolicyGroupShrinkRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetLocalDrive(v string) *ModifyPolicyGroupShrinkRequest {
	s.LocalDrive = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetLockResolution(v string) *ModifyPolicyGroupShrinkRequest {
	s.LockResolution = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetNetRedirectPolicyShrink(v string) *ModifyPolicyGroupShrinkRequest {
	s.NetRedirectPolicyShrink = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetPolicyGroupId(v string) *ModifyPolicyGroupShrinkRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetPolicyGroupName(v string) *ModifyPolicyGroupShrinkRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetResolutionHeight(v int32) *ModifyPolicyGroupShrinkRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetResolutionWidth(v int32) *ModifyPolicyGroupShrinkRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) SetWatermarkShrink(v string) *ModifyPolicyGroupShrinkRequest {
	s.WatermarkShrink = &v
	return s
}

func (s *ModifyPolicyGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
