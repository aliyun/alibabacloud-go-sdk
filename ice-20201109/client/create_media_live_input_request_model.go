// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputSettings(v []*CreateMediaLiveInputRequestInputSettings) *CreateMediaLiveInputRequest
	GetInputSettings() []*CreateMediaLiveInputRequestInputSettings
	SetName(v string) *CreateMediaLiveInputRequest
	GetName() *string
	SetSecurityGroupIds(v []*string) *CreateMediaLiveInputRequest
	GetSecurityGroupIds() []*string
	SetType(v string) *CreateMediaLiveInputRequest
	GetType() *string
}

type CreateMediaLiveInputRequest struct {
	// The input settings. An input can have up to two sources: primary and backup sources.
	//
	// This parameter is required.
	InputSettings []*CreateMediaLiveInputRequestInputSettings `json:"InputSettings,omitempty" xml:"InputSettings,omitempty" type:"Repeated"`
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
	// The input type. Valid values: RTMP_PUSH, RTMP_PULL, SRT_PUSH, SRT_PULL, and MEDIA_CONNECT.
	//
	// This parameter is required.
	//
	// example:
	//
	// RTMP_PUSH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMediaLiveInputRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputRequest) GetInputSettings() []*CreateMediaLiveInputRequestInputSettings {
	return s.InputSettings
}

func (s *CreateMediaLiveInputRequest) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveInputRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateMediaLiveInputRequest) GetType() *string {
	return s.Type
}

func (s *CreateMediaLiveInputRequest) SetInputSettings(v []*CreateMediaLiveInputRequestInputSettings) *CreateMediaLiveInputRequest {
	s.InputSettings = v
	return s
}

func (s *CreateMediaLiveInputRequest) SetName(v string) *CreateMediaLiveInputRequest {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveInputRequest) SetSecurityGroupIds(v []*string) *CreateMediaLiveInputRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateMediaLiveInputRequest) SetType(v string) *CreateMediaLiveInputRequest {
	s.Type = &v
	return s
}

func (s *CreateMediaLiveInputRequest) Validate() error {
	if s.InputSettings != nil {
		for _, item := range s.InputSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMediaLiveInputRequestInputSettings struct {
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

func (s CreateMediaLiveInputRequestInputSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputRequestInputSettings) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputRequestInputSettings) GetFlowId() *string {
	return s.FlowId
}

func (s *CreateMediaLiveInputRequestInputSettings) GetFlowOutputName() *string {
	return s.FlowOutputName
}

func (s *CreateMediaLiveInputRequestInputSettings) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *CreateMediaLiveInputRequestInputSettings) GetSrtLatency() *int32 {
	return s.SrtLatency
}

func (s *CreateMediaLiveInputRequestInputSettings) GetSrtMaxBitrate() *int32 {
	return s.SrtMaxBitrate
}

func (s *CreateMediaLiveInputRequestInputSettings) GetSrtPassphrase() *string {
	return s.SrtPassphrase
}

func (s *CreateMediaLiveInputRequestInputSettings) GetSrtPbKeyLen() *int32 {
	return s.SrtPbKeyLen
}

func (s *CreateMediaLiveInputRequestInputSettings) GetStreamName() *string {
	return s.StreamName
}

func (s *CreateMediaLiveInputRequestInputSettings) SetFlowId(v string) *CreateMediaLiveInputRequestInputSettings {
	s.FlowId = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) SetFlowOutputName(v string) *CreateMediaLiveInputRequestInputSettings {
	s.FlowOutputName = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) SetSourceUrl(v string) *CreateMediaLiveInputRequestInputSettings {
	s.SourceUrl = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) SetSrtLatency(v int32) *CreateMediaLiveInputRequestInputSettings {
	s.SrtLatency = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) SetSrtMaxBitrate(v int32) *CreateMediaLiveInputRequestInputSettings {
	s.SrtMaxBitrate = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) SetSrtPassphrase(v string) *CreateMediaLiveInputRequestInputSettings {
	s.SrtPassphrase = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) SetSrtPbKeyLen(v int32) *CreateMediaLiveInputRequestInputSettings {
	s.SrtPbKeyLen = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) SetStreamName(v string) *CreateMediaLiveInputRequestInputSettings {
	s.StreamName = &v
	return s
}

func (s *CreateMediaLiveInputRequestInputSettings) Validate() error {
	return dara.Validate(s)
}
