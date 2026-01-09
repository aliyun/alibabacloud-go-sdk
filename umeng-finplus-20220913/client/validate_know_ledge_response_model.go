// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateKnowLedgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateKnowLedgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateKnowLedgeResponse
	GetStatusCode() *int32
	SetBody(v *ValidateKnowLedgeResponseBody) *ValidateKnowLedgeResponse
	GetBody() *ValidateKnowLedgeResponseBody
}

type ValidateKnowLedgeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateKnowLedgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateKnowLedgeResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateKnowLedgeResponse) GoString() string {
	return s.String()
}

func (s *ValidateKnowLedgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateKnowLedgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateKnowLedgeResponse) GetBody() *ValidateKnowLedgeResponseBody {
	return s.Body
}

func (s *ValidateKnowLedgeResponse) SetHeaders(v map[string]*string) *ValidateKnowLedgeResponse {
	s.Headers = v
	return s
}

func (s *ValidateKnowLedgeResponse) SetStatusCode(v int32) *ValidateKnowLedgeResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateKnowLedgeResponse) SetBody(v *ValidateKnowLedgeResponseBody) *ValidateKnowLedgeResponse {
	s.Body = v
	return s
}

func (s *ValidateKnowLedgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
