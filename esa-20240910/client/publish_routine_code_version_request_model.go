// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRoutineCodeVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *PublishRoutineCodeVersionRequest
	GetCodeVersion() *string
	SetEnv(v string) *PublishRoutineCodeVersionRequest
	GetEnv() *string
	SetName(v string) *PublishRoutineCodeVersionRequest
	GetName() *string
}

type PublishRoutineCodeVersionRequest struct {
	// The version number of the code to publish.
	//
	// example:
	//
	// 1710120201067203242
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The environment name.
	//
	// This parameter is required.
	//
	// example:
	//
	// production
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The name of the Edge Routine.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-routine1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PublishRoutineCodeVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishRoutineCodeVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeVersionRequest) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *PublishRoutineCodeVersionRequest) GetEnv() *string {
	return s.Env
}

func (s *PublishRoutineCodeVersionRequest) GetName() *string {
	return s.Name
}

func (s *PublishRoutineCodeVersionRequest) SetCodeVersion(v string) *PublishRoutineCodeVersionRequest {
	s.CodeVersion = &v
	return s
}

func (s *PublishRoutineCodeVersionRequest) SetEnv(v string) *PublishRoutineCodeVersionRequest {
	s.Env = &v
	return s
}

func (s *PublishRoutineCodeVersionRequest) SetName(v string) *PublishRoutineCodeVersionRequest {
	s.Name = &v
	return s
}

func (s *PublishRoutineCodeVersionRequest) Validate() error {
	return dara.Validate(s)
}
