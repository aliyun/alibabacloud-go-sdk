// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterConnectionTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterConnectionTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterConnectionTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterConnectionTypesResponseBody) *ListClusterConnectionTypesResponse
	GetBody() *ListClusterConnectionTypesResponseBody
}

type ListClusterConnectionTypesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterConnectionTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterConnectionTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterConnectionTypesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterConnectionTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterConnectionTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterConnectionTypesResponse) GetBody() *ListClusterConnectionTypesResponseBody {
	return s.Body
}

func (s *ListClusterConnectionTypesResponse) SetHeaders(v map[string]*string) *ListClusterConnectionTypesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterConnectionTypesResponse) SetStatusCode(v int32) *ListClusterConnectionTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterConnectionTypesResponse) SetBody(v *ListClusterConnectionTypesResponseBody) *ListClusterConnectionTypesResponse {
	s.Body = v
	return s
}

func (s *ListClusterConnectionTypesResponse) Validate() error {
	return dara.Validate(s)
}
