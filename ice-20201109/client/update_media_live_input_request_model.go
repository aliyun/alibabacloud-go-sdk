// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputId(v string) *UpdateMediaLiveInputRequest
	GetInputId() *string
	SetInputSettings(v []*UpdateMediaLiveInputRequestInputSettings) *UpdateMediaLiveInputRequest
	GetInputSettings() []*UpdateMediaLiveInputRequestInputSettings
	SetName(v string) *UpdateMediaLiveInputRequest
	GetName() *string
	SetSecurityGroupIds(v []*string) *UpdateMediaLiveInputRequest
	GetSecurityGroupIds() []*string
}

type UpdateMediaLiveInputRequest struct {
	// The ID of the input.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	// The input settings. An input can have up to two sources: primary and backup sources.
	//
	// This parameter is required.
	InputSettings []*UpdateMediaLiveInputRequestInputSettings `json:"InputSettings,omitempty" xml:"InputSettings,omitempty" type:"Repeated"`
	// The name of the input. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myinput
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the security groups to be associated with the input. This parameter is required for PUSH inputs.
	//
	// example:
	//
	// ["G6G4X5T4SZYPSTT5"]
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateMediaLiveInputRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputRequest) GetInputId() *string {
	return s.InputId
}

func (s *UpdateMediaLiveInputRequest) GetInputSettings() []*UpdateMediaLiveInputRequestInputSettings {
	return s.InputSettings
}

func (s *UpdateMediaLiveInputRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveInputRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *UpdateMediaLiveInputRequest) SetInputId(v string) *UpdateMediaLiveInputRequest {
	s.InputId = &v
	return s
}

func (s *UpdateMediaLiveInputRequest) SetInputSettings(v []*UpdateMediaLiveInputRequestInputSettings) *UpdateMediaLiveInputRequest {
	s.InputSettings = v
	return s
}

func (s *UpdateMediaLiveInputRequest) SetName(v string) *UpdateMediaLiveInputRequest {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveInputRequest) SetSecurityGroupIds(v []*string) *UpdateMediaLiveInputRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *UpdateMediaLiveInputRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateMediaLiveInputRequestInputSettings struct {
	// The ID of the flow from MediaConnect. This parameter is required when Type is set to MEDIA_CONNECT.
	//
	// example:
	//
	// ******81-9693-40dc-bbab-db5e49******
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The output name of the MediaConnect flow. This parameter is required when Type is set to MEDIA_CONNECT.
	//
	// example:
	//
	// myFlowOutputName
	FlowOutputName *string `json:"FlowOutputName,omitempty" xml:"FlowOutputName,omitempty"`
	// The source URL from which the stream is pulled. This parameter is required for PULL inputs.
	//
	// example:
	//
	// rtmp://domain/app/stream
	SourceUrl     *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
	SrtLatency    *int32  `json:"SrtLatency,omitempty" xml:"SrtLatency,omitempty"`
	SrtMaxBitrate *int32  `json:"SrtMaxBitrate,omitempty" xml:"SrtMaxBitrate,omitempty"`
	SrtPassphrase *string `json:"SrtPassphrase,omitempty" xml:"SrtPassphrase,omitempty"`
	SrtPbKeyLen   *int32  `json:"SrtPbKeyLen,omitempty" xml:"SrtPbKeyLen,omitempty"`
	// The name of the pushed stream. This parameter is required for PUSH inputs. It can be up to 255 characters in length.
	//
	// example:
	//
	// mystream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s UpdateMediaLiveInputRequestInputSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputRequestInputSettings) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetFlowId() *string {
	return s.FlowId
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetFlowOutputName() *string {
	return s.FlowOutputName
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetSrtMaxBitrate() *int32 {
	return s.SrtMaxBitrate
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetSrtPbKeyLen() *int32 {
	return s.SrtPbKeyLen
}

func (s *UpdateMediaLiveInputRequestInputSettings) GetStreamName() *string {
	return s.StreamName
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetFlowId(v string) *UpdateMediaLiveInputRequestInputSettings {
	s.FlowId = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetFlowOutputName(v string) *UpdateMediaLiveInputRequestInputSettings {
	s.FlowOutputName = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetSourceUrl(v string) *UpdateMediaLiveInputRequestInputSettings {
	s.SourceUrl = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetSrtLatency(v int32) *UpdateMediaLiveInputRequestInputSettings {
	s.SrtLatency = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetSrtMaxBitrate(v int32) *UpdateMediaLiveInputRequestInputSettings {
	s.SrtMaxBitrate = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetSrtPassphrase(v string) *UpdateMediaLiveInputRequestInputSettings {
	s.SrtPassphrase = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetSrtPbKeyLen(v int32) *UpdateMediaLiveInputRequestInputSettings {
	s.SrtPbKeyLen = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) SetStreamName(v string) *UpdateMediaLiveInputRequestInputSettings {
	s.StreamName = &v
	return s
}

func (s *UpdateMediaLiveInputRequestInputSettings) Validate() error {
	return dara.Validate(s)
}
