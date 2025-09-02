// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceFileResponseBody) *CreateResourceFileResponse
	GetBody() *CreateResourceFileResponseBody
}

type CreateResourceFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceFileResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceFileResponse) GetBody() *CreateResourceFileResponseBody {
	return s.Body
}

func (s *CreateResourceFileResponse) SetHeaders(v map[string]*string) *CreateResourceFileResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceFileResponse) SetStatusCode(v int32) *CreateResourceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceFileResponse) SetBody(v *CreateResourceFileResponseBody) *CreateResourceFileResponse {
	s.Body = v
	return s
}

func (s *CreateResourceFileResponse) Validate() error {
	return dara.Validate(s)
}
