// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterVbrAttachmentAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterVbrAttachmentAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterVbrAttachmentAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterVbrAttachmentAttributeResponseBody) *UpdateTransitRouterVbrAttachmentAttributeResponse
	GetBody() *UpdateTransitRouterVbrAttachmentAttributeResponseBody
}

type UpdateTransitRouterVbrAttachmentAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterVbrAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterVbrAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) GetBody() *UpdateTransitRouterVbrAttachmentAttributeResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterVbrAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) SetStatusCode(v int32) *UpdateTransitRouterVbrAttachmentAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterVbrAttachmentAttributeResponseBody) *UpdateTransitRouterVbrAttachmentAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterVbrAttachmentAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
