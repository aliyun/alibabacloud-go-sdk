// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterEcrAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterEcrAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterEcrAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterEcrAttachmentResponseBody) *CreateTransitRouterEcrAttachmentResponse
	GetBody() *CreateTransitRouterEcrAttachmentResponseBody
}

type CreateTransitRouterEcrAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterEcrAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterEcrAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterEcrAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterEcrAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterEcrAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterEcrAttachmentResponse) GetBody() *CreateTransitRouterEcrAttachmentResponseBody {
	return s.Body
}

func (s *CreateTransitRouterEcrAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterEcrAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterEcrAttachmentResponse) SetStatusCode(v int32) *CreateTransitRouterEcrAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentResponse) SetBody(v *CreateTransitRouterEcrAttachmentResponseBody) *CreateTransitRouterEcrAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterEcrAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
