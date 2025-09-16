// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSqlInstanceRequest
	GetName() *string
	SetParent(v int64) *CreateSqlInstanceRequest
	GetParent() *int64
}

type CreateSqlInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	Parent *int64 `json:"parent,omitempty" xml:"parent,omitempty"`
}

func (s CreateSqlInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSqlInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateSqlInstanceRequest) GetParent() *int64 {
	return s.Parent
}

func (s *CreateSqlInstanceRequest) SetName(v string) *CreateSqlInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateSqlInstanceRequest) SetParent(v int64) *CreateSqlInstanceRequest {
	s.Parent = &v
	return s
}

func (s *CreateSqlInstanceRequest) Validate() error {
	return dara.Validate(s)
}
