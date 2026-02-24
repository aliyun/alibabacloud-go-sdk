// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedSearchFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSql(v string) *CreateAdvancedSearchFileRequest
	GetSql() *string
}

type CreateAdvancedSearchFileRequest struct {
	// This parameter is required.
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s CreateAdvancedSearchFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedSearchFileRequest) GoString() string {
	return s.String()
}

func (s *CreateAdvancedSearchFileRequest) GetSql() *string {
	return s.Sql
}

func (s *CreateAdvancedSearchFileRequest) SetSql(v string) *CreateAdvancedSearchFileRequest {
	s.Sql = &v
	return s
}

func (s *CreateAdvancedSearchFileRequest) Validate() error {
	return dara.Validate(s)
}
