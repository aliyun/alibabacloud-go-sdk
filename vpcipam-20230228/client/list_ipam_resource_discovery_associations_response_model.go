// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceDiscoveryAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamResourceDiscoveryAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamResourceDiscoveryAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamResourceDiscoveryAssociationsResponseBody) *ListIpamResourceDiscoveryAssociationsResponse
	GetBody() *ListIpamResourceDiscoveryAssociationsResponseBody
}

type ListIpamResourceDiscoveryAssociationsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamResourceDiscoveryAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamResourceDiscoveryAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) GetBody() *ListIpamResourceDiscoveryAssociationsResponseBody {
	return s.Body
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) SetHeaders(v map[string]*string) *ListIpamResourceDiscoveryAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) SetStatusCode(v int32) *ListIpamResourceDiscoveryAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) SetBody(v *ListIpamResourceDiscoveryAssociationsResponseBody) *ListIpamResourceDiscoveryAssociationsResponse {
	s.Body = v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
