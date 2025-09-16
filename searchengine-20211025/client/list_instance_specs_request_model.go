// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceSpecsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *ListInstanceSpecsRequest
	GetType() *string
}

type ListInstanceSpecsRequest struct {
	// The node type. Valid values:
	//
	// 	- qrs: Query Result Searcher (QRS) Worker
	//
	// 	- search: Searcher Worker
	//
	// 	- index: index node
	//
	// 	- cluster: cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// search
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListInstanceSpecsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceSpecsRequest) GetType() *string {
	return s.Type
}

func (s *ListInstanceSpecsRequest) SetType(v string) *ListInstanceSpecsRequest {
	s.Type = &v
	return s
}

func (s *ListInstanceSpecsRequest) Validate() error {
	return dara.Validate(s)
}
