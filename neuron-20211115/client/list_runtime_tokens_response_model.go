// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRuntimeTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRuntimeTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRuntimeTokensResponse
	GetStatusCode() *int32
	SetBody(v *PdpListTokenResult) *ListRuntimeTokensResponse
	GetBody() *PdpListTokenResult
}

type ListRuntimeTokensResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpListTokenResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRuntimeTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRuntimeTokensResponse) GoString() string {
	return s.String()
}

func (s *ListRuntimeTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRuntimeTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRuntimeTokensResponse) GetBody() *PdpListTokenResult {
	return s.Body
}

func (s *ListRuntimeTokensResponse) SetHeaders(v map[string]*string) *ListRuntimeTokensResponse {
	s.Headers = v
	return s
}

func (s *ListRuntimeTokensResponse) SetStatusCode(v int32) *ListRuntimeTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRuntimeTokensResponse) SetBody(v *PdpListTokenResult) *ListRuntimeTokensResponse {
	s.Body = v
	return s
}

func (s *ListRuntimeTokensResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
