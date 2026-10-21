// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImageCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImageCachesResponse
	GetStatusCode() *int32
	SetBody(v *ListImageCachesResponseBody) *ListImageCachesResponse
	GetBody() *ListImageCachesResponseBody
}

type ListImageCachesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImageCachesResponse) GoString() string {
	return s.String()
}

func (s *ListImageCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImageCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImageCachesResponse) GetBody() *ListImageCachesResponseBody {
	return s.Body
}

func (s *ListImageCachesResponse) SetHeaders(v map[string]*string) *ListImageCachesResponse {
	s.Headers = v
	return s
}

func (s *ListImageCachesResponse) SetStatusCode(v int32) *ListImageCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageCachesResponse) SetBody(v *ListImageCachesResponseBody) *ListImageCachesResponse {
	s.Body = v
	return s
}

func (s *ListImageCachesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
