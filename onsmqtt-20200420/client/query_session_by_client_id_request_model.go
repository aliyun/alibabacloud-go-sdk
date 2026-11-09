// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionByClientIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *QuerySessionByClientIdRequest
	GetClientId() *string
	SetInstanceId(v string) *QuerySessionByClientIdRequest
	GetInstanceId() *string
}

type QuerySessionByClientIdRequest struct {
	// The ID of the ApsaraMQ for MQTT client that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details*	- page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QuerySessionByClientIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionByClientIdRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionByClientIdRequest) GetClientId() *string {
	return s.ClientId
}

func (s *QuerySessionByClientIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuerySessionByClientIdRequest) SetClientId(v string) *QuerySessionByClientIdRequest {
	s.ClientId = &v
	return s
}

func (s *QuerySessionByClientIdRequest) SetInstanceId(v string) *QuerySessionByClientIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySessionByClientIdRequest) Validate() error {
	return dara.Validate(s)
}
