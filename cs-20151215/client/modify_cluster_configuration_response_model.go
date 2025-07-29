// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterConfigurationResponse
	GetStatusCode() *int32
}

type ModifyClusterConfigurationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s ModifyClusterConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterConfigurationResponse) SetHeaders(v map[string]*string) *ModifyClusterConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterConfigurationResponse) SetStatusCode(v int32) *ModifyClusterConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
