// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteRoutineRequest
	GetName() *string
}

type DeleteRoutineRequest struct {
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteRoutineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineRequest) SetName(v string) *DeleteRoutineRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineRequest) Validate() error {
	return dara.Validate(s)
}
