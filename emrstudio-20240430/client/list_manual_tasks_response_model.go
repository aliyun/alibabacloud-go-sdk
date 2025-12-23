// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManualTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManualTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListManualTasksResponseBody) *ListManualTasksResponse
	GetBody() *ListManualTasksResponseBody
}

type ListManualTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManualTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManualTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManualTasksResponse) GoString() string {
	return s.String()
}

func (s *ListManualTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManualTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManualTasksResponse) GetBody() *ListManualTasksResponseBody {
	return s.Body
}

func (s *ListManualTasksResponse) SetHeaders(v map[string]*string) *ListManualTasksResponse {
	s.Headers = v
	return s
}

func (s *ListManualTasksResponse) SetStatusCode(v int32) *ListManualTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManualTasksResponse) SetBody(v *ListManualTasksResponseBody) *ListManualTasksResponse {
	s.Body = v
	return s
}

func (s *ListManualTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
