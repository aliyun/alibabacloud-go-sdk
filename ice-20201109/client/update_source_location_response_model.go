// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSourceLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSourceLocationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSourceLocationResponseBody) *UpdateSourceLocationResponse
	GetBody() *UpdateSourceLocationResponseBody
}

type UpdateSourceLocationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSourceLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSourceLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceLocationResponse) GoString() string {
	return s.String()
}

func (s *UpdateSourceLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSourceLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSourceLocationResponse) GetBody() *UpdateSourceLocationResponseBody {
	return s.Body
}

func (s *UpdateSourceLocationResponse) SetHeaders(v map[string]*string) *UpdateSourceLocationResponse {
	s.Headers = v
	return s
}

func (s *UpdateSourceLocationResponse) SetStatusCode(v int32) *UpdateSourceLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSourceLocationResponse) SetBody(v *UpdateSourceLocationResponseBody) *UpdateSourceLocationResponse {
	s.Body = v
	return s
}

func (s *UpdateSourceLocationResponse) Validate() error {
	return dara.Validate(s)
}
