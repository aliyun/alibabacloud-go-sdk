// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceType(v string) *ListOperationsRequest
	GetServiceType() *string
}

type ListOperationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ListOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationsRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListOperationsRequest) SetServiceType(v string) *ListOperationsRequest {
	s.ServiceType = &v
	return s
}

func (s *ListOperationsRequest) Validate() error {
	return dara.Validate(s)
}
