// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineCodeVersionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *GetRoutineCodeVersionInfoRequest
	GetCodeVersion() *string
	SetName(v string) *GetRoutineCodeVersionInfoRequest
	GetName() *string
}

type GetRoutineCodeVersionInfoRequest struct {
	// The code version number to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710120201067203242
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The name of the Edge Routine.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetRoutineCodeVersionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineCodeVersionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetRoutineCodeVersionInfoRequest) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *GetRoutineCodeVersionInfoRequest) GetName() *string {
	return s.Name
}

func (s *GetRoutineCodeVersionInfoRequest) SetCodeVersion(v string) *GetRoutineCodeVersionInfoRequest {
	s.CodeVersion = &v
	return s
}

func (s *GetRoutineCodeVersionInfoRequest) SetName(v string) *GetRoutineCodeVersionInfoRequest {
	s.Name = &v
	return s
}

func (s *GetRoutineCodeVersionInfoRequest) Validate() error {
	return dara.Validate(s)
}
