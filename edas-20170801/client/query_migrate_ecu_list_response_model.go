// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMigrateEcuListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMigrateEcuListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMigrateEcuListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMigrateEcuListResponseBody) *QueryMigrateEcuListResponse
	GetBody() *QueryMigrateEcuListResponseBody
}

type QueryMigrateEcuListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMigrateEcuListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMigrateEcuListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateEcuListResponse) GoString() string {
	return s.String()
}

func (s *QueryMigrateEcuListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMigrateEcuListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMigrateEcuListResponse) GetBody() *QueryMigrateEcuListResponseBody {
	return s.Body
}

func (s *QueryMigrateEcuListResponse) SetHeaders(v map[string]*string) *QueryMigrateEcuListResponse {
	s.Headers = v
	return s
}

func (s *QueryMigrateEcuListResponse) SetStatusCode(v int32) *QueryMigrateEcuListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMigrateEcuListResponse) SetBody(v *QueryMigrateEcuListResponseBody) *QueryMigrateEcuListResponse {
	s.Body = v
	return s
}

func (s *QueryMigrateEcuListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
