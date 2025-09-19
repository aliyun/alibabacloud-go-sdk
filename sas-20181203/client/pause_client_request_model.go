// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseClientRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuids(v string) *PauseClientRequest
	GetUuids() *string
	SetValue(v string) *PauseClientRequest
	GetValue() *string
}

type PauseClientRequest struct {
	// The UUIDs of servers for which you want to enable or disable the Security Center agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// uuid-1211-sadsd-2131
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	// The status of the Security Center agent. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PauseClientRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseClientRequest) GoString() string {
	return s.String()
}

func (s *PauseClientRequest) GetUuids() *string {
	return s.Uuids
}

func (s *PauseClientRequest) GetValue() *string {
	return s.Value
}

func (s *PauseClientRequest) SetUuids(v string) *PauseClientRequest {
	s.Uuids = &v
	return s
}

func (s *PauseClientRequest) SetValue(v string) *PauseClientRequest {
	s.Value = &v
	return s
}

func (s *PauseClientRequest) Validate() error {
	return dara.Validate(s)
}
