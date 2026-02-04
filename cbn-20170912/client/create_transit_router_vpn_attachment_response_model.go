// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVpnAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterVpnAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterVpnAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterVpnAttachmentResponseBody) *CreateTransitRouterVpnAttachmentResponse
	GetBody() *CreateTransitRouterVpnAttachmentResponseBody
}

type CreateTransitRouterVpnAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterVpnAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterVpnAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpnAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpnAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterVpnAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterVpnAttachmentResponse) GetBody() *CreateTransitRouterVpnAttachmentResponseBody {
	return s.Body
}

func (s *CreateTransitRouterVpnAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterVpnAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterVpnAttachmentResponse) SetStatusCode(v int32) *CreateTransitRouterVpnAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterVpnAttachmentResponse) SetBody(v *CreateTransitRouterVpnAttachmentResponseBody) *CreateTransitRouterVpnAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterVpnAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
