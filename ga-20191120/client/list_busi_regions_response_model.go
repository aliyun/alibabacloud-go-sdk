// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusiRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBusiRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBusiRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListBusiRegionsResponseBody) *ListBusiRegionsResponse
	GetBody() *ListBusiRegionsResponseBody
}

type ListBusiRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBusiRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBusiRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBusiRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBusiRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBusiRegionsResponse) GetBody() *ListBusiRegionsResponseBody {
	return s.Body
}

func (s *ListBusiRegionsResponse) SetHeaders(v map[string]*string) *ListBusiRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListBusiRegionsResponse) SetStatusCode(v int32) *ListBusiRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBusiRegionsResponse) SetBody(v *ListBusiRegionsResponseBody) *ListBusiRegionsResponse {
	s.Body = v
	return s
}

func (s *ListBusiRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
