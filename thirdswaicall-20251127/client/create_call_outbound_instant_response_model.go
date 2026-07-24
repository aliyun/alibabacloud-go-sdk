// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallOutboundInstantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCallOutboundInstantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCallOutboundInstantResponse
	GetStatusCode() *int32
	SetBody(v *CreateCallOutboundInstantResponseBody) *CreateCallOutboundInstantResponse
	GetBody() *CreateCallOutboundInstantResponseBody
}

type CreateCallOutboundInstantResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallOutboundInstantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCallOutboundInstantResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCallOutboundInstantResponse) GoString() string {
	return s.String()
}

func (s *CreateCallOutboundInstantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCallOutboundInstantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCallOutboundInstantResponse) GetBody() *CreateCallOutboundInstantResponseBody {
	return s.Body
}

func (s *CreateCallOutboundInstantResponse) SetHeaders(v map[string]*string) *CreateCallOutboundInstantResponse {
	s.Headers = v
	return s
}

func (s *CreateCallOutboundInstantResponse) SetStatusCode(v int32) *CreateCallOutboundInstantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallOutboundInstantResponse) SetBody(v *CreateCallOutboundInstantResponseBody) *CreateCallOutboundInstantResponse {
	s.Body = v
	return s
}

func (s *CreateCallOutboundInstantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
