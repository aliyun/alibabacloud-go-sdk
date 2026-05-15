// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *RestartOutboundTaskResponseBody) *RestartOutboundTaskResponse
	GetBody() *RestartOutboundTaskResponseBody
}

type RestartOutboundTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *RestartOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartOutboundTaskResponse) GetBody() *RestartOutboundTaskResponseBody {
	return s.Body
}

func (s *RestartOutboundTaskResponse) SetHeaders(v map[string]*string) *RestartOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *RestartOutboundTaskResponse) SetStatusCode(v int32) *RestartOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartOutboundTaskResponse) SetBody(v *RestartOutboundTaskResponseBody) *RestartOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *RestartOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
