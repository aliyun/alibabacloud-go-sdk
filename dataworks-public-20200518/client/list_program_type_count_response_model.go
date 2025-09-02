// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProgramTypeCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProgramTypeCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProgramTypeCountResponse
	GetStatusCode() *int32
	SetBody(v *ListProgramTypeCountResponseBody) *ListProgramTypeCountResponse
	GetBody() *ListProgramTypeCountResponseBody
}

type ListProgramTypeCountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProgramTypeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProgramTypeCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProgramTypeCountResponse) GoString() string {
	return s.String()
}

func (s *ListProgramTypeCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProgramTypeCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProgramTypeCountResponse) GetBody() *ListProgramTypeCountResponseBody {
	return s.Body
}

func (s *ListProgramTypeCountResponse) SetHeaders(v map[string]*string) *ListProgramTypeCountResponse {
	s.Headers = v
	return s
}

func (s *ListProgramTypeCountResponse) SetStatusCode(v int32) *ListProgramTypeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProgramTypeCountResponse) SetBody(v *ListProgramTypeCountResponseBody) *ListProgramTypeCountResponse {
	s.Body = v
	return s
}

func (s *ListProgramTypeCountResponse) Validate() error {
	return dara.Validate(s)
}
