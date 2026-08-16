// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppPolicyId(v string) *ModifyAppPolicyRequest
	GetAppPolicyId() *string
	SetProductType(v string) *ModifyAppPolicyRequest
	GetProductType() *string
	SetVideoPolicy(v *ModifyAppPolicyRequestVideoPolicy) *ModifyAppPolicyRequest
	GetVideoPolicy() *ModifyAppPolicyRequestVideoPolicy
}

type ModifyAppPolicyRequest struct {
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-ee2znjktwgxu2****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The display policy.
	VideoPolicy *ModifyAppPolicyRequestVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
}

func (s ModifyAppPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyRequest) GetAppPolicyId() *string {
	return s.AppPolicyId
}

func (s *ModifyAppPolicyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyAppPolicyRequest) GetVideoPolicy() *ModifyAppPolicyRequestVideoPolicy {
	return s.VideoPolicy
}

func (s *ModifyAppPolicyRequest) SetAppPolicyId(v string) *ModifyAppPolicyRequest {
	s.AppPolicyId = &v
	return s
}

func (s *ModifyAppPolicyRequest) SetProductType(v string) *ModifyAppPolicyRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppPolicyRequest) SetVideoPolicy(v *ModifyAppPolicyRequestVideoPolicy) *ModifyAppPolicyRequest {
	s.VideoPolicy = v
	return s
}

func (s *ModifyAppPolicyRequest) Validate() error {
	if s.VideoPolicy != nil {
		if err := s.VideoPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppPolicyRequestVideoPolicy struct {
	// The frame rate (FPS).
	//
	// example:
	//
	// 60
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// The height of the resolution, in pixels.
	//
	// example:
	//
	// 1080
	SessionResolutionHeight *int32 `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	// The width of the resolution, in pixels.
	//
	// example:
	//
	// 1920
	SessionResolutionWidth *int32 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	// The streaming mode. Used together with the Webrtc parameter to specify the protocol type.
	//
	// - Webrtc=`true` and StreamingMode=`video`: WebRTC stream.
	//
	// - Webrtc=`false` and StreamingMode=`video`: video stream.
	//
	// - Webrtc=`false` and StreamingMode=`mix`: mixed stream.
	//
	// example:
	//
	// video
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// Specifies whether to use adaptive resolution.
	//
	// - `true`: The session resolution follows changes in the terminal display area. In this case, SessionResolutionWidth and SessionResolutionHeight specify the maximum resolution values.
	//
	// - `false`: The session resolution does not follow changes in the terminal display area. In this case, the resolution is fixed to the values of SessionResolutionWidth and SessionResolutionHeight.
	//
	// example:
	//
	// false
	TerminalResolutionAdaptive *bool `json:"TerminalResolutionAdaptive,omitempty" xml:"TerminalResolutionAdaptive,omitempty"`
	// The visual quality strategy.
	//
	// example:
	//
	// smooth
	VisualQualityStrategy *string `json:"VisualQualityStrategy,omitempty" xml:"VisualQualityStrategy,omitempty"`
	// Specifies whether to enable WebRTC. Used together with the StreamingMode parameter to specify the protocol type.
	//
	// - Webrtc=`true` and StreamingMode=`video`: WebRTC stream.
	//
	// - Webrtc=`false` and StreamingMode=`video`: video stream.
	//
	// - Webrtc=`false` and StreamingMode=`mix`: mixed stream.
	//
	// example:
	//
	// true
	Webrtc *bool `json:"Webrtc,omitempty" xml:"Webrtc,omitempty"`
}

func (s ModifyAppPolicyRequestVideoPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppPolicyRequestVideoPolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyRequestVideoPolicy) GetFrameRate() *int32 {
	return s.FrameRate
}

func (s *ModifyAppPolicyRequestVideoPolicy) GetSessionResolutionHeight() *int32 {
	return s.SessionResolutionHeight
}

func (s *ModifyAppPolicyRequestVideoPolicy) GetSessionResolutionWidth() *int32 {
	return s.SessionResolutionWidth
}

func (s *ModifyAppPolicyRequestVideoPolicy) GetStreamingMode() *string {
	return s.StreamingMode
}

func (s *ModifyAppPolicyRequestVideoPolicy) GetTerminalResolutionAdaptive() *bool {
	return s.TerminalResolutionAdaptive
}

func (s *ModifyAppPolicyRequestVideoPolicy) GetVisualQualityStrategy() *string {
	return s.VisualQualityStrategy
}

func (s *ModifyAppPolicyRequestVideoPolicy) GetWebrtc() *bool {
	return s.Webrtc
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetFrameRate(v int32) *ModifyAppPolicyRequestVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetSessionResolutionHeight(v int32) *ModifyAppPolicyRequestVideoPolicy {
	s.SessionResolutionHeight = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetSessionResolutionWidth(v int32) *ModifyAppPolicyRequestVideoPolicy {
	s.SessionResolutionWidth = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetStreamingMode(v string) *ModifyAppPolicyRequestVideoPolicy {
	s.StreamingMode = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetTerminalResolutionAdaptive(v bool) *ModifyAppPolicyRequestVideoPolicy {
	s.TerminalResolutionAdaptive = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetVisualQualityStrategy(v string) *ModifyAppPolicyRequestVideoPolicy {
	s.VisualQualityStrategy = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetWebrtc(v bool) *ModifyAppPolicyRequestVideoPolicy {
	s.Webrtc = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) Validate() error {
	return dara.Validate(s)
}
