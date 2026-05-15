// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeQualityProjectStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ChangeQualityProjectStatusRequest
	GetInstanceId() *string
	SetProjectId(v int64) *ChangeQualityProjectStatusRequest
	GetProjectId() *int64
	SetStatus(v int32) *ChangeQualityProjectStatusRequest
	GetStatus() *int32
}

type ChangeQualityProjectStatusRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ChangeQualityProjectStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeQualityProjectStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeQualityProjectStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeQualityProjectStatusRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ChangeQualityProjectStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeQualityProjectStatusRequest) SetInstanceId(v string) *ChangeQualityProjectStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeQualityProjectStatusRequest) SetProjectId(v int64) *ChangeQualityProjectStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *ChangeQualityProjectStatusRequest) SetStatus(v int32) *ChangeQualityProjectStatusRequest {
	s.Status = &v
	return s
}

func (s *ChangeQualityProjectStatusRequest) Validate() error {
	return dara.Validate(s)
}
