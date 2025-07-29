// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterChecksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTarget(v string) *ListClusterChecksRequest
	GetTarget() *string
	SetType(v string) *ListClusterChecksRequest
	GetType() *string
}

type ListClusterChecksRequest struct {
	// The targets to check.
	//
	// example:
	//
	// ngw-bp19ay6nnvd4cexxxx
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// The check method.
	//
	// example:
	//
	// ClusterUpgrade
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListClusterChecksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterChecksRequest) GoString() string {
	return s.String()
}

func (s *ListClusterChecksRequest) GetTarget() *string {
	return s.Target
}

func (s *ListClusterChecksRequest) GetType() *string {
	return s.Type
}

func (s *ListClusterChecksRequest) SetTarget(v string) *ListClusterChecksRequest {
	s.Target = &v
	return s
}

func (s *ListClusterChecksRequest) SetType(v string) *ListClusterChecksRequest {
	s.Type = &v
	return s
}

func (s *ListClusterChecksRequest) Validate() error {
	return dara.Validate(s)
}
