// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSimilarSecurityEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleSimilarSecurityEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleSimilarSecurityEventsResponse
	GetStatusCode() *int32
	SetBody(v *HandleSimilarSecurityEventsResponseBody) *HandleSimilarSecurityEventsResponse
	GetBody() *HandleSimilarSecurityEventsResponseBody
}

type HandleSimilarSecurityEventsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleSimilarSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleSimilarSecurityEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleSimilarSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSimilarSecurityEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleSimilarSecurityEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleSimilarSecurityEventsResponse) GetBody() *HandleSimilarSecurityEventsResponseBody {
	return s.Body
}

func (s *HandleSimilarSecurityEventsResponse) SetHeaders(v map[string]*string) *HandleSimilarSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *HandleSimilarSecurityEventsResponse) SetStatusCode(v int32) *HandleSimilarSecurityEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleSimilarSecurityEventsResponse) SetBody(v *HandleSimilarSecurityEventsResponseBody) *HandleSimilarSecurityEventsResponse {
	s.Body = v
	return s
}

func (s *HandleSimilarSecurityEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
