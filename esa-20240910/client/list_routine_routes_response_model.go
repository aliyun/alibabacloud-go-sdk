// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutineRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutineRoutesResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutineRoutesResponseBody) *ListRoutineRoutesResponse
	GetBody() *ListRoutineRoutesResponseBody
}

type ListRoutineRoutesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutineRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutineRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListRoutineRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutineRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutineRoutesResponse) GetBody() *ListRoutineRoutesResponseBody {
	return s.Body
}

func (s *ListRoutineRoutesResponse) SetHeaders(v map[string]*string) *ListRoutineRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListRoutineRoutesResponse) SetStatusCode(v int32) *ListRoutineRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutineRoutesResponse) SetBody(v *ListRoutineRoutesResponseBody) *ListRoutineRoutesResponse {
	s.Body = v
	return s
}

func (s *ListRoutineRoutesResponse) Validate() error {
	return dara.Validate(s)
}
