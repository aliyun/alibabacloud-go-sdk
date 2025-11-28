// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllBizChainResponse
	GetStatusCode() *int32
	SetBody(v *ListAllBizChainResponseBody) *ListAllBizChainResponse
	GetBody() *ListAllBizChainResponseBody
}

type ListAllBizChainResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainResponse) GoString() string {
	return s.String()
}

func (s *ListAllBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllBizChainResponse) GetBody() *ListAllBizChainResponseBody {
	return s.Body
}

func (s *ListAllBizChainResponse) SetHeaders(v map[string]*string) *ListAllBizChainResponse {
	s.Headers = v
	return s
}

func (s *ListAllBizChainResponse) SetStatusCode(v int32) *ListAllBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllBizChainResponse) SetBody(v *ListAllBizChainResponseBody) *ListAllBizChainResponse {
	s.Body = v
	return s
}

func (s *ListAllBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
