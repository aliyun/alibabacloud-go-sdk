// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvInfosResponse
	GetStatusCode() *int32
	SetBody(v *PdpListEnvInfoResult) *ListEnvInfosResponse
	GetBody() *PdpListEnvInfoResult
}

type ListEnvInfosResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpListEnvInfoResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvInfosResponse) GoString() string {
	return s.String()
}

func (s *ListEnvInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvInfosResponse) GetBody() *PdpListEnvInfoResult {
	return s.Body
}

func (s *ListEnvInfosResponse) SetHeaders(v map[string]*string) *ListEnvInfosResponse {
	s.Headers = v
	return s
}

func (s *ListEnvInfosResponse) SetStatusCode(v int32) *ListEnvInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvInfosResponse) SetBody(v *PdpListEnvInfoResult) *ListEnvInfosResponse {
	s.Body = v
	return s
}

func (s *ListEnvInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
