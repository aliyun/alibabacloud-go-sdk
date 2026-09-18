// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalPublicUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePersonalPublicUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePersonalPublicUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreatePersonalPublicUrlResponseBody) *CreatePersonalPublicUrlResponse
	GetBody() *CreatePersonalPublicUrlResponseBody
}

type CreatePersonalPublicUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePersonalPublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePersonalPublicUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalPublicUrlResponse) GoString() string {
	return s.String()
}

func (s *CreatePersonalPublicUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePersonalPublicUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePersonalPublicUrlResponse) GetBody() *CreatePersonalPublicUrlResponseBody {
	return s.Body
}

func (s *CreatePersonalPublicUrlResponse) SetHeaders(v map[string]*string) *CreatePersonalPublicUrlResponse {
	s.Headers = v
	return s
}

func (s *CreatePersonalPublicUrlResponse) SetStatusCode(v int32) *CreatePersonalPublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePersonalPublicUrlResponse) SetBody(v *CreatePersonalPublicUrlResponseBody) *CreatePersonalPublicUrlResponse {
	s.Body = v
	return s
}

func (s *CreatePersonalPublicUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
