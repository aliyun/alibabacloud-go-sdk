// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSwimmingLanesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllSwimmingLanesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllSwimmingLanesResponse
	GetStatusCode() *int32
	SetBody(v *ListAllSwimmingLanesResponseBody) *ListAllSwimmingLanesResponse
	GetBody() *ListAllSwimmingLanesResponseBody
}

type ListAllSwimmingLanesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllSwimmingLanesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllSwimmingLanesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLanesResponse) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLanesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllSwimmingLanesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllSwimmingLanesResponse) GetBody() *ListAllSwimmingLanesResponseBody {
	return s.Body
}

func (s *ListAllSwimmingLanesResponse) SetHeaders(v map[string]*string) *ListAllSwimmingLanesResponse {
	s.Headers = v
	return s
}

func (s *ListAllSwimmingLanesResponse) SetStatusCode(v int32) *ListAllSwimmingLanesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllSwimmingLanesResponse) SetBody(v *ListAllSwimmingLanesResponseBody) *ListAllSwimmingLanesResponse {
	s.Body = v
	return s
}

func (s *ListAllSwimmingLanesResponse) Validate() error {
	return dara.Validate(s)
}
