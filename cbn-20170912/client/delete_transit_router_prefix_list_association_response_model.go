// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPrefixListAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterPrefixListAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterPrefixListAssociationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterPrefixListAssociationResponseBody) *DeleteTransitRouterPrefixListAssociationResponse
	GetBody() *DeleteTransitRouterPrefixListAssociationResponseBody
}

type DeleteTransitRouterPrefixListAssociationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterPrefixListAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterPrefixListAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPrefixListAssociationResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPrefixListAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterPrefixListAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterPrefixListAssociationResponse) GetBody() *DeleteTransitRouterPrefixListAssociationResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterPrefixListAssociationResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterPrefixListAssociationResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationResponse) SetStatusCode(v int32) *DeleteTransitRouterPrefixListAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationResponse) SetBody(v *DeleteTransitRouterPrefixListAssociationResponseBody) *DeleteTransitRouterPrefixListAssociationResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationResponse) Validate() error {
	return dara.Validate(s)
}
