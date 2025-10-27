// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDiscoverDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDiscoverDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDiscoverDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *QueryDiscoverDatabaseResponseBody) *QueryDiscoverDatabaseResponse
	GetBody() *QueryDiscoverDatabaseResponseBody
}

type QueryDiscoverDatabaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDiscoverDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDiscoverDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDiscoverDatabaseResponse) GoString() string {
	return s.String()
}

func (s *QueryDiscoverDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDiscoverDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDiscoverDatabaseResponse) GetBody() *QueryDiscoverDatabaseResponseBody {
	return s.Body
}

func (s *QueryDiscoverDatabaseResponse) SetHeaders(v map[string]*string) *QueryDiscoverDatabaseResponse {
	s.Headers = v
	return s
}

func (s *QueryDiscoverDatabaseResponse) SetStatusCode(v int32) *QueryDiscoverDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDiscoverDatabaseResponse) SetBody(v *QueryDiscoverDatabaseResponseBody) *QueryDiscoverDatabaseResponse {
	s.Body = v
	return s
}

func (s *QueryDiscoverDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
