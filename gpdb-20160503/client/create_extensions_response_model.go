// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *CreateExtensionsResponseBody) *CreateExtensionsResponse
	GetBody() *CreateExtensionsResponseBody
}

type CreateExtensionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExtensionsResponse) GoString() string {
	return s.String()
}

func (s *CreateExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExtensionsResponse) GetBody() *CreateExtensionsResponseBody {
	return s.Body
}

func (s *CreateExtensionsResponse) SetHeaders(v map[string]*string) *CreateExtensionsResponse {
	s.Headers = v
	return s
}

func (s *CreateExtensionsResponse) SetStatusCode(v int32) *CreateExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExtensionsResponse) SetBody(v *CreateExtensionsResponseBody) *CreateExtensionsResponse {
	s.Body = v
	return s
}

func (s *CreateExtensionsResponse) Validate() error {
	return dara.Validate(s)
}
