// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtended(v bool) *ListAllNodeRequest
	GetExtended() *bool
}

type ListAllNodeRequest struct {
	// The Java Virtual Machine (JVM) heap memory usage of the node.
	//
	// example:
	//
	// false
	Extended *bool `json:"extended,omitempty" xml:"extended,omitempty"`
}

func (s ListAllNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllNodeRequest) GoString() string {
	return s.String()
}

func (s *ListAllNodeRequest) GetExtended() *bool {
	return s.Extended
}

func (s *ListAllNodeRequest) SetExtended(v bool) *ListAllNodeRequest {
	s.Extended = &v
	return s
}

func (s *ListAllNodeRequest) Validate() error {
	return dara.Validate(s)
}
