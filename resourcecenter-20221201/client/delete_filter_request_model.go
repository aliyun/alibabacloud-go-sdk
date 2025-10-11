// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterName(v string) *DeleteFilterRequest
	GetFilterName() *string
}

type DeleteFilterRequest struct {
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s DeleteFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilterRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilterRequest) GetFilterName() *string {
	return s.FilterName
}

func (s *DeleteFilterRequest) SetFilterName(v string) *DeleteFilterRequest {
	s.FilterName = &v
	return s
}

func (s *DeleteFilterRequest) Validate() error {
	return dara.Validate(s)
}
