// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowUnbindPhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlowUnbindPhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlowUnbindPhoneResponse
	GetStatusCode() *int32
	SetBody(v *FlowUnbindPhoneResponseBody) *FlowUnbindPhoneResponse
	GetBody() *FlowUnbindPhoneResponseBody
}

type FlowUnbindPhoneResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowUnbindPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlowUnbindPhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s FlowUnbindPhoneResponse) GoString() string {
	return s.String()
}

func (s *FlowUnbindPhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlowUnbindPhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlowUnbindPhoneResponse) GetBody() *FlowUnbindPhoneResponseBody {
	return s.Body
}

func (s *FlowUnbindPhoneResponse) SetHeaders(v map[string]*string) *FlowUnbindPhoneResponse {
	s.Headers = v
	return s
}

func (s *FlowUnbindPhoneResponse) SetStatusCode(v int32) *FlowUnbindPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *FlowUnbindPhoneResponse) SetBody(v *FlowUnbindPhoneResponseBody) *FlowUnbindPhoneResponse {
	s.Body = v
	return s
}

func (s *FlowUnbindPhoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
