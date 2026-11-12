// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMqttTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientOS(v string) *GetMqttTicketShrinkRequest
	GetClientOS() *string
	SetClientType(v int32) *GetMqttTicketShrinkRequest
	GetClientType() *int32
	SetClientVersion(v string) *GetMqttTicketShrinkRequest
	GetClientVersion() *string
	SetPropertiesShrink(v string) *GetMqttTicketShrinkRequest
	GetPropertiesShrink() *string
	SetUuid(v string) *GetMqttTicketShrinkRequest
	GetUuid() *string
}

type GetMqttTicketShrinkRequest struct {
	// The client operating system.
	//
	// This parameter is required.
	//
	// example:
	//
	// Windows 10
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client type. Valid values:
	//
	// - 0: hardware terminal
	//
	// - 1: Windows software terminal
	//
	// - 2: macOS software terminal
	//
	// - 3: iOS client
	//
	// - 4: Android client
	//
	// - 5: web client
	//
	// - 6: tenant console
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The client version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The extended properties of the client, represented as a JSON object where both keys and values are strings. Example: `{"a":"1","b":"2"}`. If this parameter is not specified, an empty property set is used.
	//
	// example:
	//
	// {"a":"1","b":"2"}
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The UUID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetMqttTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMqttTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMqttTicketShrinkRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GetMqttTicketShrinkRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetMqttTicketShrinkRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetMqttTicketShrinkRequest) GetPropertiesShrink() *string {
	return s.PropertiesShrink
}

func (s *GetMqttTicketShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetMqttTicketShrinkRequest) SetClientOS(v string) *GetMqttTicketShrinkRequest {
	s.ClientOS = &v
	return s
}

func (s *GetMqttTicketShrinkRequest) SetClientType(v int32) *GetMqttTicketShrinkRequest {
	s.ClientType = &v
	return s
}

func (s *GetMqttTicketShrinkRequest) SetClientVersion(v string) *GetMqttTicketShrinkRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetMqttTicketShrinkRequest) SetPropertiesShrink(v string) *GetMqttTicketShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *GetMqttTicketShrinkRequest) SetUuid(v string) *GetMqttTicketShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *GetMqttTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
