// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHealthCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHealthCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHealthCheckResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHealthCheckResponseBody) *ModifyHealthCheckResponse
	GetBody() *ModifyHealthCheckResponseBody
}

type ModifyHealthCheckResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHealthCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHealthCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHealthCheckResponse) GetBody() *ModifyHealthCheckResponseBody {
	return s.Body
}

func (s *ModifyHealthCheckResponse) SetHeaders(v map[string]*string) *ModifyHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *ModifyHealthCheckResponse) SetStatusCode(v int32) *ModifyHealthCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHealthCheckResponse) SetBody(v *ModifyHealthCheckResponseBody) *ModifyHealthCheckResponse {
	s.Body = v
	return s
}

func (s *ModifyHealthCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
