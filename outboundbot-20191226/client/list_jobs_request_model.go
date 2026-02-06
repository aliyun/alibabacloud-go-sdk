// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListJobsRequest
	GetInstanceId() *string
	SetJobId(v []*string) *ListJobsRequest
	GetJobId() []*string
}

type ListJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24fb9a8d-f20e-4ee2-a11c-094dda68c5cc
	JobId []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListJobsRequest) GetJobId() []*string {
	return s.JobId
}

func (s *ListJobsRequest) SetInstanceId(v string) *ListJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListJobsRequest) SetJobId(v []*string) *ListJobsRequest {
	s.JobId = v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
