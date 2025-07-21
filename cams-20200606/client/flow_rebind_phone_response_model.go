// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowRebindPhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlowRebindPhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlowRebindPhoneResponse
	GetStatusCode() *int32
	SetBody(v *FlowRebindPhoneResponseBody) *FlowRebindPhoneResponse
	GetBody() *FlowRebindPhoneResponseBody
}

type FlowRebindPhoneResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowRebindPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlowRebindPhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s FlowRebindPhoneResponse) GoString() string {
	return s.String()
}

func (s *FlowRebindPhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlowRebindPhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlowRebindPhoneResponse) GetBody() *FlowRebindPhoneResponseBody {
	return s.Body
}

func (s *FlowRebindPhoneResponse) SetHeaders(v map[string]*string) *FlowRebindPhoneResponse {
	s.Headers = v
	return s
}

func (s *FlowRebindPhoneResponse) SetStatusCode(v int32) *FlowRebindPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *FlowRebindPhoneResponse) SetBody(v *FlowRebindPhoneResponseBody) *FlowRebindPhoneResponse {
	s.Body = v
	return s
}

func (s *FlowRebindPhoneResponse) Validate() error {
	return dara.Validate(s)
}
