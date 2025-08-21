// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDocResponse
	GetStatusCode() *int32
	SetBody(v *CreateDocResponseBody) *CreateDocResponse
	GetBody() *CreateDocResponseBody
}

type CreateDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDocResponse) GoString() string {
	return s.String()
}

func (s *CreateDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDocResponse) GetBody() *CreateDocResponseBody {
	return s.Body
}

func (s *CreateDocResponse) SetHeaders(v map[string]*string) *CreateDocResponse {
	s.Headers = v
	return s
}

func (s *CreateDocResponse) SetStatusCode(v int32) *CreateDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocResponse) SetBody(v *CreateDocResponseBody) *CreateDocResponse {
	s.Body = v
	return s
}

func (s *CreateDocResponse) Validate() error {
	return dara.Validate(s)
}
