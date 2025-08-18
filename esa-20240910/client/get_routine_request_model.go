// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetRoutineRequest
	GetName() *string
}

type GetRoutineRequest struct {
	// The routine name.
	//
	// This parameter is required.
	//
	// example:
	//
	// GetRoutine
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetRoutineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineRequest) GetName() *string {
	return s.Name
}

func (s *GetRoutineRequest) SetName(v string) *GetRoutineRequest {
	s.Name = &v
	return s
}

func (s *GetRoutineRequest) Validate() error {
	return dara.Validate(s)
}
