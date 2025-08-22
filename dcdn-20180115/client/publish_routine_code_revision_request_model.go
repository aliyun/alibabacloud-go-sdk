// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRoutineCodeRevisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvs(v map[string]interface{}) *PublishRoutineCodeRevisionRequest
	GetEnvs() map[string]interface{}
	SetName(v string) *PublishRoutineCodeRevisionRequest
	GetName() *string
	SetSelectCodeRevision(v string) *PublishRoutineCodeRevisionRequest
	GetSelectCodeRevision() *string
}

type PublishRoutineCodeRevisionRequest struct {
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
	Envs map[string]interface{} `json:"Envs,omitempty" xml:"Envs,omitempty"`
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

func (s PublishRoutineCodeRevisionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishRoutineCodeRevisionRequest) GoString() string {
	return s.String()
}

func (s *PublishRoutineCodeRevisionRequest) GetEnvs() map[string]interface{} {
	return s.Envs
}

func (s *PublishRoutineCodeRevisionRequest) GetName() *string {
	return s.Name
}

func (s *PublishRoutineCodeRevisionRequest) GetSelectCodeRevision() *string {
	return s.SelectCodeRevision
}

func (s *PublishRoutineCodeRevisionRequest) SetEnvs(v map[string]interface{}) *PublishRoutineCodeRevisionRequest {
	s.Envs = v
	return s
}

func (s *PublishRoutineCodeRevisionRequest) SetName(v string) *PublishRoutineCodeRevisionRequest {
	s.Name = &v
	return s
}

func (s *PublishRoutineCodeRevisionRequest) SetSelectCodeRevision(v string) *PublishRoutineCodeRevisionRequest {
	s.SelectCodeRevision = &v
	return s
}

func (s *PublishRoutineCodeRevisionRequest) Validate() error {
	return dara.Validate(s)
}
