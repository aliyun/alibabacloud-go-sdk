// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShiftSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListShiftSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListShiftSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *ListShiftSchedulesResponseBody) *ListShiftSchedulesResponse
	GetBody() *ListShiftSchedulesResponseBody
}

type ListShiftSchedulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShiftSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListShiftSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListShiftSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListShiftSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListShiftSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListShiftSchedulesResponse) GetBody() *ListShiftSchedulesResponseBody {
	return s.Body
}

func (s *ListShiftSchedulesResponse) SetHeaders(v map[string]*string) *ListShiftSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListShiftSchedulesResponse) SetStatusCode(v int32) *ListShiftSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShiftSchedulesResponse) SetBody(v *ListShiftSchedulesResponseBody) *ListShiftSchedulesResponse {
	s.Body = v
	return s
}

func (s *ListShiftSchedulesResponse) Validate() error {
	return dara.Validate(s)
}
