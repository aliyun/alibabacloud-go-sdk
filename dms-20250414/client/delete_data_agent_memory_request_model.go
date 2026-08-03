// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *DeleteDataAgentMemoryRequest
	GetDMSUnit() *string
	SetUuid(v string) *DeleteDataAgentMemoryRequest
	GetUuid() *string
}

type DeleteDataAgentMemoryRequest struct {
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The memory UUID.
	//
	// example:
	//
	// ed3f67***********ed
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DeleteDataAgentMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMemoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMemoryRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *DeleteDataAgentMemoryRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteDataAgentMemoryRequest) SetDMSUnit(v string) *DeleteDataAgentMemoryRequest {
	s.DMSUnit = &v
	return s
}

func (s *DeleteDataAgentMemoryRequest) SetUuid(v string) *DeleteDataAgentMemoryRequest {
	s.Uuid = &v
	return s
}

func (s *DeleteDataAgentMemoryRequest) Validate() error {
	return dara.Validate(s)
}
