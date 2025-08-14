// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectRouterAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectRouterAssociationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectRouterAssociationResponseBody) *ModifyExpressConnectRouterAssociationResponse
	GetBody() *ModifyExpressConnectRouterAssociationResponseBody
}

type ModifyExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectRouterAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectRouterAssociationResponse) GetBody() *ModifyExpressConnectRouterAssociationResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponse) SetBody(v *ModifyExpressConnectRouterAssociationResponseBody) *ModifyExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationResponse) Validate() error {
	return dara.Validate(s)
}
