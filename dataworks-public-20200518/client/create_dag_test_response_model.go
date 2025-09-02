// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDagTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDagTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDagTestResponse
	GetStatusCode() *int32
	SetBody(v *CreateDagTestResponseBody) *CreateDagTestResponse
	GetBody() *CreateDagTestResponseBody
}

type CreateDagTestResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDagTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDagTestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDagTestResponse) GoString() string {
	return s.String()
}

func (s *CreateDagTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDagTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDagTestResponse) GetBody() *CreateDagTestResponseBody {
	return s.Body
}

func (s *CreateDagTestResponse) SetHeaders(v map[string]*string) *CreateDagTestResponse {
	s.Headers = v
	return s
}

func (s *CreateDagTestResponse) SetStatusCode(v int32) *CreateDagTestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDagTestResponse) SetBody(v *CreateDagTestResponseBody) *CreateDagTestResponse {
	s.Body = v
	return s
}

func (s *CreateDagTestResponse) Validate() error {
	return dara.Validate(s)
}
