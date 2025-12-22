// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeedbacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeedbacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeedbacksResponse
	GetStatusCode() *int32
	SetBody(v *ListFeedbacksResponseBody) *ListFeedbacksResponse
	GetBody() *ListFeedbacksResponseBody
}

type ListFeedbacksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeedbacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeedbacksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeedbacksResponse) GoString() string {
	return s.String()
}

func (s *ListFeedbacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeedbacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeedbacksResponse) GetBody() *ListFeedbacksResponseBody {
	return s.Body
}

func (s *ListFeedbacksResponse) SetHeaders(v map[string]*string) *ListFeedbacksResponse {
	s.Headers = v
	return s
}

func (s *ListFeedbacksResponse) SetStatusCode(v int32) *ListFeedbacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeedbacksResponse) SetBody(v *ListFeedbacksResponseBody) *ListFeedbacksResponse {
	s.Body = v
	return s
}

func (s *ListFeedbacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
