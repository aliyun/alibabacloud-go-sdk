// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *ListSchedulesResponseBody) *ListSchedulesResponse
	GetBody() *ListSchedulesResponseBody
}

type ListSchedulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSchedulesResponse) GetBody() *ListSchedulesResponseBody {
	return s.Body
}

func (s *ListSchedulesResponse) SetHeaders(v map[string]*string) *ListSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListSchedulesResponse) SetStatusCode(v int32) *ListSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchedulesResponse) SetBody(v *ListSchedulesResponseBody) *ListSchedulesResponse {
	s.Body = v
	return s
}

func (s *ListSchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
