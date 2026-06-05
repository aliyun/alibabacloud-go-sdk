// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHiveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateHiveRequest
	GetDescription() *string
	SetName(v string) *CreateHiveRequest
	GetName() *string
}

type CreateHiveRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateHiveRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHiveRequest) GoString() string {
	return s.String()
}

func (s *CreateHiveRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHiveRequest) GetName() *string {
	return s.Name
}

func (s *CreateHiveRequest) SetDescription(v string) *CreateHiveRequest {
	s.Description = &v
	return s
}

func (s *CreateHiveRequest) SetName(v string) *CreateHiveRequest {
	s.Name = &v
	return s
}

func (s *CreateHiveRequest) Validate() error {
	return dara.Validate(s)
}
