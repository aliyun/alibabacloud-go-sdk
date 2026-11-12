// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMqttTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientOS(v string) *GetMqttTicketRequest
	GetClientOS() *string
	SetClientType(v int32) *GetMqttTicketRequest
	GetClientType() *int32
	SetClientVersion(v string) *GetMqttTicketRequest
	GetClientVersion() *string
	SetProperties(v map[string]interface{}) *GetMqttTicketRequest
	GetProperties() map[string]interface{}
	SetUuid(v string) *GetMqttTicketRequest
	GetUuid() *string
}

type GetMqttTicketRequest struct {
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
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The UUID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetMqttTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMqttTicketRequest) GoString() string {
	return s.String()
}

func (s *GetMqttTicketRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GetMqttTicketRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetMqttTicketRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetMqttTicketRequest) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetMqttTicketRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetMqttTicketRequest) SetClientOS(v string) *GetMqttTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetMqttTicketRequest) SetClientType(v int32) *GetMqttTicketRequest {
	s.ClientType = &v
	return s
}

func (s *GetMqttTicketRequest) SetClientVersion(v string) *GetMqttTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetMqttTicketRequest) SetProperties(v map[string]interface{}) *GetMqttTicketRequest {
	s.Properties = v
	return s
}

func (s *GetMqttTicketRequest) SetUuid(v string) *GetMqttTicketRequest {
	s.Uuid = &v
	return s
}

func (s *GetMqttTicketRequest) Validate() error {
	return dara.Validate(s)
}
