// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAssetScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAssetScopeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAssetScopeResponseBody) *UpdateAssetScopeResponse
	GetBody() *UpdateAssetScopeResponseBody
}

type UpdateAssetScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAssetScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAssetScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAssetScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAssetScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAssetScopeResponse) GetBody() *UpdateAssetScopeResponseBody {
	return s.Body
}

func (s *UpdateAssetScopeResponse) SetHeaders(v map[string]*string) *UpdateAssetScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAssetScopeResponse) SetStatusCode(v int32) *UpdateAssetScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAssetScopeResponse) SetBody(v *UpdateAssetScopeResponseBody) *UpdateAssetScopeResponse {
	s.Body = v
	return s
}

func (s *UpdateAssetScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
