// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *CloseConnectionRequest
	GetClientId() *string
	SetInstanceId(v string) *CloseConnectionRequest
	GetInstanceId() *string
}

type CloseConnectionRequest struct {
	// The client ID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the MSMQ MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloseConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseConnectionRequest) GoString() string {
	return s.String()
}

func (s *CloseConnectionRequest) GetClientId() *string {
	return s.ClientId
}

func (s *CloseConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloseConnectionRequest) SetClientId(v string) *CloseConnectionRequest {
	s.ClientId = &v
	return s
}

func (s *CloseConnectionRequest) SetInstanceId(v string) *CloseConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *CloseConnectionRequest) Validate() error {
	return dara.Validate(s)
}
