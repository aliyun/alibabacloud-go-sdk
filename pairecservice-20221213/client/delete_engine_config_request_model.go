// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEngineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteAll(v bool) *DeleteEngineConfigRequest
	GetDeleteAll() *bool
	SetInstanceId(v string) *DeleteEngineConfigRequest
	GetInstanceId() *string
}

type DeleteEngineConfigRequest struct {
	// Specifies whether to delete all engine configurations with the same name.
	//
	// example:
	//
	// False
	DeleteAll *bool `json:"DeleteAll,omitempty" xml:"DeleteAll,omitempty"`
	// The instance ID. For information about how to obtain the instance ID, see [ListInstances](https://help.aliyun.com/document_detail/2411819.html).
	//
	// example:
	//
	// pairec-cn-***test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteEngineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEngineConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteEngineConfigRequest) GetDeleteAll() *bool {
	return s.DeleteAll
}

func (s *DeleteEngineConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteEngineConfigRequest) SetDeleteAll(v bool) *DeleteEngineConfigRequest {
	s.DeleteAll = &v
	return s
}

func (s *DeleteEngineConfigRequest) SetInstanceId(v string) *DeleteEngineConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteEngineConfigRequest) Validate() error {
	return dara.Validate(s)
}
