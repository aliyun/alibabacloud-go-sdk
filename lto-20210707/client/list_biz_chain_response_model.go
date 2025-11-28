// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBizChainResponse
	GetStatusCode() *int32
	SetBody(v *ListBizChainResponseBody) *ListBizChainResponse
	GetBody() *ListBizChainResponseBody
}

type ListBizChainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainResponse) GoString() string {
	return s.String()
}

func (s *ListBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBizChainResponse) GetBody() *ListBizChainResponseBody {
	return s.Body
}

func (s *ListBizChainResponse) SetHeaders(v map[string]*string) *ListBizChainResponse {
	s.Headers = v
	return s
}

func (s *ListBizChainResponse) SetStatusCode(v int32) *ListBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBizChainResponse) SetBody(v *ListBizChainResponseBody) *ListBizChainResponse {
	s.Body = v
	return s
}

func (s *ListBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
