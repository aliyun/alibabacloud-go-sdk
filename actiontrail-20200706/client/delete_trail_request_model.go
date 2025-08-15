// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteTrailRequest
	GetName() *string
}

type DeleteTrailRequest struct {
	// The name of the trail that you want to delete.
	//
	// The name must be 6 to 36 characters in length. The name must start with a lowercase letter and can contain lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// > The name must be unique within your Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteTrailRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrailRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrailRequest) GetName() *string {
	return s.Name
}

func (s *DeleteTrailRequest) SetName(v string) *DeleteTrailRequest {
	s.Name = &v
	return s
}

func (s *DeleteTrailRequest) Validate() error {
	return dara.Validate(s)
}
