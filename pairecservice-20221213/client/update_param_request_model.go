// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateParamRequest
	GetInstanceId() *string
	SetValue(v string) *UpdateParamRequest
	GetValue() *string
}

type UpdateParamRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// house
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateParamRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateParamRequest) GoString() string {
	return s.String()
}

func (s *UpdateParamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateParamRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateParamRequest) SetInstanceId(v string) *UpdateParamRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateParamRequest) SetValue(v string) *UpdateParamRequest {
	s.Value = &v
	return s
}

func (s *UpdateParamRequest) Validate() error {
	return dara.Validate(s)
}
