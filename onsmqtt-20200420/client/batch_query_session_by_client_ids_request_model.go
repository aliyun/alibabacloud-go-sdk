// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQuerySessionByClientIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIdList(v []*string) *BatchQuerySessionByClientIdsRequest
	GetClientIdList() []*string
	SetInstanceId(v string) *BatchQuerySessionByClientIdsRequest
	GetInstanceId() *string
}

type BatchQuerySessionByClientIdsRequest struct {
	// The list of clients.
	//
	// This parameter is required.
	//
	// example:
	//
	// ClientIdList.1
	ClientIdList []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must match the actual instance ID used by the client. You can obtain the ID on the **Instance Details*	- page in the [console](https://mqtt.console.aliyun.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s BatchQuerySessionByClientIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQuerySessionByClientIdsRequest) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsRequest) GetClientIdList() []*string {
	return s.ClientIdList
}

func (s *BatchQuerySessionByClientIdsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchQuerySessionByClientIdsRequest) SetClientIdList(v []*string) *BatchQuerySessionByClientIdsRequest {
	s.ClientIdList = v
	return s
}

func (s *BatchQuerySessionByClientIdsRequest) SetInstanceId(v string) *BatchQuerySessionByClientIdsRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchQuerySessionByClientIdsRequest) Validate() error {
	return dara.Validate(s)
}
