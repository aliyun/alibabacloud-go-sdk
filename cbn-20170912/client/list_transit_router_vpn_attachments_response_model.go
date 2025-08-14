// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVpnAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterVpnAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterVpnAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterVpnAttachmentsResponseBody) *ListTransitRouterVpnAttachmentsResponse
	GetBody() *ListTransitRouterVpnAttachmentsResponseBody
}

type ListTransitRouterVpnAttachmentsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterVpnAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterVpnAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterVpnAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterVpnAttachmentsResponse) GetBody() *ListTransitRouterVpnAttachmentsResponseBody {
	return s.Body
}

func (s *ListTransitRouterVpnAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterVpnAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponse) SetStatusCode(v int32) *ListTransitRouterVpnAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponse) SetBody(v *ListTransitRouterVpnAttachmentsResponseBody) *ListTransitRouterVpnAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterVpnAttachmentsResponse) Validate() error {
	return dara.Validate(s)
}
