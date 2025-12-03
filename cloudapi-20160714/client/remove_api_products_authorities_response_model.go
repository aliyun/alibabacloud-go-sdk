// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApiProductsAuthoritiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveApiProductsAuthoritiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveApiProductsAuthoritiesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveApiProductsAuthoritiesResponseBody) *RemoveApiProductsAuthoritiesResponse
	GetBody() *RemoveApiProductsAuthoritiesResponseBody
}

type RemoveApiProductsAuthoritiesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApiProductsAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApiProductsAuthoritiesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveApiProductsAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *RemoveApiProductsAuthoritiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveApiProductsAuthoritiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveApiProductsAuthoritiesResponse) GetBody() *RemoveApiProductsAuthoritiesResponseBody {
	return s.Body
}

func (s *RemoveApiProductsAuthoritiesResponse) SetHeaders(v map[string]*string) *RemoveApiProductsAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *RemoveApiProductsAuthoritiesResponse) SetStatusCode(v int32) *RemoveApiProductsAuthoritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApiProductsAuthoritiesResponse) SetBody(v *RemoveApiProductsAuthoritiesResponseBody) *RemoveApiProductsAuthoritiesResponse {
	s.Body = v
	return s
}

func (s *RemoveApiProductsAuthoritiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
