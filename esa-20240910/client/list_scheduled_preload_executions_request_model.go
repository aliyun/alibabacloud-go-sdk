// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPreloadExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListScheduledPreloadExecutionsRequest
	GetId() *string
}

type ListScheduledPreloadExecutionsRequest struct {
	// The ID of the scheduled prefetch task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ListScheduledPreloadExecutions
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListScheduledPreloadExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsRequest) GetId() *string {
	return s.Id
}

func (s *ListScheduledPreloadExecutionsRequest) SetId(v string) *ListScheduledPreloadExecutionsRequest {
	s.Id = &v
	return s
}

func (s *ListScheduledPreloadExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
