// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpApiResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpApiResponseBody) *CreateHttpApiResponse
	GetBody() *CreateHttpApiResponseBody
}

type CreateHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpApiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpApiResponse) GetBody() *CreateHttpApiResponseBody {
	return s.Body
}

func (s *CreateHttpApiResponse) SetHeaders(v map[string]*string) *CreateHttpApiResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpApiResponse) SetStatusCode(v int32) *CreateHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpApiResponse) SetBody(v *CreateHttpApiResponseBody) *CreateHttpApiResponse {
	s.Body = v
	return s
}

func (s *CreateHttpApiResponse) Validate() error {
	return dara.Validate(s)
}
