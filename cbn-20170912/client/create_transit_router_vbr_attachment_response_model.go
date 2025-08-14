// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVbrAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterVbrAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterVbrAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterVbrAttachmentResponseBody) *CreateTransitRouterVbrAttachmentResponse
	GetBody() *CreateTransitRouterVbrAttachmentResponseBody
}

type CreateTransitRouterVbrAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterVbrAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterVbrAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVbrAttachmentResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVbrAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterVbrAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterVbrAttachmentResponse) GetBody() *CreateTransitRouterVbrAttachmentResponseBody {
	return s.Body
}

func (s *CreateTransitRouterVbrAttachmentResponse) SetHeaders(v map[string]*string) *CreateTransitRouterVbrAttachmentResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterVbrAttachmentResponse) SetStatusCode(v int32) *CreateTransitRouterVbrAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentResponse) SetBody(v *CreateTransitRouterVbrAttachmentResponseBody) *CreateTransitRouterVbrAttachmentResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterVbrAttachmentResponse) Validate() error {
	return dara.Validate(s)
}
