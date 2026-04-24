// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageAssociatedProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListImageAssociatedProjectsRequest
	GetId() *string
}

type ListImageAssociatedProjectsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListImageAssociatedProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageAssociatedProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListImageAssociatedProjectsRequest) GetId() *string {
	return s.Id
}

func (s *ListImageAssociatedProjectsRequest) SetId(v string) *ListImageAssociatedProjectsRequest {
	s.Id = &v
	return s
}

func (s *ListImageAssociatedProjectsRequest) Validate() error {
	return dara.Validate(s)
}
