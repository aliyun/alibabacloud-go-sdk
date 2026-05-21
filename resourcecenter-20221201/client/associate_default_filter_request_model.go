// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateDefaultFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterName(v string) *AssociateDefaultFilterRequest
	GetFilterName() *string
}

type AssociateDefaultFilterRequest struct {
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// My Filters
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s AssociateDefaultFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateDefaultFilterRequest) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterRequest) GetFilterName() *string {
	return s.FilterName
}

func (s *AssociateDefaultFilterRequest) SetFilterName(v string) *AssociateDefaultFilterRequest {
	s.FilterName = &v
	return s
}

func (s *AssociateDefaultFilterRequest) Validate() error {
	return dara.Validate(s)
}
