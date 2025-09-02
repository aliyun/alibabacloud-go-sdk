// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendInstanceResponse
	GetStatusCode() *int32
	SetBody(v *SuspendInstanceResponseBody) *SuspendInstanceResponse
	GetBody() *SuspendInstanceResponseBody
}

type SuspendInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendInstanceResponse) GoString() string {
	return s.String()
}

func (s *SuspendInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendInstanceResponse) GetBody() *SuspendInstanceResponseBody {
	return s.Body
}

func (s *SuspendInstanceResponse) SetHeaders(v map[string]*string) *SuspendInstanceResponse {
	s.Headers = v
	return s
}

func (s *SuspendInstanceResponse) SetStatusCode(v int32) *SuspendInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendInstanceResponse) SetBody(v *SuspendInstanceResponseBody) *SuspendInstanceResponse {
	s.Body = v
	return s
}

func (s *SuspendInstanceResponse) Validate() error {
	return dara.Validate(s)
}
