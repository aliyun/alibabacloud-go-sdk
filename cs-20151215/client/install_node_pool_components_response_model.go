// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallNodePoolComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallNodePoolComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallNodePoolComponentsResponse
	GetStatusCode() *int32
	SetBody(v *InstallNodePoolComponentsResponseBody) *InstallNodePoolComponentsResponse
	GetBody() *InstallNodePoolComponentsResponseBody
}

type InstallNodePoolComponentsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallNodePoolComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallNodePoolComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsResponse) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallNodePoolComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallNodePoolComponentsResponse) GetBody() *InstallNodePoolComponentsResponseBody {
	return s.Body
}

func (s *InstallNodePoolComponentsResponse) SetHeaders(v map[string]*string) *InstallNodePoolComponentsResponse {
	s.Headers = v
	return s
}

func (s *InstallNodePoolComponentsResponse) SetStatusCode(v int32) *InstallNodePoolComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallNodePoolComponentsResponse) SetBody(v *InstallNodePoolComponentsResponseBody) *InstallNodePoolComponentsResponse {
	s.Body = v
	return s
}

func (s *InstallNodePoolComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
