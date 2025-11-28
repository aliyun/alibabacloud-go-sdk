// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaaSAntChainBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBaaSAntChainBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBaaSAntChainBizChainResponse
	GetStatusCode() *int32
	SetBody(v *AddBaaSAntChainBizChainResponseBody) *AddBaaSAntChainBizChainResponse
	GetBody() *AddBaaSAntChainBizChainResponseBody
}

type AddBaaSAntChainBizChainResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBaaSAntChainBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBaaSAntChainBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBaaSAntChainBizChainResponse) GoString() string {
	return s.String()
}

func (s *AddBaaSAntChainBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBaaSAntChainBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBaaSAntChainBizChainResponse) GetBody() *AddBaaSAntChainBizChainResponseBody {
	return s.Body
}

func (s *AddBaaSAntChainBizChainResponse) SetHeaders(v map[string]*string) *AddBaaSAntChainBizChainResponse {
	s.Headers = v
	return s
}

func (s *AddBaaSAntChainBizChainResponse) SetStatusCode(v int32) *AddBaaSAntChainBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBaaSAntChainBizChainResponse) SetBody(v *AddBaaSAntChainBizChainResponseBody) *AddBaaSAntChainBizChainResponse {
	s.Body = v
	return s
}

func (s *AddBaaSAntChainBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
