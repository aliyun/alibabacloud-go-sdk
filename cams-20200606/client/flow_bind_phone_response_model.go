// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowBindPhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlowBindPhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlowBindPhoneResponse
	GetStatusCode() *int32
	SetBody(v *FlowBindPhoneResponseBody) *FlowBindPhoneResponse
	GetBody() *FlowBindPhoneResponseBody
}

type FlowBindPhoneResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlowBindPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlowBindPhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s FlowBindPhoneResponse) GoString() string {
	return s.String()
}

func (s *FlowBindPhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlowBindPhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlowBindPhoneResponse) GetBody() *FlowBindPhoneResponseBody {
	return s.Body
}

func (s *FlowBindPhoneResponse) SetHeaders(v map[string]*string) *FlowBindPhoneResponse {
	s.Headers = v
	return s
}

func (s *FlowBindPhoneResponse) SetStatusCode(v int32) *FlowBindPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *FlowBindPhoneResponse) SetBody(v *FlowBindPhoneResponseBody) *FlowBindPhoneResponse {
	s.Body = v
	return s
}

func (s *FlowBindPhoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
