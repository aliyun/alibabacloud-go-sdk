// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CheckSqlContentRequest
	GetInstanceId() *string
	SetJobName(v string) *CheckSqlContentRequest
	GetJobName() *string
	SetRegionId(v string) *CheckSqlContentRequest
	GetRegionId() *string
	SetSqlContent(v string) *CheckSqlContentRequest
	GetSqlContent() *string
}

type CheckSqlContentRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SqlContent *string `json:"SqlContent,omitempty" xml:"SqlContent,omitempty"`
}

func (s CheckSqlContentRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlContentRequest) GoString() string {
	return s.String()
}

func (s *CheckSqlContentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckSqlContentRequest) GetJobName() *string {
	return s.JobName
}

func (s *CheckSqlContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckSqlContentRequest) GetSqlContent() *string {
	return s.SqlContent
}

func (s *CheckSqlContentRequest) SetInstanceId(v string) *CheckSqlContentRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckSqlContentRequest) SetJobName(v string) *CheckSqlContentRequest {
	s.JobName = &v
	return s
}

func (s *CheckSqlContentRequest) SetRegionId(v string) *CheckSqlContentRequest {
	s.RegionId = &v
	return s
}

func (s *CheckSqlContentRequest) SetSqlContent(v string) *CheckSqlContentRequest {
	s.SqlContent = &v
	return s
}

func (s *CheckSqlContentRequest) Validate() error {
	return dara.Validate(s)
}
