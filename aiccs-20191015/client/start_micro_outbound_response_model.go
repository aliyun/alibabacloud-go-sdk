// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMicroOutboundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartMicroOutboundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartMicroOutboundResponse
	GetStatusCode() *int32
	SetBody(v *StartMicroOutboundResponseBody) *StartMicroOutboundResponse
	GetBody() *StartMicroOutboundResponseBody
}

type StartMicroOutboundResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMicroOutboundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMicroOutboundResponse) String() string {
	return dara.Prettify(s)
}

func (s StartMicroOutboundResponse) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartMicroOutboundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartMicroOutboundResponse) GetBody() *StartMicroOutboundResponseBody {
	return s.Body
}

func (s *StartMicroOutboundResponse) SetHeaders(v map[string]*string) *StartMicroOutboundResponse {
	s.Headers = v
	return s
}

func (s *StartMicroOutboundResponse) SetStatusCode(v int32) *StartMicroOutboundResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMicroOutboundResponse) SetBody(v *StartMicroOutboundResponseBody) *StartMicroOutboundResponse {
	s.Body = v
	return s
}

func (s *StartMicroOutboundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
