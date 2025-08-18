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
	// The ID of the scheduled prefetch task.
	//
	// example:
	//
	// ResetScheduledPreloadJob
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
