// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteQualityProjectRequest
	GetInstanceId() *string
	SetProjectId(v int64) *DeleteQualityProjectRequest
	GetProjectId() *int64
}

type DeleteQualityProjectRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteQualityProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteQualityProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteQualityProjectRequest) SetInstanceId(v string) *DeleteQualityProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQualityProjectRequest) SetProjectId(v int64) *DeleteQualityProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteQualityProjectRequest) Validate() error {
	return dara.Validate(s)
}
