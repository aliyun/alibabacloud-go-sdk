// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncEditTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncEditTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncEditTimelineResponse
	GetStatusCode() *int32
	SetBody(v *AsyncEditTimelineResponseBody) *AsyncEditTimelineResponse
	GetBody() *AsyncEditTimelineResponseBody
}

type AsyncEditTimelineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncEditTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncEditTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncEditTimelineResponse) GoString() string {
	return s.String()
}

func (s *AsyncEditTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncEditTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncEditTimelineResponse) GetBody() *AsyncEditTimelineResponseBody {
	return s.Body
}

func (s *AsyncEditTimelineResponse) SetHeaders(v map[string]*string) *AsyncEditTimelineResponse {
	s.Headers = v
	return s
}

func (s *AsyncEditTimelineResponse) SetStatusCode(v int32) *AsyncEditTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncEditTimelineResponse) SetBody(v *AsyncEditTimelineResponseBody) *AsyncEditTimelineResponse {
	s.Body = v
	return s
}

func (s *AsyncEditTimelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
