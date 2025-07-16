// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectFeatureEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureEntityId(v string) *GetProjectFeatureEntityResponseBody
	GetFeatureEntityId() *string
	SetJoinId(v string) *GetProjectFeatureEntityResponseBody
	GetJoinId() *string
	SetName(v string) *GetProjectFeatureEntityResponseBody
	GetName() *string
	SetProjectName(v string) *GetProjectFeatureEntityResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetProjectFeatureEntityResponseBody
	GetRequestId() *string
	SetWorkspaceId(v string) *GetProjectFeatureEntityResponseBody
	GetWorkspaceId() *string
}

type GetProjectFeatureEntityResponseBody struct {
	// example:
	//
	// 3
	FeatureEntityId *string `json:"FeatureEntityId,omitempty" xml:"FeatureEntityId,omitempty"`
	// example:
	//
	// user_id
	JoinId *string `json:"JoinId,omitempty" xml:"JoinId,omitempty"`
	// example:
	//
	// feature_entity_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 37D19490-AB69-567D-A852-407C94E510E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 34245
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetProjectFeatureEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectFeatureEntityResponseBody) GetFeatureEntityId() *string {
	return s.FeatureEntityId
}

func (s *GetProjectFeatureEntityResponseBody) GetJoinId() *string {
	return s.JoinId
}

func (s *GetProjectFeatureEntityResponseBody) GetName() *string {
	return s.Name
}

func (s *GetProjectFeatureEntityResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetProjectFeatureEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectFeatureEntityResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetProjectFeatureEntityResponseBody) SetFeatureEntityId(v string) *GetProjectFeatureEntityResponseBody {
	s.FeatureEntityId = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetJoinId(v string) *GetProjectFeatureEntityResponseBody {
	s.JoinId = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetName(v string) *GetProjectFeatureEntityResponseBody {
	s.Name = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetProjectName(v string) *GetProjectFeatureEntityResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetRequestId(v string) *GetProjectFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) SetWorkspaceId(v string) *GetProjectFeatureEntityResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetProjectFeatureEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
