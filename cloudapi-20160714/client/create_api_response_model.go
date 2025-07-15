// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiResponseBody) *CreateApiResponse
	GetBody() *CreateApiResponseBody
}

type CreateApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiResponse) GoString() string {
	return s.String()
}

func (s *CreateApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiResponse) GetBody() *CreateApiResponseBody {
	return s.Body
}

func (s *CreateApiResponse) SetHeaders(v map[string]*string) *CreateApiResponse {
	s.Headers = v
	return s
}

func (s *CreateApiResponse) SetStatusCode(v int32) *CreateApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiResponse) SetBody(v *CreateApiResponseBody) *CreateApiResponse {
	s.Body = v
	return s
}

func (s *CreateApiResponse) Validate() error {
	return dara.Validate(s)
}
