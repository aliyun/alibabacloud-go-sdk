// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaceDbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateFaceDbRequest
	GetName() *string
}

type CreateFaceDbRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateFaceDbRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFaceDbRequest) GoString() string {
	return s.String()
}

func (s *CreateFaceDbRequest) GetName() *string {
	return s.Name
}

func (s *CreateFaceDbRequest) SetName(v string) *CreateFaceDbRequest {
	s.Name = &v
	return s
}

func (s *CreateFaceDbRequest) Validate() error {
	return dara.Validate(s)
}
