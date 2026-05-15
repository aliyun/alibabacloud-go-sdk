// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetQualityProjectLogRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetQualityProjectLogRequest
	GetProjectId() *int64
}

type GetQualityProjectLogRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProjectId  *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetQualityProjectLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectLogRequest) GoString() string {
	return s.String()
}

func (s *GetQualityProjectLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQualityProjectLogRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityProjectLogRequest) SetInstanceId(v string) *GetQualityProjectLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityProjectLogRequest) SetProjectId(v int64) *GetQualityProjectLogRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQualityProjectLogRequest) Validate() error {
	return dara.Validate(s)
}
