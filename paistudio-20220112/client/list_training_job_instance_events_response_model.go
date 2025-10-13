// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobInstanceEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrainingJobInstanceEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrainingJobInstanceEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrainingJobInstanceEventsResponseBody) *ListTrainingJobInstanceEventsResponse
	GetBody() *ListTrainingJobInstanceEventsResponseBody
}

type ListTrainingJobInstanceEventsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobInstanceEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobInstanceEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceEventsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrainingJobInstanceEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrainingJobInstanceEventsResponse) GetBody() *ListTrainingJobInstanceEventsResponseBody {
	return s.Body
}

func (s *ListTrainingJobInstanceEventsResponse) SetHeaders(v map[string]*string) *ListTrainingJobInstanceEventsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobInstanceEventsResponse) SetStatusCode(v int32) *ListTrainingJobInstanceEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobInstanceEventsResponse) SetBody(v *ListTrainingJobInstanceEventsResponseBody) *ListTrainingJobInstanceEventsResponse {
	s.Body = v
	return s
}

func (s *ListTrainingJobInstanceEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
