// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimedViewAttitudeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTimedViewAttitudeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTimedViewAttitudeResponse
	GetStatusCode() *int32
	SetBody(v *ListTimedViewAttitudeResponseBody) *ListTimedViewAttitudeResponse
	GetBody() *ListTimedViewAttitudeResponseBody
}

type ListTimedViewAttitudeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTimedViewAttitudeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTimedViewAttitudeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTimedViewAttitudeResponse) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTimedViewAttitudeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTimedViewAttitudeResponse) GetBody() *ListTimedViewAttitudeResponseBody {
	return s.Body
}

func (s *ListTimedViewAttitudeResponse) SetHeaders(v map[string]*string) *ListTimedViewAttitudeResponse {
	s.Headers = v
	return s
}

func (s *ListTimedViewAttitudeResponse) SetStatusCode(v int32) *ListTimedViewAttitudeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTimedViewAttitudeResponse) SetBody(v *ListTimedViewAttitudeResponseBody) *ListTimedViewAttitudeResponse {
	s.Body = v
	return s
}

func (s *ListTimedViewAttitudeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
