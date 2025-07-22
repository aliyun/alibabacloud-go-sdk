// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterAssociationAllowedPrefixResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse
	GetStatusCode() *int32
	SetBody(v *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse
	GetBody() *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody
}

type ModifyExpressConnectRouterAssociationAllowedPrefixResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) GetBody() *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	return s.Body
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetBody(v *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.Body = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) Validate() error {
	return dara.Validate(s)
}
