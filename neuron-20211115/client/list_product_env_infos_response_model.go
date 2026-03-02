// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductEnvInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductEnvInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductEnvInfosResponse
	GetStatusCode() *int32
	SetBody(v *EnvListResult) *ListProductEnvInfosResponse
	GetBody() *EnvListResult
}

type ListProductEnvInfosResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnvListResult     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductEnvInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductEnvInfosResponse) GoString() string {
	return s.String()
}

func (s *ListProductEnvInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductEnvInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductEnvInfosResponse) GetBody() *EnvListResult {
	return s.Body
}

func (s *ListProductEnvInfosResponse) SetHeaders(v map[string]*string) *ListProductEnvInfosResponse {
	s.Headers = v
	return s
}

func (s *ListProductEnvInfosResponse) SetStatusCode(v int32) *ListProductEnvInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductEnvInfosResponse) SetBody(v *EnvListResult) *ListProductEnvInfosResponse {
	s.Body = v
	return s
}

func (s *ListProductEnvInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
