// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOutboundCallRestrictionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOutboundCallRestrictionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOutboundCallRestrictionResponse
	GetStatusCode() *int32
	SetBody(v *CreateOutboundCallRestrictionResponseBody) *CreateOutboundCallRestrictionResponse
	GetBody() *CreateOutboundCallRestrictionResponseBody
}

type CreateOutboundCallRestrictionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOutboundCallRestrictionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOutboundCallRestrictionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundCallRestrictionResponse) GoString() string {
	return s.String()
}

func (s *CreateOutboundCallRestrictionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOutboundCallRestrictionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOutboundCallRestrictionResponse) GetBody() *CreateOutboundCallRestrictionResponseBody {
	return s.Body
}

func (s *CreateOutboundCallRestrictionResponse) SetHeaders(v map[string]*string) *CreateOutboundCallRestrictionResponse {
	s.Headers = v
	return s
}

func (s *CreateOutboundCallRestrictionResponse) SetStatusCode(v int32) *CreateOutboundCallRestrictionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOutboundCallRestrictionResponse) SetBody(v *CreateOutboundCallRestrictionResponseBody) *CreateOutboundCallRestrictionResponse {
	s.Body = v
	return s
}

func (s *CreateOutboundCallRestrictionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
