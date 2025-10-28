// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMigrateRegionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMigrateRegionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMigrateRegionListResponse
	GetStatusCode() *int32
	SetBody(v *QueryMigrateRegionListResponseBody) *QueryMigrateRegionListResponse
	GetBody() *QueryMigrateRegionListResponseBody
}

type QueryMigrateRegionListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMigrateRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMigrateRegionListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMigrateRegionListResponse) GoString() string {
	return s.String()
}

func (s *QueryMigrateRegionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMigrateRegionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMigrateRegionListResponse) GetBody() *QueryMigrateRegionListResponseBody {
	return s.Body
}

func (s *QueryMigrateRegionListResponse) SetHeaders(v map[string]*string) *QueryMigrateRegionListResponse {
	s.Headers = v
	return s
}

func (s *QueryMigrateRegionListResponse) SetStatusCode(v int32) *QueryMigrateRegionListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMigrateRegionListResponse) SetBody(v *QueryMigrateRegionListResponseBody) *QueryMigrateRegionListResponse {
	s.Body = v
	return s
}

func (s *QueryMigrateRegionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
