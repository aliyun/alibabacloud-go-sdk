// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricOrganizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaaSFabricOrganizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaaSFabricOrganizationResponse
	GetStatusCode() *int32
	SetBody(v *ListBaaSFabricOrganizationResponseBody) *ListBaaSFabricOrganizationResponse
	GetBody() *ListBaaSFabricOrganizationResponseBody
}

type ListBaaSFabricOrganizationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaaSFabricOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaaSFabricOrganizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricOrganizationResponse) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricOrganizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaaSFabricOrganizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaaSFabricOrganizationResponse) GetBody() *ListBaaSFabricOrganizationResponseBody {
	return s.Body
}

func (s *ListBaaSFabricOrganizationResponse) SetHeaders(v map[string]*string) *ListBaaSFabricOrganizationResponse {
	s.Headers = v
	return s
}

func (s *ListBaaSFabricOrganizationResponse) SetStatusCode(v int32) *ListBaaSFabricOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaaSFabricOrganizationResponse) SetBody(v *ListBaaSFabricOrganizationResponseBody) *ListBaaSFabricOrganizationResponse {
	s.Body = v
	return s
}

func (s *ListBaaSFabricOrganizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
