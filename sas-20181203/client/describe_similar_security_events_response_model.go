// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarSecurityEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSimilarSecurityEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSimilarSecurityEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSimilarSecurityEventsResponseBody) *DescribeSimilarSecurityEventsResponse
	GetBody() *DescribeSimilarSecurityEventsResponseBody
}

type DescribeSimilarSecurityEventsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSimilarSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSimilarSecurityEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSimilarSecurityEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSimilarSecurityEventsResponse) GetBody() *DescribeSimilarSecurityEventsResponseBody {
	return s.Body
}

func (s *DescribeSimilarSecurityEventsResponse) SetHeaders(v map[string]*string) *DescribeSimilarSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimilarSecurityEventsResponse) SetStatusCode(v int32) *DescribeSimilarSecurityEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponse) SetBody(v *DescribeSimilarSecurityEventsResponseBody) *DescribeSimilarSecurityEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeSimilarSecurityEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
