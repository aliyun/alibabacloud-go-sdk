// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisterCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMqttInstanceId(v string) *GetRegisterCodeRequest
	GetMqttInstanceId() *string
}

type GetRegisterCodeRequest struct {
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-7mz2d******
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s GetRegisterCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegisterCodeRequest) GoString() string {
	return s.String()
}

func (s *GetRegisterCodeRequest) GetMqttInstanceId() *string {
	return s.MqttInstanceId
}

func (s *GetRegisterCodeRequest) SetMqttInstanceId(v string) *GetRegisterCodeRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *GetRegisterCodeRequest) Validate() error {
	return dara.Validate(s)
}
