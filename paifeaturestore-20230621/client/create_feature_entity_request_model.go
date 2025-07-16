// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFeatureEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJoinId(v string) *CreateFeatureEntityRequest
	GetJoinId() *string
	SetName(v string) *CreateFeatureEntityRequest
	GetName() *string
	SetProjectId(v string) *CreateFeatureEntityRequest
	GetProjectId() *string
}

type CreateFeatureEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateFeatureEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFeatureEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateFeatureEntityRequest) GetJoinId() *string {
	return s.JoinId
}

func (s *CreateFeatureEntityRequest) GetName() *string {
	return s.Name
}

func (s *CreateFeatureEntityRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateFeatureEntityRequest) SetJoinId(v string) *CreateFeatureEntityRequest {
	s.JoinId = &v
	return s
}

func (s *CreateFeatureEntityRequest) SetName(v string) *CreateFeatureEntityRequest {
	s.Name = &v
	return s
}

func (s *CreateFeatureEntityRequest) SetProjectId(v string) *CreateFeatureEntityRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFeatureEntityRequest) Validate() error {
	return dara.Validate(s)
}
