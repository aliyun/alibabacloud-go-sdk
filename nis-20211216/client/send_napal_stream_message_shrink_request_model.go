// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNapalStreamMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationShrink(v string) *SendNapalStreamMessageShrinkRequest
	GetConfigurationShrink() *string
	SetMessageShrink(v string) *SendNapalStreamMessageShrinkRequest
	GetMessageShrink() *string
	SetMetadataShrink(v string) *SendNapalStreamMessageShrinkRequest
	GetMetadataShrink() *string
}

type SendNapalStreamMessageShrinkRequest struct {
	// The request configuration object.
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The message object that contains user input and session context information.
	MessageShrink *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The additional request information.
	MetadataShrink *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
}

func (s SendNapalStreamMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageShrinkRequest) GetConfigurationShrink() *string {
	return s.ConfigurationShrink
}

func (s *SendNapalStreamMessageShrinkRequest) GetMessageShrink() *string {
	return s.MessageShrink
}

func (s *SendNapalStreamMessageShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *SendNapalStreamMessageShrinkRequest) SetConfigurationShrink(v string) *SendNapalStreamMessageShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *SendNapalStreamMessageShrinkRequest) SetMessageShrink(v string) *SendNapalStreamMessageShrinkRequest {
	s.MessageShrink = &v
	return s
}

func (s *SendNapalStreamMessageShrinkRequest) SetMetadataShrink(v string) *SendNapalStreamMessageShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *SendNapalStreamMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
