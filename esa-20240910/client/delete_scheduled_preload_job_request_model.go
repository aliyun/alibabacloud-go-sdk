// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPreloadJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteScheduledPreloadJobRequest
	GetId() *string
}

type DeleteScheduledPreloadJobRequest struct {
	// The ID of the scheduled prefetch task.
	//
	// This parameter is required.
	//
	// example:
	//
	// DeleteScheduledPreloadJob
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteScheduledPreloadJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPreloadJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadJobRequest) GetId() *string {
	return s.Id
}

func (s *DeleteScheduledPreloadJobRequest) SetId(v string) *DeleteScheduledPreloadJobRequest {
	s.Id = &v
	return s
}

func (s *DeleteScheduledPreloadJobRequest) Validate() error {
	return dara.Validate(s)
}
