// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCleanClusterUserPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CleanClusterUserPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CleanClusterUserPermissionsResponse
	GetStatusCode() *int32
}

type CleanClusterUserPermissionsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CleanClusterUserPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CleanClusterUserPermissionsResponse) GoString() string {
	return s.String()
}

func (s *CleanClusterUserPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CleanClusterUserPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CleanClusterUserPermissionsResponse) SetHeaders(v map[string]*string) *CleanClusterUserPermissionsResponse {
	s.Headers = v
	return s
}

func (s *CleanClusterUserPermissionsResponse) SetStatusCode(v int32) *CleanClusterUserPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CleanClusterUserPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
