// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdInsertionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteAdInsertionRequest
	GetName() *string
}

type DeleteAdInsertionRequest struct {
	// The name of the configuration that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_ad
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteAdInsertionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdInsertionRequest) GoString() string {
	return s.String()
}

func (s *DeleteAdInsertionRequest) GetName() *string {
	return s.Name
}

func (s *DeleteAdInsertionRequest) SetName(v string) *DeleteAdInsertionRequest {
	s.Name = &v
	return s
}

func (s *DeleteAdInsertionRequest) Validate() error {
	return dara.Validate(s)
}
