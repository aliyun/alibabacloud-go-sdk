// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTableLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTableLevelResponse
	GetStatusCode() *int32
	SetBody(v *CreateTableLevelResponseBody) *CreateTableLevelResponse
	GetBody() *CreateTableLevelResponseBody
}

type CreateTableLevelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTableLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTableLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTableLevelResponse) GoString() string {
	return s.String()
}

func (s *CreateTableLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTableLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTableLevelResponse) GetBody() *CreateTableLevelResponseBody {
	return s.Body
}

func (s *CreateTableLevelResponse) SetHeaders(v map[string]*string) *CreateTableLevelResponse {
	s.Headers = v
	return s
}

func (s *CreateTableLevelResponse) SetStatusCode(v int32) *CreateTableLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTableLevelResponse) SetBody(v *CreateTableLevelResponseBody) *CreateTableLevelResponse {
	s.Body = v
	return s
}

func (s *CreateTableLevelResponse) Validate() error {
	return dara.Validate(s)
}
