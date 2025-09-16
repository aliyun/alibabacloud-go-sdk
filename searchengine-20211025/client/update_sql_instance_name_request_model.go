// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateSqlInstanceNameRequest
	GetName() *string
}

type UpdateSqlInstanceNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateSqlInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceNameRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSqlInstanceNameRequest) SetName(v string) *UpdateSqlInstanceNameRequest {
	s.Name = &v
	return s
}

func (s *UpdateSqlInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
