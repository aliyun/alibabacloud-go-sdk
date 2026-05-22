// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineCodeVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *GetRoutineCodeVersionRequest
	GetCodeVersion() *string
	SetName(v string) *GetRoutineCodeVersionRequest
	GetName() *string
}

type GetRoutineCodeVersionRequest struct {
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
	// GetRoutineCodeVersion
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetRoutineCodeVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineCodeVersionRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineCodeVersionRequest) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *GetRoutineCodeVersionRequest) GetName() *string {
	return s.Name
}

func (s *GetRoutineCodeVersionRequest) SetCodeVersion(v string) *GetRoutineCodeVersionRequest {
	s.CodeVersion = &v
	return s
}

func (s *GetRoutineCodeVersionRequest) SetName(v string) *GetRoutineCodeVersionRequest {
	s.Name = &v
	return s
}

func (s *GetRoutineCodeVersionRequest) Validate() error {
	return dara.Validate(s)
}
