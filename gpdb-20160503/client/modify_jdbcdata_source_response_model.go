// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJDBCDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyJDBCDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyJDBCDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyJDBCDataSourceResponseBody) *ModifyJDBCDataSourceResponse
	GetBody() *ModifyJDBCDataSourceResponseBody
}

type ModifyJDBCDataSourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyJDBCDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyJDBCDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyJDBCDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyJDBCDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyJDBCDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyJDBCDataSourceResponse) GetBody() *ModifyJDBCDataSourceResponseBody {
	return s.Body
}

func (s *ModifyJDBCDataSourceResponse) SetHeaders(v map[string]*string) *ModifyJDBCDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyJDBCDataSourceResponse) SetStatusCode(v int32) *ModifyJDBCDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyJDBCDataSourceResponse) SetBody(v *ModifyJDBCDataSourceResponseBody) *ModifyJDBCDataSourceResponse {
	s.Body = v
	return s
}

func (s *ModifyJDBCDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
