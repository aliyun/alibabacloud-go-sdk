// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sClusterUserConfigExpireResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sClusterUserConfigExpireResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sClusterUserConfigExpireResponse
	GetStatusCode() *int32
}

type UpdateK8sClusterUserConfigExpireResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateK8sClusterUserConfigExpireResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sClusterUserConfigExpireResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sClusterUserConfigExpireResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sClusterUserConfigExpireResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sClusterUserConfigExpireResponse) SetHeaders(v map[string]*string) *UpdateK8sClusterUserConfigExpireResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sClusterUserConfigExpireResponse) SetStatusCode(v int32) *UpdateK8sClusterUserConfigExpireResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sClusterUserConfigExpireResponse) Validate() error {
	return dara.Validate(s)
}
