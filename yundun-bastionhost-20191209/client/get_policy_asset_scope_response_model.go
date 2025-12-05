// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyAssetScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyAssetScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyAssetScopeResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyAssetScopeResponseBody) *GetPolicyAssetScopeResponse
	GetBody() *GetPolicyAssetScopeResponseBody
}

type GetPolicyAssetScopeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyAssetScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyAssetScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyAssetScopeResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyAssetScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyAssetScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyAssetScopeResponse) GetBody() *GetPolicyAssetScopeResponseBody {
	return s.Body
}

func (s *GetPolicyAssetScopeResponse) SetHeaders(v map[string]*string) *GetPolicyAssetScopeResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyAssetScopeResponse) SetStatusCode(v int32) *GetPolicyAssetScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyAssetScopeResponse) SetBody(v *GetPolicyAssetScopeResponseBody) *GetPolicyAssetScopeResponse {
	s.Body = v
	return s
}

func (s *GetPolicyAssetScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
