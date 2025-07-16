// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *CreateInstanceRequest
	GetType() *string
}

type CreateInstanceRequest struct {
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetType() *string {
	return s.Type
}

func (s *CreateInstanceRequest) SetType(v string) *CreateInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
