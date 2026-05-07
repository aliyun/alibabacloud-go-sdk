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
	DeleteAll *bool `json:"DeleteAll,omitempty" xml:"DeleteAll,omitempty"`
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
