// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecursionZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRecursionZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRecursionZoneResponse
	GetStatusCode() *int32
	SetBody(v *AddRecursionZoneResponseBody) *AddRecursionZoneResponse
	GetBody() *AddRecursionZoneResponseBody
}

type AddRecursionZoneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecursionZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRecursionZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRecursionZoneResponse) GoString() string {
	return s.String()
}

func (s *AddRecursionZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRecursionZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRecursionZoneResponse) GetBody() *AddRecursionZoneResponseBody {
	return s.Body
}

func (s *AddRecursionZoneResponse) SetHeaders(v map[string]*string) *AddRecursionZoneResponse {
	s.Headers = v
	return s
}

func (s *AddRecursionZoneResponse) SetStatusCode(v int32) *AddRecursionZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecursionZoneResponse) SetBody(v *AddRecursionZoneResponseBody) *AddRecursionZoneResponse {
	s.Body = v
	return s
}

func (s *AddRecursionZoneResponse) Validate() error {
	return dara.Validate(s)
}
