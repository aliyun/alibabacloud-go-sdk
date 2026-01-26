// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnCallSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOnCallSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOnCallSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *ListOnCallSchedulesResponseBody) *ListOnCallSchedulesResponse
	GetBody() *ListOnCallSchedulesResponseBody
}

type ListOnCallSchedulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOnCallSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOnCallSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOnCallSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOnCallSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOnCallSchedulesResponse) GetBody() *ListOnCallSchedulesResponseBody {
	return s.Body
}

func (s *ListOnCallSchedulesResponse) SetHeaders(v map[string]*string) *ListOnCallSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListOnCallSchedulesResponse) SetStatusCode(v int32) *ListOnCallSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnCallSchedulesResponse) SetBody(v *ListOnCallSchedulesResponseBody) *ListOnCallSchedulesResponse {
	s.Body = v
	return s
}

func (s *ListOnCallSchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
