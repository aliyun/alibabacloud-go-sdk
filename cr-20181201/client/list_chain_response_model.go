// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChainResponse
	GetStatusCode() *int32
	SetBody(v *ListChainResponseBody) *ListChainResponse
	GetBody() *ListChainResponseBody
}

type ListChainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChainResponse) GoString() string {
	return s.String()
}

func (s *ListChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChainResponse) GetBody() *ListChainResponseBody {
	return s.Body
}

func (s *ListChainResponse) SetHeaders(v map[string]*string) *ListChainResponse {
	s.Headers = v
	return s
}

func (s *ListChainResponse) SetStatusCode(v int32) *ListChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChainResponse) SetBody(v *ListChainResponseBody) *ListChainResponse {
	s.Body = v
	return s
}

func (s *ListChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
