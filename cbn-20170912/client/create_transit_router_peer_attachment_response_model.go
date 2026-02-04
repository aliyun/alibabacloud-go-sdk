// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPeerAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterPeerAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterPeerAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterPeerAttachmentResponseBody) *CreateTransitRouterPeerAttachmentResponse
	GetBody() *CreateTransitRouterPeerAttachmentResponseBody
}

type CreateTransitRouterPeerAttachmentResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterPeerAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterPeerAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPeerAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPeerAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterPeerAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterPeerAttachmentResponse) GetBody() *CreateTransitRouterPeerAttachmentResponseBody {
	return s.Body
}

func (s *CreateTransitRouterPeerAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterPeerAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterPeerAttachmentResponse) SetStatusCode(v int32) *CreateTransitRouterPeerAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterPeerAttachmentResponse) SetBody(v *CreateTransitRouterPeerAttachmentResponseBody) *CreateTransitRouterPeerAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterPeerAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
