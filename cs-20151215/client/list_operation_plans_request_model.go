// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListOperationPlansRequest
	GetClusterId() *string
	SetType(v string) *ListOperationPlansRequest
	GetType() *string
}

type ListOperationPlansRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// c29ced64b3dfe4f33b57ca0aa9f68****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The operation plan type.
	//
	// example:
	//
	// cluster_upgrade
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListOperationPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationPlansRequest) GoString() string {
	return s.String()
}

func (s *ListOperationPlansRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListOperationPlansRequest) GetType() *string {
	return s.Type
}

func (s *ListOperationPlansRequest) SetClusterId(v string) *ListOperationPlansRequest {
	s.ClusterId = &v
	return s
}

func (s *ListOperationPlansRequest) SetType(v string) *ListOperationPlansRequest {
	s.Type = &v
	return s
}

func (s *ListOperationPlansRequest) Validate() error {
	return dara.Validate(s)
}
