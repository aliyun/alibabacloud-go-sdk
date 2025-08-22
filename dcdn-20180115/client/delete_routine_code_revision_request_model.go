// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineCodeRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteRoutineCodeRevisionRequest
	GetName() *string
	SetSelectCodeRevision(v string) *DeleteRoutineCodeRevisionRequest
	GetSelectCodeRevision() *string
}

type DeleteRoutineCodeRevisionRequest struct {
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the version that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	SelectCodeRevision *string `json:"SelectCodeRevision,omitempty" xml:"SelectCodeRevision,omitempty"`
}

func (s DeleteRoutineCodeRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineCodeRevisionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeRevisionRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineCodeRevisionRequest) GetSelectCodeRevision() *string {
	return s.SelectCodeRevision
}

func (s *DeleteRoutineCodeRevisionRequest) SetName(v string) *DeleteRoutineCodeRevisionRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineCodeRevisionRequest) SetSelectCodeRevision(v string) *DeleteRoutineCodeRevisionRequest {
	s.SelectCodeRevision = &v
	return s
}

func (s *DeleteRoutineCodeRevisionRequest) Validate() error {
	return dara.Validate(s)
}
