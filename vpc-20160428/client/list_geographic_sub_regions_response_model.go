// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeographicSubRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGeographicSubRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGeographicSubRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListGeographicSubRegionsResponseBody) *ListGeographicSubRegionsResponse
	GetBody() *ListGeographicSubRegionsResponseBody
}

type ListGeographicSubRegionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGeographicSubRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGeographicSubRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGeographicSubRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListGeographicSubRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGeographicSubRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGeographicSubRegionsResponse) GetBody() *ListGeographicSubRegionsResponseBody {
	return s.Body
}

func (s *ListGeographicSubRegionsResponse) SetHeaders(v map[string]*string) *ListGeographicSubRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListGeographicSubRegionsResponse) SetStatusCode(v int32) *ListGeographicSubRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGeographicSubRegionsResponse) SetBody(v *ListGeographicSubRegionsResponseBody) *ListGeographicSubRegionsResponse {
	s.Body = v
	return s
}

func (s *ListGeographicSubRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
