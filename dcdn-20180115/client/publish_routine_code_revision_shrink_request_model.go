// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRoutineCodeRevisionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvsShrink(v string) *PublishRoutineCodeRevisionShrinkRequest
	GetEnvsShrink() *string
	SetName(v string) *PublishRoutineCodeRevisionShrinkRequest
	GetName() *string
	SetSelectCodeRevision(v string) *PublishRoutineCodeRevisionShrinkRequest
	GetSelectCodeRevision() *string
}

type PublishRoutineCodeRevisionShrinkRequest struct {
	// The environment to which you want to publish the code.
	//
	// >
	//
	// 	- production: the name of the environment, including the environment name (SpecName) and the domain name whitelist (AllowedHosts).
	//
	// 	- presetCanary: You can add canary release environments based on your business requirements. This parameter is optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["production","presetCanaryZhejiang"]
	EnvsShrink *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The name of the routine. The name must be unique among the routines that belong to the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the routine code that you want to publish.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1620876959997924701
	SelectCodeRevision *string `json:"SelectCodeRevision,omitempty" xml:"SelectCodeRevision,omitempty"`
}

func (s PublishRoutineCodeRevisionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishRoutineCodeRevisionShrinkRequest) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionShrinkRequest) GetEnvsShrink() *string {
	return s.EnvsShrink
}

func (s *PublishRoutineCodeRevisionShrinkRequest) GetName() *string {
	return s.Name
}

func (s *PublishRoutineCodeRevisionShrinkRequest) GetSelectCodeRevision() *string {
	return s.SelectCodeRevision
}

func (s *PublishRoutineCodeRevisionShrinkRequest) SetEnvsShrink(v string) *PublishRoutineCodeRevisionShrinkRequest {
	s.EnvsShrink = &v
	return s
}

func (s *PublishRoutineCodeRevisionShrinkRequest) SetName(v string) *PublishRoutineCodeRevisionShrinkRequest {
	s.Name = &v
	return s
}

func (s *PublishRoutineCodeRevisionShrinkRequest) SetSelectCodeRevision(v string) *PublishRoutineCodeRevisionShrinkRequest {
	s.SelectCodeRevision = &v
	return s
}

func (s *PublishRoutineCodeRevisionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
