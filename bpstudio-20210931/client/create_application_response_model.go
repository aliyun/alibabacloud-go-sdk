// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse
	GetBody() *CreateApplicationResponseBody
}

type CreateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationResponse) GetBody() *CreateApplicationResponseBody {
	return s.Body
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetStatusCode(v int32) *CreateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationResponse) Validate() error {
	return dara.Validate(s)
}
