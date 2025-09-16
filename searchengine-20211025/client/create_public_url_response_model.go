// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePublicUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePublicUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePublicUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreatePublicUrlResponseBody) *CreatePublicUrlResponse
	GetBody() *CreatePublicUrlResponseBody
}

type CreatePublicUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePublicUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePublicUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePublicUrlResponse) GoString() string {
	return s.String()
}

func (s *CreatePublicUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePublicUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePublicUrlResponse) GetBody() *CreatePublicUrlResponseBody {
	return s.Body
}

func (s *CreatePublicUrlResponse) SetHeaders(v map[string]*string) *CreatePublicUrlResponse {
	s.Headers = v
	return s
}

func (s *CreatePublicUrlResponse) SetStatusCode(v int32) *CreatePublicUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePublicUrlResponse) SetBody(v *CreatePublicUrlResponseBody) *CreatePublicUrlResponse {
	s.Body = v
	return s
}

func (s *CreatePublicUrlResponse) Validate() error {
	return dara.Validate(s)
}
