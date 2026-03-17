// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProbeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProbeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProbeTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateProbeTaskResponseBody) *CreateProbeTaskResponse
	GetBody() *CreateProbeTaskResponseBody
}

type CreateProbeTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProbeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProbeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProbeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateProbeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProbeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProbeTaskResponse) GetBody() *CreateProbeTaskResponseBody {
	return s.Body
}

func (s *CreateProbeTaskResponse) SetHeaders(v map[string]*string) *CreateProbeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateProbeTaskResponse) SetStatusCode(v int32) *CreateProbeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProbeTaskResponse) SetBody(v *CreateProbeTaskResponseBody) *CreateProbeTaskResponse {
	s.Body = v
	return s
}

func (s *CreateProbeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
