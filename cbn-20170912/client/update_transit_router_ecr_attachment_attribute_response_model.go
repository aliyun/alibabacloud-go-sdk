// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterEcrAttachmentAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterEcrAttachmentAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterEcrAttachmentAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterEcrAttachmentAttributeResponseBody) *UpdateTransitRouterEcrAttachmentAttributeResponse
	GetBody() *UpdateTransitRouterEcrAttachmentAttributeResponseBody
}

type UpdateTransitRouterEcrAttachmentAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterEcrAttachmentAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterEcrAttachmentAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterEcrAttachmentAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponse) GetBody() *UpdateTransitRouterEcrAttachmentAttributeResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterEcrAttachmentAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponse) SetStatusCode(v int32) *UpdateTransitRouterEcrAttachmentAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponse) SetBody(v *UpdateTransitRouterEcrAttachmentAttributeResponseBody) *UpdateTransitRouterEcrAttachmentAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterEcrAttachmentAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
