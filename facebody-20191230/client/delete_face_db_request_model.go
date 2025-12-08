// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceDbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteFaceDbRequest
	GetName() *string
}

type DeleteFaceDbRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteFaceDbRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceDbRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbRequest) GetName() *string {
	return s.Name
}

func (s *DeleteFaceDbRequest) SetName(v string) *DeleteFaceDbRequest {
	s.Name = &v
	return s
}

func (s *DeleteFaceDbRequest) Validate() error {
	return dara.Validate(s)
}
