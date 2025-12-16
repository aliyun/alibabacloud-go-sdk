// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsActive(v int32) *ListQueryProcessorsRequest
	GetIsActive() *int32
}

type ListQueryProcessorsRequest struct {
	// The scope of query analysis rules to be queried. Default value: 0. Valid values:
	//
	// 	- 0: queries all query analysis rules.
	//
	// 	- 1: queries the default query analysis rules.
	//
	// 	- 2: queries the query analysis rules that are not the default rules.
	//
	// example:
	//
	// 0
	IsActive *int32 `json:"isActive,omitempty" xml:"isActive,omitempty"`
}

func (s ListQueryProcessorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorsRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsRequest) GetIsActive() *int32 {
	return s.IsActive
}

func (s *ListQueryProcessorsRequest) SetIsActive(v int32) *ListQueryProcessorsRequest {
	s.IsActive = &v
	return s
}

func (s *ListQueryProcessorsRequest) Validate() error {
	return dara.Validate(s)
}
