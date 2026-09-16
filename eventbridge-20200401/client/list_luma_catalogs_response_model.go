// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaCatalogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLumaCatalogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLumaCatalogsResponse
	GetStatusCode() *int32
	SetBody(v *ListLumaCatalogsResponseBody) *ListLumaCatalogsResponse
	GetBody() *ListLumaCatalogsResponseBody
}

type ListLumaCatalogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLumaCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLumaCatalogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsResponse) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLumaCatalogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLumaCatalogsResponse) GetBody() *ListLumaCatalogsResponseBody {
	return s.Body
}

func (s *ListLumaCatalogsResponse) SetHeaders(v map[string]*string) *ListLumaCatalogsResponse {
	s.Headers = v
	return s
}

func (s *ListLumaCatalogsResponse) SetStatusCode(v int32) *ListLumaCatalogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLumaCatalogsResponse) SetBody(v *ListLumaCatalogsResponseBody) *ListLumaCatalogsResponse {
	s.Body = v
	return s
}

func (s *ListLumaCatalogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
