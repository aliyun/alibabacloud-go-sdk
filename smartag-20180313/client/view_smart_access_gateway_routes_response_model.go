// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ViewSmartAccessGatewayRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ViewSmartAccessGatewayRoutesResponse
	GetStatusCode() *int32
	SetBody(v *ViewSmartAccessGatewayRoutesResponseBody) *ViewSmartAccessGatewayRoutesResponse
	GetBody() *ViewSmartAccessGatewayRoutesResponseBody
}

type ViewSmartAccessGatewayRoutesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ViewSmartAccessGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ViewSmartAccessGatewayRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ViewSmartAccessGatewayRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ViewSmartAccessGatewayRoutesResponse) GetBody() *ViewSmartAccessGatewayRoutesResponseBody {
	return s.Body
}

func (s *ViewSmartAccessGatewayRoutesResponse) SetHeaders(v map[string]*string) *ViewSmartAccessGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponse) SetStatusCode(v int32) *ViewSmartAccessGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponse) SetBody(v *ViewSmartAccessGatewayRoutesResponseBody) *ViewSmartAccessGatewayRoutesResponse {
	s.Body = v
	return s
}

func (s *ViewSmartAccessGatewayRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
