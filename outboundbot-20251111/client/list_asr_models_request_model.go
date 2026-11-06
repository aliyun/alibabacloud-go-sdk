// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsrModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAsrModelsRequest
	GetInstanceId() *string
}

type ListAsrModelsRequest struct {
	// The instance ID.
	//
	// example:
	//
	// fc21ccd7-6f0a-4261-ac63-d079bfe3bc2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAsrModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsrModelsRequest) GoString() string {
	return s.String()
}

func (s *ListAsrModelsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAsrModelsRequest) SetInstanceId(v string) *ListAsrModelsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAsrModelsRequest) Validate() error {
	return dara.Validate(s)
}
