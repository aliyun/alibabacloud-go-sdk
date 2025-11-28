// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizChainDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBizChainDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBizChainDataResponse
	GetStatusCode() *int32
	SetBody(v *ListBizChainDataResponseBody) *ListBizChainDataResponse
	GetBody() *ListBizChainDataResponseBody
}

type ListBizChainDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBizChainDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBizChainDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainDataResponse) GoString() string {
	return s.String()
}

func (s *ListBizChainDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBizChainDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBizChainDataResponse) GetBody() *ListBizChainDataResponseBody {
	return s.Body
}

func (s *ListBizChainDataResponse) SetHeaders(v map[string]*string) *ListBizChainDataResponse {
	s.Headers = v
	return s
}

func (s *ListBizChainDataResponse) SetStatusCode(v int32) *ListBizChainDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBizChainDataResponse) SetBody(v *ListBizChainDataResponseBody) *ListBizChainDataResponse {
	s.Body = v
	return s
}

func (s *ListBizChainDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
