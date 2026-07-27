// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityFollowerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFollowerId(v int64) *DeleteQualityFollowerRequest
	GetFollowerId() *int64
	SetProjectId(v int64) *DeleteQualityFollowerRequest
	GetProjectId() *int64
	SetProjectName(v string) *DeleteQualityFollowerRequest
	GetProjectName() *string
}

type DeleteQualityFollowerRequest struct {
	// The ID of the follower. You can call the [GetQualityFollower](https://help.aliyun.com/document_detail/174000.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	FollowerId *int64 `json:"FollowerId,omitempty" xml:"FollowerId,omitempty"`
	// The ID of the DataWorks workspace. You can log in to the DataWorks console and go to the Workspace Management page to obtain the workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the engine or data source where the partition expression is located. You can log in to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Data Quality page to obtain the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// autotest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DeleteQualityFollowerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityFollowerRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityFollowerRequest) GetFollowerId() *int64 {
	return s.FollowerId
}

func (s *DeleteQualityFollowerRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteQualityFollowerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteQualityFollowerRequest) SetFollowerId(v int64) *DeleteQualityFollowerRequest {
	s.FollowerId = &v
	return s
}

func (s *DeleteQualityFollowerRequest) SetProjectId(v int64) *DeleteQualityFollowerRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteQualityFollowerRequest) SetProjectName(v string) *DeleteQualityFollowerRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteQualityFollowerRequest) Validate() error {
	return dara.Validate(s)
}
