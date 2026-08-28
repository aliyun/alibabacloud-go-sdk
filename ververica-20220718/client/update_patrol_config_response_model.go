// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatrolConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePatrolConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePatrolConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePatrolConfigResponseBody) *UpdatePatrolConfigResponse
	GetBody() *UpdatePatrolConfigResponseBody
}

type UpdatePatrolConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePatrolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePatrolConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatrolConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdatePatrolConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePatrolConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePatrolConfigResponse) GetBody() *UpdatePatrolConfigResponseBody {
	return s.Body
}

func (s *UpdatePatrolConfigResponse) SetHeaders(v map[string]*string) *UpdatePatrolConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdatePatrolConfigResponse) SetStatusCode(v int32) *UpdatePatrolConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePatrolConfigResponse) SetBody(v *UpdatePatrolConfigResponseBody) *UpdatePatrolConfigResponse {
	s.Body = v
	return s
}

func (s *UpdatePatrolConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
