// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetScheduledPreloadJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ResetScheduledPreloadJobRequest
	GetId() *string
}

type ResetScheduledPreloadJobRequest struct {
	// The scheduled preload job ID.
	//
	// 	Notice: The scheduled preload job ID. This parameter is required. You can obtain the ID from the response of CreateScheduledPreloadJob after creating a job, or query existing job IDs by calling GetScheduledPreloadJob or ListScheduledPreloadJobs.
	//
	// example:
	//
	// 665d3af3621bccf3fe29e1a4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ResetScheduledPreloadJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *ResetScheduledPreloadJobRequest) GetId() *string {
	return s.Id
}

func (s *ResetScheduledPreloadJobRequest) SetId(v string) *ResetScheduledPreloadJobRequest {
	s.Id = &v
	return s
}

func (s *ResetScheduledPreloadJobRequest) Validate() error {
	return dara.Validate(s)
}
