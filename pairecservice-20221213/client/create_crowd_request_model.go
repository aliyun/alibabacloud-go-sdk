// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrowdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateCrowdRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCrowdRequest
	GetInstanceId() *string
	SetLabel(v string) *CreateCrowdRequest
	GetLabel() *string
	SetName(v string) *CreateCrowdRequest
	GetName() *string
	SetSource(v string) *CreateCrowdRequest
	GetSource() *string
	SetUsers(v string) *CreateCrowdRequest
	GetUsers() *string
}

type CreateCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// os=android
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xx人群
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// user1,user2,user3
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s CreateCrowdRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCrowdRequest) GoString() string {
	return s.String()
}

func (s *CreateCrowdRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCrowdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCrowdRequest) GetLabel() *string {
	return s.Label
}

func (s *CreateCrowdRequest) GetName() *string {
	return s.Name
}

func (s *CreateCrowdRequest) GetSource() *string {
	return s.Source
}

func (s *CreateCrowdRequest) GetUsers() *string {
	return s.Users
}

func (s *CreateCrowdRequest) SetDescription(v string) *CreateCrowdRequest {
	s.Description = &v
	return s
}

func (s *CreateCrowdRequest) SetInstanceId(v string) *CreateCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCrowdRequest) SetLabel(v string) *CreateCrowdRequest {
	s.Label = &v
	return s
}

func (s *CreateCrowdRequest) SetName(v string) *CreateCrowdRequest {
	s.Name = &v
	return s
}

func (s *CreateCrowdRequest) SetSource(v string) *CreateCrowdRequest {
	s.Source = &v
	return s
}

func (s *CreateCrowdRequest) SetUsers(v string) *CreateCrowdRequest {
	s.Users = &v
	return s
}

func (s *CreateCrowdRequest) Validate() error {
	return dara.Validate(s)
}
