// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdInsertionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetAdInsertionRequest
	GetName() *string
}

type GetAdInsertionRequest struct {
	// The name of the ad insertion configuration that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_ad
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAdInsertionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdInsertionRequest) GoString() string {
	return s.String()
}

func (s *GetAdInsertionRequest) GetName() *string {
	return s.Name
}

func (s *GetAdInsertionRequest) SetName(v string) *GetAdInsertionRequest {
	s.Name = &v
	return s
}

func (s *GetAdInsertionRequest) Validate() error {
	return dara.Validate(s)
}
