// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainConsortiumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaaSAntChainConsortiumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaaSAntChainConsortiumResponse
	GetStatusCode() *int32
	SetBody(v *ListBaaSAntChainConsortiumResponseBody) *ListBaaSAntChainConsortiumResponse
	GetBody() *ListBaaSAntChainConsortiumResponseBody
}

type ListBaaSAntChainConsortiumResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaaSAntChainConsortiumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaaSAntChainConsortiumResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainConsortiumResponse) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainConsortiumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaaSAntChainConsortiumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaaSAntChainConsortiumResponse) GetBody() *ListBaaSAntChainConsortiumResponseBody {
	return s.Body
}

func (s *ListBaaSAntChainConsortiumResponse) SetHeaders(v map[string]*string) *ListBaaSAntChainConsortiumResponse {
	s.Headers = v
	return s
}

func (s *ListBaaSAntChainConsortiumResponse) SetStatusCode(v int32) *ListBaaSAntChainConsortiumResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaaSAntChainConsortiumResponse) SetBody(v *ListBaaSAntChainConsortiumResponseBody) *ListBaaSAntChainConsortiumResponse {
	s.Body = v
	return s
}

func (s *ListBaaSAntChainConsortiumResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
