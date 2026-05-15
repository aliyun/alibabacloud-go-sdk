// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityProjectDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetQualityProjectDetailRequest
	GetInstanceId() *string
	SetProjectId(v int64) *GetQualityProjectDetailRequest
	GetProjectId() *int64
}

type GetQualityProjectDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15977801
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetQualityProjectDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQualityProjectDetailRequest) GoString() string {
	return s.String()
}

func (s *GetQualityProjectDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetQualityProjectDetailRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetQualityProjectDetailRequest) SetInstanceId(v string) *GetQualityProjectDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQualityProjectDetailRequest) SetProjectId(v int64) *GetQualityProjectDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *GetQualityProjectDetailRequest) Validate() error {
	return dara.Validate(s)
}
