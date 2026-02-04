// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVbrAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterVbrAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterVbrAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterVbrAttachmentsResponseBody) *ListTransitRouterVbrAttachmentsResponse
	GetBody() *ListTransitRouterVbrAttachmentsResponseBody
}

type ListTransitRouterVbrAttachmentsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterVbrAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterVbrAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVbrAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVbrAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterVbrAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterVbrAttachmentsResponse) GetBody() *ListTransitRouterVbrAttachmentsResponseBody {
	return s.Body
}

func (s *ListTransitRouterVbrAttachmentsResponse) SetHeaders(v map[string]*string) *ListTransitRouterVbrAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponse) SetStatusCode(v int32) *ListTransitRouterVbrAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponse) SetBody(v *ListTransitRouterVbrAttachmentsResponseBody) *ListTransitRouterVbrAttachmentsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterVbrAttachmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
