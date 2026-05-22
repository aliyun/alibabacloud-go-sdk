// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineCodeVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *DeleteRoutineCodeVersionRequest
	GetCodeVersion() *string
	SetName(v string) *DeleteRoutineCodeVersionRequest
	GetName() *string
}

type DeleteRoutineCodeVersionRequest struct {
	// The code version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710120201067203242
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The routine name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-routine1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteRoutineCodeVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineCodeVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeVersionRequest) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *DeleteRoutineCodeVersionRequest) GetName() *string {
	return s.Name
}

func (s *DeleteRoutineCodeVersionRequest) SetCodeVersion(v string) *DeleteRoutineCodeVersionRequest {
	s.CodeVersion = &v
	return s
}

func (s *DeleteRoutineCodeVersionRequest) SetName(v string) *DeleteRoutineCodeVersionRequest {
	s.Name = &v
	return s
}

func (s *DeleteRoutineCodeVersionRequest) Validate() error {
	return dara.Validate(s)
}
