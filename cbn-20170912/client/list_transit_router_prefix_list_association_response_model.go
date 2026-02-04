// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterPrefixListAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterPrefixListAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterPrefixListAssociationResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterPrefixListAssociationResponseBody) *ListTransitRouterPrefixListAssociationResponse
	GetBody() *ListTransitRouterPrefixListAssociationResponseBody
}

type ListTransitRouterPrefixListAssociationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterPrefixListAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterPrefixListAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterPrefixListAssociationResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterPrefixListAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterPrefixListAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterPrefixListAssociationResponse) GetBody() *ListTransitRouterPrefixListAssociationResponseBody {
	return s.Body
}

func (s *ListTransitRouterPrefixListAssociationResponse) SetHeaders(v map[string]*string) *ListTransitRouterPrefixListAssociationResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponse) SetStatusCode(v int32) *ListTransitRouterPrefixListAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponse) SetBody(v *ListTransitRouterPrefixListAssociationResponseBody) *ListTransitRouterPrefixListAssociationResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterPrefixListAssociationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
