// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationPlansForRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListOperationPlansForRegionRequest
	GetClusterId() *string
	SetState(v string) *ListOperationPlansForRegionRequest
	GetState() *string
	SetType(v string) *ListOperationPlansForRegionRequest
	GetType() *string
}

type ListOperationPlansForRegionRequest struct {
	// example:
	//
	// c02b3e03be10643e8a644a843ffcb****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// Scheduled
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// CLUSTER_UPGRADE_MASTER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListOperationPlansForRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationPlansForRegionRequest) GoString() string {
	return s.String()
}

func (s *ListOperationPlansForRegionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListOperationPlansForRegionRequest) GetState() *string {
	return s.State
}

func (s *ListOperationPlansForRegionRequest) GetType() *string {
	return s.Type
}

func (s *ListOperationPlansForRegionRequest) SetClusterId(v string) *ListOperationPlansForRegionRequest {
	s.ClusterId = &v
	return s
}

func (s *ListOperationPlansForRegionRequest) SetState(v string) *ListOperationPlansForRegionRequest {
	s.State = &v
	return s
}

func (s *ListOperationPlansForRegionRequest) SetType(v string) *ListOperationPlansForRegionRequest {
	s.Type = &v
	return s
}

func (s *ListOperationPlansForRegionRequest) Validate() error {
	return dara.Validate(s)
}
