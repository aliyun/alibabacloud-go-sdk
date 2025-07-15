// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyPhysicalConnectionLOAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyPhysicalConnectionLOAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyPhysicalConnectionLOAResponse
	GetStatusCode() *int32
	SetBody(v *ApplyPhysicalConnectionLOAResponseBody) *ApplyPhysicalConnectionLOAResponse
	GetBody() *ApplyPhysicalConnectionLOAResponseBody
}

type ApplyPhysicalConnectionLOAResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyPhysicalConnectionLOAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyPhysicalConnectionLOAResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyPhysicalConnectionLOAResponse) GoString() string {
	return s.String()
}

func (s *ApplyPhysicalConnectionLOAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyPhysicalConnectionLOAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyPhysicalConnectionLOAResponse) GetBody() *ApplyPhysicalConnectionLOAResponseBody {
	return s.Body
}

func (s *ApplyPhysicalConnectionLOAResponse) SetHeaders(v map[string]*string) *ApplyPhysicalConnectionLOAResponse {
	s.Headers = v
	return s
}

func (s *ApplyPhysicalConnectionLOAResponse) SetStatusCode(v int32) *ApplyPhysicalConnectionLOAResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyPhysicalConnectionLOAResponse) SetBody(v *ApplyPhysicalConnectionLOAResponseBody) *ApplyPhysicalConnectionLOAResponse {
	s.Body = v
	return s
}

func (s *ApplyPhysicalConnectionLOAResponse) Validate() error {
	return dara.Validate(s)
}
