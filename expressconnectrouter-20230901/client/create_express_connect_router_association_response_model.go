// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectRouterAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExpressConnectRouterAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExpressConnectRouterAssociationResponse
	GetStatusCode() *int32
	SetBody(v *CreateExpressConnectRouterAssociationResponseBody) *CreateExpressConnectRouterAssociationResponse
	GetBody() *CreateExpressConnectRouterAssociationResponseBody
}

type CreateExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectRouterAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExpressConnectRouterAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExpressConnectRouterAssociationResponse) GetBody() *CreateExpressConnectRouterAssociationResponseBody {
	return s.Body
}

func (s *CreateExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *CreateExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *CreateExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponse) SetBody(v *CreateExpressConnectRouterAssociationResponseBody) *CreateExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponse) Validate() error {
	return dara.Validate(s)
}
