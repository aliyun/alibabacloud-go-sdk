// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlInstanceContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateSqlInstanceContentRequest
	GetContent() *string
}

type UpdateSqlInstanceContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// select 	- from test
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s UpdateSqlInstanceContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlInstanceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateSqlInstanceContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateSqlInstanceContentRequest) SetContent(v string) *UpdateSqlInstanceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateSqlInstanceContentRequest) Validate() error {
	return dara.Validate(s)
}
