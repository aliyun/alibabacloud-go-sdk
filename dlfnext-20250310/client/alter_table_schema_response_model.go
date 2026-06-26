// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterTableSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterTableSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterTableSchemaResponse
	GetStatusCode() *int32
}

type AlterTableSchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AlterTableSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterTableSchemaResponse) GoString() string {
	return s.String()
}

func (s *AlterTableSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterTableSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterTableSchemaResponse) SetHeaders(v map[string]*string) *AlterTableSchemaResponse {
	s.Headers = v
	return s
}

func (s *AlterTableSchemaResponse) SetStatusCode(v int32) *AlterTableSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterTableSchemaResponse) Validate() error {
	return dara.Validate(s)
}
