// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaaSFabricBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBaaSFabricBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBaaSFabricBizChainResponse
	GetStatusCode() *int32
	SetBody(v *AddBaaSFabricBizChainResponseBody) *AddBaaSFabricBizChainResponse
	GetBody() *AddBaaSFabricBizChainResponseBody
}

type AddBaaSFabricBizChainResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBaaSFabricBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBaaSFabricBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBaaSFabricBizChainResponse) GoString() string {
	return s.String()
}

func (s *AddBaaSFabricBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBaaSFabricBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBaaSFabricBizChainResponse) GetBody() *AddBaaSFabricBizChainResponseBody {
	return s.Body
}

func (s *AddBaaSFabricBizChainResponse) SetHeaders(v map[string]*string) *AddBaaSFabricBizChainResponse {
	s.Headers = v
	return s
}

func (s *AddBaaSFabricBizChainResponse) SetStatusCode(v int32) *AddBaaSFabricBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBaaSFabricBizChainResponse) SetBody(v *AddBaaSFabricBizChainResponseBody) *AddBaaSFabricBizChainResponse {
	s.Body = v
	return s
}

func (s *AddBaaSFabricBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
