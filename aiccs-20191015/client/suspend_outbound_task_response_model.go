// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *SuspendOutboundTaskResponseBody) *SuspendOutboundTaskResponse
	GetBody() *SuspendOutboundTaskResponseBody
}

type SuspendOutboundTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *SuspendOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendOutboundTaskResponse) GetBody() *SuspendOutboundTaskResponseBody {
	return s.Body
}

func (s *SuspendOutboundTaskResponse) SetHeaders(v map[string]*string) *SuspendOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *SuspendOutboundTaskResponse) SetStatusCode(v int32) *SuspendOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendOutboundTaskResponse) SetBody(v *SuspendOutboundTaskResponseBody) *SuspendOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *SuspendOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
