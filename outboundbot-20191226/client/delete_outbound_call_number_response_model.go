// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundCallNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOutboundCallNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOutboundCallNumberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOutboundCallNumberResponseBody) *DeleteOutboundCallNumberResponse
	GetBody() *DeleteOutboundCallNumberResponseBody
}

type DeleteOutboundCallNumberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOutboundCallNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOutboundCallNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundCallNumberResponse) GoString() string {
	return s.String()
}

func (s *DeleteOutboundCallNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOutboundCallNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOutboundCallNumberResponse) GetBody() *DeleteOutboundCallNumberResponseBody {
	return s.Body
}

func (s *DeleteOutboundCallNumberResponse) SetHeaders(v map[string]*string) *DeleteOutboundCallNumberResponse {
	s.Headers = v
	return s
}

func (s *DeleteOutboundCallNumberResponse) SetStatusCode(v int32) *DeleteOutboundCallNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOutboundCallNumberResponse) SetBody(v *DeleteOutboundCallNumberResponseBody) *DeleteOutboundCallNumberResponse {
	s.Body = v
	return s
}

func (s *DeleteOutboundCallNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
