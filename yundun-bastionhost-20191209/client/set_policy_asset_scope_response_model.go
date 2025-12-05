// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyAssetScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolicyAssetScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolicyAssetScopeResponse
	GetStatusCode() *int32
	SetBody(v *SetPolicyAssetScopeResponseBody) *SetPolicyAssetScopeResponse
	GetBody() *SetPolicyAssetScopeResponseBody
}

type SetPolicyAssetScopeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolicyAssetScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolicyAssetScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyAssetScopeResponse) GoString() string {
	return s.String()
}

func (s *SetPolicyAssetScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolicyAssetScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolicyAssetScopeResponse) GetBody() *SetPolicyAssetScopeResponseBody {
	return s.Body
}

func (s *SetPolicyAssetScopeResponse) SetHeaders(v map[string]*string) *SetPolicyAssetScopeResponse {
	s.Headers = v
	return s
}

func (s *SetPolicyAssetScopeResponse) SetStatusCode(v int32) *SetPolicyAssetScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolicyAssetScopeResponse) SetBody(v *SetPolicyAssetScopeResponseBody) *SetPolicyAssetScopeResponse {
	s.Body = v
	return s
}

func (s *SetPolicyAssetScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
