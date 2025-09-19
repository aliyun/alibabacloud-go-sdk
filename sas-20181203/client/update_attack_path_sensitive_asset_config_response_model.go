// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttackPathSensitiveAssetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAttackPathSensitiveAssetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAttackPathSensitiveAssetConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAttackPathSensitiveAssetConfigResponseBody) *UpdateAttackPathSensitiveAssetConfigResponse
	GetBody() *UpdateAttackPathSensitiveAssetConfigResponseBody
}

type UpdateAttackPathSensitiveAssetConfigResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAttackPathSensitiveAssetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAttackPathSensitiveAssetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttackPathSensitiveAssetConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateAttackPathSensitiveAssetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAttackPathSensitiveAssetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAttackPathSensitiveAssetConfigResponse) GetBody() *UpdateAttackPathSensitiveAssetConfigResponseBody {
	return s.Body
}

func (s *UpdateAttackPathSensitiveAssetConfigResponse) SetHeaders(v map[string]*string) *UpdateAttackPathSensitiveAssetConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigResponse) SetStatusCode(v int32) *UpdateAttackPathSensitiveAssetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigResponse) SetBody(v *UpdateAttackPathSensitiveAssetConfigResponseBody) *UpdateAttackPathSensitiveAssetConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateAttackPathSensitiveAssetConfigResponse) Validate() error {
	return dara.Validate(s)
}
