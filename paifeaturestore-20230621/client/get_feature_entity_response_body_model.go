// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *GetFeatureEntityResponseBody
	GetGmtCreateTime() *string
	SetJoinId(v string) *GetFeatureEntityResponseBody
	GetJoinId() *string
	SetName(v string) *GetFeatureEntityResponseBody
	GetName() *string
	SetOwner(v string) *GetFeatureEntityResponseBody
	GetOwner() *string
	SetProjectId(v string) *GetFeatureEntityResponseBody
	GetProjectId() *string
	SetProjectName(v string) *GetFeatureEntityResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetFeatureEntityResponseBody
	GetRequestId() *string
}

type GetFeatureEntityResponseBody struct {
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
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
	// 123456789*****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// E23EFF09-58AA-5420-934F-8453AE01548D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFeatureEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureEntityResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetFeatureEntityResponseBody) GetJoinId() *string {
	return s.JoinId
}

func (s *GetFeatureEntityResponseBody) GetName() *string {
	return s.Name
}

func (s *GetFeatureEntityResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetFeatureEntityResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetFeatureEntityResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetFeatureEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFeatureEntityResponseBody) SetGmtCreateTime(v string) *GetFeatureEntityResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetJoinId(v string) *GetFeatureEntityResponseBody {
	s.JoinId = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetName(v string) *GetFeatureEntityResponseBody {
	s.Name = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetOwner(v string) *GetFeatureEntityResponseBody {
	s.Owner = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetProjectId(v string) *GetFeatureEntityResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetProjectName(v string) *GetFeatureEntityResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetFeatureEntityResponseBody) SetRequestId(v string) *GetFeatureEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
