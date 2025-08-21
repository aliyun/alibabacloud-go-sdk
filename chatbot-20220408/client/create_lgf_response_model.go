// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLgfResponse
	GetStatusCode() *int32
	SetBody(v *CreateLgfResponseBody) *CreateLgfResponse
	GetBody() *CreateLgfResponseBody
}

type CreateLgfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLgfResponse) GoString() string {
	return s.String()
}

func (s *CreateLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLgfResponse) GetBody() *CreateLgfResponseBody {
	return s.Body
}

func (s *CreateLgfResponse) SetHeaders(v map[string]*string) *CreateLgfResponse {
	s.Headers = v
	return s
}

func (s *CreateLgfResponse) SetStatusCode(v int32) *CreateLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLgfResponse) SetBody(v *CreateLgfResponseBody) *CreateLgfResponse {
	s.Body = v
	return s
}

func (s *CreateLgfResponse) Validate() error {
	return dara.Validate(s)
}
