// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClickHouseDBTimezonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClickHouseDBTimezonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClickHouseDBTimezonesResponse
	GetStatusCode() *int32
	SetBody(v *ListClickHouseDBTimezonesResponseBody) *ListClickHouseDBTimezonesResponse
	GetBody() *ListClickHouseDBTimezonesResponseBody
}

type ListClickHouseDBTimezonesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClickHouseDBTimezonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClickHouseDBTimezonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClickHouseDBTimezonesResponse) GoString() string {
	return s.String()
}

func (s *ListClickHouseDBTimezonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClickHouseDBTimezonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClickHouseDBTimezonesResponse) GetBody() *ListClickHouseDBTimezonesResponseBody {
	return s.Body
}

func (s *ListClickHouseDBTimezonesResponse) SetHeaders(v map[string]*string) *ListClickHouseDBTimezonesResponse {
	s.Headers = v
	return s
}

func (s *ListClickHouseDBTimezonesResponse) SetStatusCode(v int32) *ListClickHouseDBTimezonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClickHouseDBTimezonesResponse) SetBody(v *ListClickHouseDBTimezonesResponseBody) *ListClickHouseDBTimezonesResponse {
	s.Body = v
	return s
}

func (s *ListClickHouseDBTimezonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
