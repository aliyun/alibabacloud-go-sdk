// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayApiGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayApiGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayApiGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayApiGroupsResponseBody) *ListGatewayApiGroupsResponse
	GetBody() *ListGatewayApiGroupsResponseBody
}

type ListGatewayApiGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayApiGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayApiGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayApiGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayApiGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayApiGroupsResponse) GetBody() *ListGatewayApiGroupsResponseBody {
	return s.Body
}

func (s *ListGatewayApiGroupsResponse) SetHeaders(v map[string]*string) *ListGatewayApiGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayApiGroupsResponse) SetStatusCode(v int32) *ListGatewayApiGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayApiGroupsResponse) SetBody(v *ListGatewayApiGroupsResponseBody) *ListGatewayApiGroupsResponse {
	s.Body = v
	return s
}

func (s *ListGatewayApiGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
