// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPrefixListAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterPrefixListAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterPrefixListAssociationResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterPrefixListAssociationResponseBody) *CreateTransitRouterPrefixListAssociationResponse
	GetBody() *CreateTransitRouterPrefixListAssociationResponseBody
}

type CreateTransitRouterPrefixListAssociationResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterPrefixListAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterPrefixListAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPrefixListAssociationResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPrefixListAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterPrefixListAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterPrefixListAssociationResponse) GetBody() *CreateTransitRouterPrefixListAssociationResponseBody {
	return s.Body
}

func (s *CreateTransitRouterPrefixListAssociationResponse) SetHeaders(v map[string]*string) *CreateTransitRouterPrefixListAssociationResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationResponse) SetStatusCode(v int32) *CreateTransitRouterPrefixListAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationResponse) SetBody(v *CreateTransitRouterPrefixListAssociationResponseBody) *CreateTransitRouterPrefixListAssociationResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationResponse) Validate() error {
	return dara.Validate(s)
}
