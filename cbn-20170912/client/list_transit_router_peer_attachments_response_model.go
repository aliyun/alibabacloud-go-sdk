// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterPeerAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterPeerAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterPeerAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterPeerAttachmentsResponseBody) *ListTransitRouterPeerAttachmentsResponse
	GetBody() *ListTransitRouterPeerAttachmentsResponseBody
}

type ListTransitRouterPeerAttachmentsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterPeerAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterPeerAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPeerAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPeerAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterPeerAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterPeerAttachmentsResponse) GetBody() *ListTransitRouterPeerAttachmentsResponseBody {
	return s.Body
}

func (s *ListTransitRouterPeerAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterPeerAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponse) SetStatusCode(v int32) *ListTransitRouterPeerAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponse) SetBody(v *ListTransitRouterPeerAttachmentsResponseBody) *ListTransitRouterPeerAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterPeerAttachmentsResponse) Validate() error {
	return dara.Validate(s)
}
