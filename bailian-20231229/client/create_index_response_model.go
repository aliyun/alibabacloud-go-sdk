// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIndexResponse
	GetStatusCode() *int32
	SetBody(v *CreateIndexResponseBody) *CreateIndexResponse
	GetBody() *CreateIndexResponseBody
}

type CreateIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIndexResponse) GetBody() *CreateIndexResponseBody {
	return s.Body
}

func (s *CreateIndexResponse) SetHeaders(v map[string]*string) *CreateIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexResponse) SetStatusCode(v int32) *CreateIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndexResponse) SetBody(v *CreateIndexResponseBody) *CreateIndexResponse {
	s.Body = v
	return s
}

func (s *CreateIndexResponse) Validate() error {
	return dara.Validate(s)
}
