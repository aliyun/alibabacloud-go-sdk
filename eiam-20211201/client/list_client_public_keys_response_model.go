// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientPublicKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientPublicKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientPublicKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListClientPublicKeysResponseBody) *ListClientPublicKeysResponse
	GetBody() *ListClientPublicKeysResponseBody
}

type ListClientPublicKeysResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientPublicKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientPublicKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientPublicKeysResponse) GoString() string {
	return s.String()
}

func (s *ListClientPublicKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientPublicKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientPublicKeysResponse) GetBody() *ListClientPublicKeysResponseBody {
	return s.Body
}

func (s *ListClientPublicKeysResponse) SetHeaders(v map[string]*string) *ListClientPublicKeysResponse {
	s.Headers = v
	return s
}

func (s *ListClientPublicKeysResponse) SetStatusCode(v int32) *ListClientPublicKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientPublicKeysResponse) SetBody(v *ListClientPublicKeysResponseBody) *ListClientPublicKeysResponse {
	s.Body = v
	return s
}

func (s *ListClientPublicKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
