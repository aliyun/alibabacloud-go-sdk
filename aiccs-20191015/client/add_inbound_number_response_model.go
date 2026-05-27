// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInboundNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddInboundNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddInboundNumberResponse
	GetStatusCode() *int32
	SetBody(v *AddInboundNumberResponseBody) *AddInboundNumberResponse
	GetBody() *AddInboundNumberResponseBody
}

type AddInboundNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInboundNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInboundNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddInboundNumberResponse) GoString() string {
	return s.String()
}

func (s *AddInboundNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddInboundNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddInboundNumberResponse) GetBody() *AddInboundNumberResponseBody {
	return s.Body
}

func (s *AddInboundNumberResponse) SetHeaders(v map[string]*string) *AddInboundNumberResponse {
	s.Headers = v
	return s
}

func (s *AddInboundNumberResponse) SetStatusCode(v int32) *AddInboundNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInboundNumberResponse) SetBody(v *AddInboundNumberResponseBody) *AddInboundNumberResponse {
	s.Body = v
	return s
}

func (s *AddInboundNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
