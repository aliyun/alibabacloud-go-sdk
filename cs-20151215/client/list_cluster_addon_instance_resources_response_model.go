// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterAddonInstanceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterAddonInstanceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterAddonInstanceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterAddonInstanceResourcesResponseBody) *ListClusterAddonInstanceResourcesResponse
	GetBody() *ListClusterAddonInstanceResourcesResponseBody
}

type ListClusterAddonInstanceResourcesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterAddonInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterAddonInstanceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterAddonInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterAddonInstanceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterAddonInstanceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterAddonInstanceResourcesResponse) GetBody() *ListClusterAddonInstanceResourcesResponseBody {
	return s.Body
}

func (s *ListClusterAddonInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListClusterAddonInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponse) SetStatusCode(v int32) *ListClusterAddonInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponse) SetBody(v *ListClusterAddonInstanceResourcesResponseBody) *ListClusterAddonInstanceResourcesResponse {
	s.Body = v
	return s
}

func (s *ListClusterAddonInstanceResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
