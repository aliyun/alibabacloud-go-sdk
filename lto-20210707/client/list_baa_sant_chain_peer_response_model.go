// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSAntChainPeerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaaSAntChainPeerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaaSAntChainPeerResponse
	GetStatusCode() *int32
	SetBody(v *ListBaaSAntChainPeerResponseBody) *ListBaaSAntChainPeerResponse
	GetBody() *ListBaaSAntChainPeerResponseBody
}

type ListBaaSAntChainPeerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaaSAntChainPeerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaaSAntChainPeerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSAntChainPeerResponse) GoString() string {
	return s.String()
}

func (s *ListBaaSAntChainPeerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaaSAntChainPeerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaaSAntChainPeerResponse) GetBody() *ListBaaSAntChainPeerResponseBody {
	return s.Body
}

func (s *ListBaaSAntChainPeerResponse) SetHeaders(v map[string]*string) *ListBaaSAntChainPeerResponse {
	s.Headers = v
	return s
}

func (s *ListBaaSAntChainPeerResponse) SetStatusCode(v int32) *ListBaaSAntChainPeerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaaSAntChainPeerResponse) SetBody(v *ListBaaSAntChainPeerResponseBody) *ListBaaSAntChainPeerResponse {
	s.Body = v
	return s
}

func (s *ListBaaSAntChainPeerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
