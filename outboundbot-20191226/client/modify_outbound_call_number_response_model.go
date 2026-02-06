// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOutboundCallNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOutboundCallNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOutboundCallNumberResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOutboundCallNumberResponseBody) *ModifyOutboundCallNumberResponse
	GetBody() *ModifyOutboundCallNumberResponseBody
}

type ModifyOutboundCallNumberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOutboundCallNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOutboundCallNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOutboundCallNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyOutboundCallNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOutboundCallNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOutboundCallNumberResponse) GetBody() *ModifyOutboundCallNumberResponseBody {
	return s.Body
}

func (s *ModifyOutboundCallNumberResponse) SetHeaders(v map[string]*string) *ModifyOutboundCallNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyOutboundCallNumberResponse) SetStatusCode(v int32) *ModifyOutboundCallNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOutboundCallNumberResponse) SetBody(v *ModifyOutboundCallNumberResponseBody) *ModifyOutboundCallNumberResponse {
	s.Body = v
	return s
}

func (s *ModifyOutboundCallNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
