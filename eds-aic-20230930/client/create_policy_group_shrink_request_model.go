// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCameraRedirect(v string) *CreatePolicyGroupShrinkRequest
	GetCameraRedirect() *string
	SetClipboard(v string) *CreatePolicyGroupShrinkRequest
	GetClipboard() *string
	SetHtml5FileTransfer(v string) *CreatePolicyGroupShrinkRequest
	GetHtml5FileTransfer() *string
	SetLocalDrive(v string) *CreatePolicyGroupShrinkRequest
	GetLocalDrive() *string
	SetLockResolution(v string) *CreatePolicyGroupShrinkRequest
	GetLockResolution() *string
	SetNetRedirectPolicyShrink(v string) *CreatePolicyGroupShrinkRequest
	GetNetRedirectPolicyShrink() *string
	SetPolicyGroupName(v string) *CreatePolicyGroupShrinkRequest
	GetPolicyGroupName() *string
	SetPolicyType(v string) *CreatePolicyGroupShrinkRequest
	GetPolicyType() *string
	SetResolutionHeight(v int32) *CreatePolicyGroupShrinkRequest
	GetResolutionHeight() *int32
	SetResolutionWidth(v int32) *CreatePolicyGroupShrinkRequest
	GetResolutionWidth() *int32
	SetWatermarkShrink(v string) *CreatePolicyGroupShrinkRequest
	GetWatermarkShrink() *string
}

type CreatePolicyGroupShrinkRequest struct {
	// Specifies whether to enable the webcam redirection feature.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	CameraRedirect *string `json:"CameraRedirect,omitempty" xml:"CameraRedirect,omitempty"`
	// The read/write permissions on the clipboard.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: read and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// readwrite
	Clipboard *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	// The file transfer policy of the Alibaba Cloud Workspace web client.
	//
	// Valid values:
	//
	// 	- all: File upload and download are supported.
	//
	// 	- download: Only file download is supported.
	//
	// 	- upload: Only file upload is supported.
	//
	// 	- off: File upload or download is forbidden.
	//
	// example:
	//
	// off
	Html5FileTransfer *string `json:"Html5FileTransfer,omitempty" xml:"Html5FileTransfer,omitempty"`
	// The read/write permissions on the on-premises drive.
	//
	// Valid values:
	//
	// 	- read: read-only.
	//
	// 	- readwrite: ready and write.
	//
	// 	- off: read/write disabled.
	//
	// example:
	//
	// off
	LocalDrive *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	// Specifies whether to lock the resolution.
	//
	// Valid values:
	//
	// 	- off
	//
	// 	- on
	//
	// example:
	//
	// off
	LockResolution *string `json:"LockResolution,omitempty" xml:"LockResolution,omitempty"`
	// The network redirection policy.
	NetRedirectPolicyShrink *string `json:"NetRedirectPolicy,omitempty" xml:"NetRedirectPolicy,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// defaultPolicy
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
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
	ResolutionWidth *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
	WatermarkShrink *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s CreatePolicyGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyGroupShrinkRequest) GetCameraRedirect() *string {
	return s.CameraRedirect
}

func (s *CreatePolicyGroupShrinkRequest) GetClipboard() *string {
	return s.Clipboard
}

func (s *CreatePolicyGroupShrinkRequest) GetHtml5FileTransfer() *string {
	return s.Html5FileTransfer
}

func (s *CreatePolicyGroupShrinkRequest) GetLocalDrive() *string {
	return s.LocalDrive
}

func (s *CreatePolicyGroupShrinkRequest) GetLockResolution() *string {
	return s.LockResolution
}

func (s *CreatePolicyGroupShrinkRequest) GetNetRedirectPolicyShrink() *string {
	return s.NetRedirectPolicyShrink
}

func (s *CreatePolicyGroupShrinkRequest) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *CreatePolicyGroupShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePolicyGroupShrinkRequest) GetResolutionHeight() *int32 {
	return s.ResolutionHeight
}

func (s *CreatePolicyGroupShrinkRequest) GetResolutionWidth() *int32 {
	return s.ResolutionWidth
}

func (s *CreatePolicyGroupShrinkRequest) GetWatermarkShrink() *string {
	return s.WatermarkShrink
}

func (s *CreatePolicyGroupShrinkRequest) SetCameraRedirect(v string) *CreatePolicyGroupShrinkRequest {
	s.CameraRedirect = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetClipboard(v string) *CreatePolicyGroupShrinkRequest {
	s.Clipboard = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetHtml5FileTransfer(v string) *CreatePolicyGroupShrinkRequest {
	s.Html5FileTransfer = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetLocalDrive(v string) *CreatePolicyGroupShrinkRequest {
	s.LocalDrive = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetLockResolution(v string) *CreatePolicyGroupShrinkRequest {
	s.LockResolution = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetNetRedirectPolicyShrink(v string) *CreatePolicyGroupShrinkRequest {
	s.NetRedirectPolicyShrink = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetPolicyGroupName(v string) *CreatePolicyGroupShrinkRequest {
	s.PolicyGroupName = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetPolicyType(v string) *CreatePolicyGroupShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetResolutionHeight(v int32) *CreatePolicyGroupShrinkRequest {
	s.ResolutionHeight = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetResolutionWidth(v int32) *CreatePolicyGroupShrinkRequest {
	s.ResolutionWidth = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) SetWatermarkShrink(v string) *CreatePolicyGroupShrinkRequest {
	s.WatermarkShrink = &v
	return s
}

func (s *CreatePolicyGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
