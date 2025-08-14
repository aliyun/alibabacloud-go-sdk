// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterEcrAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterEcrAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterEcrAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterEcrAttachmentsResponseBody) *ListTransitRouterEcrAttachmentsResponse
	GetBody() *ListTransitRouterEcrAttachmentsResponseBody
}

type ListTransitRouterEcrAttachmentsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterEcrAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterEcrAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterEcrAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterEcrAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterEcrAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterEcrAttachmentsResponse) GetBody() *ListTransitRouterEcrAttachmentsResponseBody {
	return s.Body
}

func (s *ListTransitRouterEcrAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterEcrAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponse) SetStatusCode(v int32) *ListTransitRouterEcrAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponse) SetBody(v *ListTransitRouterEcrAttachmentsResponseBody) *ListTransitRouterEcrAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterEcrAttachmentsResponse) Validate() error {
	return dara.Validate(s)
}
