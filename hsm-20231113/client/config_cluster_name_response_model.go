// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigClusterNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigClusterNameResponse
	GetStatusCode() *int32
	SetBody(v *ConfigClusterNameResponseBody) *ConfigClusterNameResponse
	GetBody() *ConfigClusterNameResponseBody
}

type ConfigClusterNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigClusterNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigClusterNameResponse) GetBody() *ConfigClusterNameResponseBody {
	return s.Body
}

func (s *ConfigClusterNameResponse) SetHeaders(v map[string]*string) *ConfigClusterNameResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterNameResponse) SetStatusCode(v int32) *ConfigClusterNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterNameResponse) SetBody(v *ConfigClusterNameResponseBody) *ConfigClusterNameResponse {
	s.Body = v
	return s
}

func (s *ConfigClusterNameResponse) Validate() error {
	return dara.Validate(s)
}
