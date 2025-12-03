// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCheckRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCheckRunResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCheckRunResponseBody) *UpdateCheckRunResponse
	GetBody() *UpdateCheckRunResponseBody
}

type UpdateCheckRunResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCheckRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCheckRunResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckRunResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCheckRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCheckRunResponse) GetBody() *UpdateCheckRunResponseBody {
	return s.Body
}

func (s *UpdateCheckRunResponse) SetHeaders(v map[string]*string) *UpdateCheckRunResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckRunResponse) SetStatusCode(v int32) *UpdateCheckRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCheckRunResponse) SetBody(v *UpdateCheckRunResponseBody) *UpdateCheckRunResponse {
	s.Body = v
	return s
}

func (s *UpdateCheckRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
