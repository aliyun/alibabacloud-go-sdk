// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBsnFabricBizChainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBsnFabricBizChainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBsnFabricBizChainResponse
	GetStatusCode() *int32
	SetBody(v *AddBsnFabricBizChainResponseBody) *AddBsnFabricBizChainResponse
	GetBody() *AddBsnFabricBizChainResponseBody
}

type AddBsnFabricBizChainResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBsnFabricBizChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBsnFabricBizChainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBsnFabricBizChainResponse) GoString() string {
	return s.String()
}

func (s *AddBsnFabricBizChainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBsnFabricBizChainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBsnFabricBizChainResponse) GetBody() *AddBsnFabricBizChainResponseBody {
	return s.Body
}

func (s *AddBsnFabricBizChainResponse) SetHeaders(v map[string]*string) *AddBsnFabricBizChainResponse {
	s.Headers = v
	return s
}

func (s *AddBsnFabricBizChainResponse) SetStatusCode(v int32) *AddBsnFabricBizChainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBsnFabricBizChainResponse) SetBody(v *AddBsnFabricBizChainResponseBody) *AddBsnFabricBizChainResponse {
	s.Body = v
	return s
}

func (s *AddBsnFabricBizChainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
