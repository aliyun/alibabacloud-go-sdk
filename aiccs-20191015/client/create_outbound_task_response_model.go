// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateOutboundTaskResponseBody) *CreateOutboundTaskResponse
	GetBody() *CreateOutboundTaskResponseBody
}

type CreateOutboundTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOutboundTaskResponse) GetBody() *CreateOutboundTaskResponseBody {
	return s.Body
}

func (s *CreateOutboundTaskResponse) SetHeaders(v map[string]*string) *CreateOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOutboundTaskResponse) SetStatusCode(v int32) *CreateOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOutboundTaskResponse) SetBody(v *CreateOutboundTaskResponseBody) *CreateOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *CreateOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
