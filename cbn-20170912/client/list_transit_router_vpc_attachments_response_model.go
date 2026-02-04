// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVpcAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterVpcAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterVpcAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterVpcAttachmentsResponseBody) *ListTransitRouterVpcAttachmentsResponse
	GetBody() *ListTransitRouterVpcAttachmentsResponseBody
}

type ListTransitRouterVpcAttachmentsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterVpcAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterVpcAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpcAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpcAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterVpcAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterVpcAttachmentsResponse) GetBody() *ListTransitRouterVpcAttachmentsResponseBody {
	return s.Body
}

func (s *ListTransitRouterVpcAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterVpcAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponse) SetStatusCode(v int32) *ListTransitRouterVpcAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponse) SetBody(v *ListTransitRouterVpcAttachmentsResponseBody) *ListTransitRouterVpcAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterVpcAttachmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
