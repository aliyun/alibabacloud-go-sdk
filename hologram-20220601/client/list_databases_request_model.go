// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternal(v bool) *ListDatabasesRequest
	GetExternal() *bool
}

type ListDatabasesRequest struct {
	// example:
	//
	// false
	External *bool `json:"external,omitempty" xml:"external,omitempty"`
}

func (s ListDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) GetExternal() *bool {
	return s.External
}

func (s *ListDatabasesRequest) SetExternal(v bool) *ListDatabasesRequest {
	s.External = &v
	return s
}

func (s *ListDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
