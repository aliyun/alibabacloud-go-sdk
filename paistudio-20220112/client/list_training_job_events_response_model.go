// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrainingJobEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrainingJobEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrainingJobEventsResponseBody) *ListTrainingJobEventsResponse
	GetBody() *ListTrainingJobEventsResponseBody
}

type ListTrainingJobEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobEventsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrainingJobEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrainingJobEventsResponse) GetBody() *ListTrainingJobEventsResponseBody {
	return s.Body
}

func (s *ListTrainingJobEventsResponse) SetHeaders(v map[string]*string) *ListTrainingJobEventsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobEventsResponse) SetStatusCode(v int32) *ListTrainingJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobEventsResponse) SetBody(v *ListTrainingJobEventsResponseBody) *ListTrainingJobEventsResponse {
	s.Body = v
	return s
}

func (s *ListTrainingJobEventsResponse) Validate() error {
	return dara.Validate(s)
}
