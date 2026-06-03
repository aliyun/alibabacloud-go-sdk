// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPnsPackageStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPnsPackageStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *QueryPnsPackageStatisticsResponseBody) *QueryPnsPackageStatisticsResponse
	GetBody() *QueryPnsPackageStatisticsResponseBody
}

type QueryPnsPackageStatisticsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPnsPackageStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPnsPackageStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPnsPackageStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPnsPackageStatisticsResponse) GetBody() *QueryPnsPackageStatisticsResponseBody {
	return s.Body
}

func (s *QueryPnsPackageStatisticsResponse) SetHeaders(v map[string]*string) *QueryPnsPackageStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QueryPnsPackageStatisticsResponse) SetStatusCode(v int32) *QueryPnsPackageStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPnsPackageStatisticsResponse) SetBody(v *QueryPnsPackageStatisticsResponseBody) *QueryPnsPackageStatisticsResponse {
	s.Body = v
	return s
}

func (s *QueryPnsPackageStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
