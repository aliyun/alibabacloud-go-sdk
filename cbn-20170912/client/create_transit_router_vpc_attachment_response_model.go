// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVpcAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterVpcAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterVpcAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterVpcAttachmentResponseBody) *CreateTransitRouterVpcAttachmentResponse
	GetBody() *CreateTransitRouterVpcAttachmentResponseBody
}

type CreateTransitRouterVpcAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterVpcAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterVpcAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVpcAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVpcAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterVpcAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterVpcAttachmentResponse) GetBody() *CreateTransitRouterVpcAttachmentResponseBody {
	return s.Body
}

func (s *CreateTransitRouterVpcAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterVpcAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentResponse) SetStatusCode(v int32) *CreateTransitRouterVpcAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterVpcAttachmentResponse) SetBody(v *CreateTransitRouterVpcAttachmentResponseBody) *CreateTransitRouterVpcAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterVpcAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
