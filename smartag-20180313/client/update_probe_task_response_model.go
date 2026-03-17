// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProbeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProbeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProbeTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProbeTaskResponseBody) *UpdateProbeTaskResponse
	GetBody() *UpdateProbeTaskResponseBody
}

type UpdateProbeTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProbeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProbeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProbeTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateProbeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProbeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProbeTaskResponse) GetBody() *UpdateProbeTaskResponseBody {
	return s.Body
}

func (s *UpdateProbeTaskResponse) SetHeaders(v map[string]*string) *UpdateProbeTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateProbeTaskResponse) SetStatusCode(v int32) *UpdateProbeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProbeTaskResponse) SetBody(v *UpdateProbeTaskResponseBody) *UpdateProbeTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateProbeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
