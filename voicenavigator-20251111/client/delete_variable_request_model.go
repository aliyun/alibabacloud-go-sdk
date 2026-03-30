// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteVariableRequest
	GetInstanceId() *string
	SetVariableId(v string) *DeleteVariableRequest
	GetVariableId() *string
}

type DeleteVariableRequest struct {
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb1
	VariableId *string `json:"VariableId,omitempty" xml:"VariableId,omitempty"`
}

func (s DeleteVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteVariableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVariableRequest) GetVariableId() *string {
	return s.VariableId
}

func (s *DeleteVariableRequest) SetInstanceId(v string) *DeleteVariableRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVariableRequest) SetVariableId(v string) *DeleteVariableRequest {
	s.VariableId = &v
	return s
}

func (s *DeleteVariableRequest) Validate() error {
	return dara.Validate(s)
}
