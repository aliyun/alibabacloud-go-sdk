// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportConfigJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetExportConfigJobRequest
	GetInstanceId() *string
	SetJobId(v string) *GetExportConfigJobRequest
	GetJobId() *string
	SetRegionId(v string) *GetExportConfigJobRequest
	GetRegionId() *string
}

type GetExportConfigJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetExportConfigJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExportConfigJobRequest) GoString() string {
	return s.String()
}

func (s *GetExportConfigJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetExportConfigJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetExportConfigJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetExportConfigJobRequest) SetInstanceId(v string) *GetExportConfigJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetExportConfigJobRequest) SetJobId(v string) *GetExportConfigJobRequest {
	s.JobId = &v
	return s
}

func (s *GetExportConfigJobRequest) SetRegionId(v string) *GetExportConfigJobRequest {
	s.RegionId = &v
	return s
}

func (s *GetExportConfigJobRequest) Validate() error {
	return dara.Validate(s)
}
