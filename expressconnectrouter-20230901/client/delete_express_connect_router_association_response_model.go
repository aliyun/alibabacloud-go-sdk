// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectRouterAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExpressConnectRouterAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExpressConnectRouterAssociationResponseBody) *DeleteExpressConnectRouterAssociationResponse
	GetBody() *DeleteExpressConnectRouterAssociationResponseBody
}

type DeleteExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExpressConnectRouterAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExpressConnectRouterAssociationResponse) GetBody() *DeleteExpressConnectRouterAssociationResponseBody {
	return s.Body
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetBody(v *DeleteExpressConnectRouterAssociationResponseBody) *DeleteExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponse) Validate() error {
	return dara.Validate(s)
}
