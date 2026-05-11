// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int64) *ModifyInstanceRequest
	GetConcurrency() *int64
	SetDescription(v string) *ModifyInstanceRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyInstanceRequest
	GetInstanceId() *string
	SetName(v string) *ModifyInstanceRequest
	GetName() *string
}

type ModifyInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Concurrency *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *ModifyInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceRequest) GetName() *string {
	return s.Name
}

func (s *ModifyInstanceRequest) SetConcurrency(v int64) *ModifyInstanceRequest {
	s.Concurrency = &v
	return s
}

func (s *ModifyInstanceRequest) SetDescription(v string) *ModifyInstanceRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetName(v string) *ModifyInstanceRequest {
	s.Name = &v
	return s
}

func (s *ModifyInstanceRequest) Validate() error {
	return dara.Validate(s)
}
