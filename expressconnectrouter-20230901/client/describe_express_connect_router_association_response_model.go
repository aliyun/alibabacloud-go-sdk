// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterAssociationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAssociationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectRouterAssociationResponseBody) *DescribeExpressConnectRouterAssociationResponse
	GetBody() *DescribeExpressConnectRouterAssociationResponseBody
}

type DescribeExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectRouterAssociationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectRouterAssociationResponse) GetBody() *DescribeExpressConnectRouterAssociationResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetBody(v *DescribeExpressConnectRouterAssociationResponseBody) *DescribeExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponse) Validate() error {
	return dara.Validate(s)
}
