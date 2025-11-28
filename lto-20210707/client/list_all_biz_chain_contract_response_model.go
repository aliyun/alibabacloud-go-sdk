// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllBizChainContractResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllBizChainContractResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllBizChainContractResponse
	GetStatusCode() *int32
	SetBody(v *ListAllBizChainContractResponseBody) *ListAllBizChainContractResponse
	GetBody() *ListAllBizChainContractResponseBody
}

type ListAllBizChainContractResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllBizChainContractResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllBizChainContractResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllBizChainContractResponse) GoString() string {
	return s.String()
}

func (s *ListAllBizChainContractResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllBizChainContractResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllBizChainContractResponse) GetBody() *ListAllBizChainContractResponseBody {
	return s.Body
}

func (s *ListAllBizChainContractResponse) SetHeaders(v map[string]*string) *ListAllBizChainContractResponse {
	s.Headers = v
	return s
}

func (s *ListAllBizChainContractResponse) SetStatusCode(v int32) *ListAllBizChainContractResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllBizChainContractResponse) SetBody(v *ListAllBizChainContractResponseBody) *ListAllBizChainContractResponse {
	s.Body = v
	return s
}

func (s *ListAllBizChainContractResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
