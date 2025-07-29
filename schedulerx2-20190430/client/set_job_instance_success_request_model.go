// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetJobInstanceSuccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *SetJobInstanceSuccessRequest
	GetGroupId() *string
	SetJobId(v int64) *SetJobInstanceSuccessRequest
	GetJobId() *int64
	SetJobInstanceId(v int64) *SetJobInstanceSuccessRequest
	GetJobInstanceId() *int64
	SetNamespace(v string) *SetJobInstanceSuccessRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *SetJobInstanceSuccessRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *SetJobInstanceSuccessRequest
	GetRegionId() *string
}

type SetJobInstanceSuccessRequest struct {
	// The application group ID. You can obtain the application group ID on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	JobInstanceId *int64 `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The source of the namespace. This parameter is required only for a special third party.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetJobInstanceSuccessRequest) String() string {
	return dara.Prettify(s)
}

func (s SetJobInstanceSuccessRequest) GoString() string {
	return s.String()
}

func (s *SetJobInstanceSuccessRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SetJobInstanceSuccessRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *SetJobInstanceSuccessRequest) GetJobInstanceId() *int64 {
	return s.JobInstanceId
}

func (s *SetJobInstanceSuccessRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SetJobInstanceSuccessRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *SetJobInstanceSuccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetJobInstanceSuccessRequest) SetGroupId(v string) *SetJobInstanceSuccessRequest {
	s.GroupId = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetJobId(v int64) *SetJobInstanceSuccessRequest {
	s.JobId = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetJobInstanceId(v int64) *SetJobInstanceSuccessRequest {
	s.JobInstanceId = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetNamespace(v string) *SetJobInstanceSuccessRequest {
	s.Namespace = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetNamespaceSource(v string) *SetJobInstanceSuccessRequest {
	s.NamespaceSource = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) SetRegionId(v string) *SetJobInstanceSuccessRequest {
	s.RegionId = &v
	return s
}

func (s *SetJobInstanceSuccessRequest) Validate() error {
	return dara.Validate(s)
}
