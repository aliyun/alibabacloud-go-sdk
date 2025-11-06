// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaultInjectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaultInjectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaultInjectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaultInjectionResponseBody) *DeleteFaultInjectionResponse
	GetBody() *DeleteFaultInjectionResponseBody
}

type DeleteFaultInjectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaultInjectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaultInjectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaultInjectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaultInjectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaultInjectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaultInjectionResponse) GetBody() *DeleteFaultInjectionResponseBody {
	return s.Body
}

func (s *DeleteFaultInjectionResponse) SetHeaders(v map[string]*string) *DeleteFaultInjectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaultInjectionResponse) SetStatusCode(v int32) *DeleteFaultInjectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaultInjectionResponse) SetBody(v *DeleteFaultInjectionResponseBody) *DeleteFaultInjectionResponse {
	s.Body = v
	return s
}

func (s *DeleteFaultInjectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
