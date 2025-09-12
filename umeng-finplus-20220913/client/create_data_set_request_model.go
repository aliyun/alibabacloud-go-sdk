// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateDataSetRequest
	GetName() *string
	SetType(v string) *CreateDataSetRequest
	GetType() *string
}

type CreateDataSetRequest struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSetRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSetRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataSetRequest) GetType() *string {
	return s.Type
}

func (s *CreateDataSetRequest) SetName(v string) *CreateDataSetRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSetRequest) SetType(v string) *CreateDataSetRequest {
	s.Type = &v
	return s
}

func (s *CreateDataSetRequest) Validate() error {
	return dara.Validate(s)
}
