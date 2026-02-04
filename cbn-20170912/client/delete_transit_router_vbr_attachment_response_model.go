// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterVbrAttachmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterVbrAttachmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterVbrAttachmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterVbrAttachmentResponseBody) *DeleteTransitRouterVbrAttachmentResponse
	GetBody() *DeleteTransitRouterVbrAttachmentResponseBody
}

type DeleteTransitRouterVbrAttachmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterVbrAttachmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterVbrAttachmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterVbrAttachmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterVbrAttachmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterVbrAttachmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterVbrAttachmentResponse) GetBody() *DeleteTransitRouterVbrAttachmentResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterVbrAttachmentResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterVbrAttachmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentResponse) SetStatusCode(v int32) *DeleteTransitRouterVbrAttachmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentResponse) SetBody(v *DeleteTransitRouterVbrAttachmentResponseBody) *DeleteTransitRouterVbrAttachmentResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterVbrAttachmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
