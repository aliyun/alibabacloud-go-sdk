// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVpcAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterVpcAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterVpcAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterVpcAttachmentResponseBody) *DeleteTransitRouterVpcAttachmentResponse
	GetBody() *DeleteTransitRouterVpcAttachmentResponseBody
}

type DeleteTransitRouterVpcAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterVpcAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterVpcAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVpcAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVpcAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterVpcAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterVpcAttachmentResponse) GetBody() *DeleteTransitRouterVpcAttachmentResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterVpcAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterVpcAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentResponse) SetStatusCode(v int32) *DeleteTransitRouterVpcAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentResponse) SetBody(v *DeleteTransitRouterVpcAttachmentResponseBody) *DeleteTransitRouterVpcAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterVpcAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
