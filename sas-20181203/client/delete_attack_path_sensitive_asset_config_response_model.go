// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackPathSensitiveAssetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAttackPathSensitiveAssetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAttackPathSensitiveAssetConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAttackPathSensitiveAssetConfigResponseBody) *DeleteAttackPathSensitiveAssetConfigResponse
	GetBody() *DeleteAttackPathSensitiveAssetConfigResponseBody
}

type DeleteAttackPathSensitiveAssetConfigResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAttackPathSensitiveAssetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAttackPathSensitiveAssetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackPathSensitiveAssetConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAttackPathSensitiveAssetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAttackPathSensitiveAssetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAttackPathSensitiveAssetConfigResponse) GetBody() *DeleteAttackPathSensitiveAssetConfigResponseBody {
	return s.Body
}

func (s *DeleteAttackPathSensitiveAssetConfigResponse) SetHeaders(v map[string]*string) *DeleteAttackPathSensitiveAssetConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAttackPathSensitiveAssetConfigResponse) SetStatusCode(v int32) *DeleteAttackPathSensitiveAssetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAttackPathSensitiveAssetConfigResponse) SetBody(v *DeleteAttackPathSensitiveAssetConfigResponseBody) *DeleteAttackPathSensitiveAssetConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteAttackPathSensitiveAssetConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
