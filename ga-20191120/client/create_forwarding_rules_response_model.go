// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateForwardingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateForwardingRulesResponse
	GetStatusCode() *int32
	SetBody(v *CreateForwardingRulesResponseBody) *CreateForwardingRulesResponse
	GetBody() *CreateForwardingRulesResponseBody
}

type CreateForwardingRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateForwardingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateForwardingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateForwardingRulesResponse) GetBody() *CreateForwardingRulesResponseBody {
	return s.Body
}

func (s *CreateForwardingRulesResponse) SetHeaders(v map[string]*string) *CreateForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateForwardingRulesResponse) SetStatusCode(v int32) *CreateForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateForwardingRulesResponse) SetBody(v *CreateForwardingRulesResponseBody) *CreateForwardingRulesResponse {
	s.Body = v
	return s
}

func (s *CreateForwardingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
