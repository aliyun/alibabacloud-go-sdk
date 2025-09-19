// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttackPathSensitiveAssetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAttackPathSensitiveAssetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAttackPathSensitiveAssetConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateAttackPathSensitiveAssetConfigResponseBody) *CreateAttackPathSensitiveAssetConfigResponse
	GetBody() *CreateAttackPathSensitiveAssetConfigResponseBody
}

type CreateAttackPathSensitiveAssetConfigResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAttackPathSensitiveAssetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAttackPathSensitiveAssetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathSensitiveAssetConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAttackPathSensitiveAssetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAttackPathSensitiveAssetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAttackPathSensitiveAssetConfigResponse) GetBody() *CreateAttackPathSensitiveAssetConfigResponseBody {
	return s.Body
}

func (s *CreateAttackPathSensitiveAssetConfigResponse) SetHeaders(v map[string]*string) *CreateAttackPathSensitiveAssetConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigResponse) SetStatusCode(v int32) *CreateAttackPathSensitiveAssetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigResponse) SetBody(v *CreateAttackPathSensitiveAssetConfigResponseBody) *CreateAttackPathSensitiveAssetConfigResponse {
	s.Body = v
	return s
}

func (s *CreateAttackPathSensitiveAssetConfigResponse) Validate() error {
	return dara.Validate(s)
}
