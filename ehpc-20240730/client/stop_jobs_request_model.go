// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *StopJobsRequest
	GetClusterId() *string
	SetJobIds(v []*string) *StopJobsRequest
	GetJobIds() []*string
}

type StopJobsRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of the jobs that you want to stop.
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s StopJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s StopJobsRequest) GoString() string {
	return s.String()
}

func (s *StopJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *StopJobsRequest) GetJobIds() []*string {
	return s.JobIds
}

func (s *StopJobsRequest) SetClusterId(v string) *StopJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *StopJobsRequest) SetJobIds(v []*string) *StopJobsRequest {
	s.JobIds = v
	return s
}

func (s *StopJobsRequest) Validate() error {
	return dara.Validate(s)
}
