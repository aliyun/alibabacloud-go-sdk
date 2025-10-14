// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteLogsResponseBody) *DescribeSiteLogsResponse
	GetBody() *DescribeSiteLogsResponseBody
}

type DescribeSiteLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteLogsResponse) GetBody() *DescribeSiteLogsResponseBody {
	return s.Body
}

func (s *DescribeSiteLogsResponse) SetHeaders(v map[string]*string) *DescribeSiteLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteLogsResponse) SetStatusCode(v int32) *DescribeSiteLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteLogsResponse) SetBody(v *DescribeSiteLogsResponseBody) *DescribeSiteLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
