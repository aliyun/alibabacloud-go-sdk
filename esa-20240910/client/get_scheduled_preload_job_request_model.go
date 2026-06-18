// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledPreloadJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetScheduledPreloadJobRequest
	GetId() *string
}

type GetScheduledPreloadJobRequest struct {
	// The task ID of the scheduled preload job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 665d3af3621bccf3fe29e1a4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetScheduledPreloadJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledPreloadJobRequest) GetId() *string {
	return s.Id
}

func (s *GetScheduledPreloadJobRequest) SetId(v string) *GetScheduledPreloadJobRequest {
	s.Id = &v
	return s
}

func (s *GetScheduledPreloadJobRequest) Validate() error {
	return dara.Validate(s)
}
