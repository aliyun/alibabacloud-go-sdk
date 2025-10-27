// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathSensitiveAssetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAttackPathSensitiveAssetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAttackPathSensitiveAssetConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAttackPathSensitiveAssetConfigResponseBody) *GetAttackPathSensitiveAssetConfigResponse
	GetBody() *GetAttackPathSensitiveAssetConfigResponseBody
}

type GetAttackPathSensitiveAssetConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttackPathSensitiveAssetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttackPathSensitiveAssetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathSensitiveAssetConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAttackPathSensitiveAssetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAttackPathSensitiveAssetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAttackPathSensitiveAssetConfigResponse) GetBody() *GetAttackPathSensitiveAssetConfigResponseBody {
	return s.Body
}

func (s *GetAttackPathSensitiveAssetConfigResponse) SetHeaders(v map[string]*string) *GetAttackPathSensitiveAssetConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAttackPathSensitiveAssetConfigResponse) SetStatusCode(v int32) *GetAttackPathSensitiveAssetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttackPathSensitiveAssetConfigResponse) SetBody(v *GetAttackPathSensitiveAssetConfigResponseBody) *GetAttackPathSensitiveAssetConfigResponse {
	s.Body = v
	return s
}

func (s *GetAttackPathSensitiveAssetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
