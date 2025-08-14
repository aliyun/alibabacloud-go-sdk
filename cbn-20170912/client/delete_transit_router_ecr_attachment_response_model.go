// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterEcrAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterEcrAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterEcrAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterEcrAttachmentResponseBody) *DeleteTransitRouterEcrAttachmentResponse
	GetBody() *DeleteTransitRouterEcrAttachmentResponseBody
}

type DeleteTransitRouterEcrAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterEcrAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterEcrAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterEcrAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterEcrAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterEcrAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterEcrAttachmentResponse) GetBody() *DeleteTransitRouterEcrAttachmentResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterEcrAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterEcrAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentResponse) SetStatusCode(v int32) *DeleteTransitRouterEcrAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentResponse) SetBody(v *DeleteTransitRouterEcrAttachmentResponseBody) *DeleteTransitRouterEcrAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterEcrAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
