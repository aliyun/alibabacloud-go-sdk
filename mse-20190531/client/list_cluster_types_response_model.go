// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterTypesResponseBody) *ListClusterTypesResponse
	GetBody() *ListClusterTypesResponseBody
}

type ListClusterTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTypesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterTypesResponse) GetBody() *ListClusterTypesResponseBody {
	return s.Body
}

func (s *ListClusterTypesResponse) SetHeaders(v map[string]*string) *ListClusterTypesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTypesResponse) SetStatusCode(v int32) *ListClusterTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterTypesResponse) SetBody(v *ListClusterTypesResponseBody) *ListClusterTypesResponse {
	s.Body = v
	return s
}

func (s *ListClusterTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
