// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaaSAntChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaaSAntChainResponse
	GetStatusCode() *int32
	SetBody(v *ListBaaSAntChainResponseBody) *ListBaaSAntChainResponse
	GetBody() *ListBaaSAntChainResponseBody
}

type ListBaaSAntChainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaaSAntChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaaSAntChainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainResponse) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaaSAntChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaaSAntChainResponse) GetBody() *ListBaaSAntChainResponseBody {
	return s.Body
}

func (s *ListBaaSAntChainResponse) SetHeaders(v map[string]*string) *ListBaaSAntChainResponse {
	s.Headers = v
	return s
}

func (s *ListBaaSAntChainResponse) SetStatusCode(v int32) *ListBaaSAntChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaaSAntChainResponse) SetBody(v *ListBaaSAntChainResponseBody) *ListBaaSAntChainResponse {
	s.Body = v
	return s
}

func (s *ListBaaSAntChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
