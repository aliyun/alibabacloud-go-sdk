// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneEffectiveScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionZoneEffectiveScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionZoneEffectiveScopeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionZoneEffectiveScopeResponseBody) *UpdateRecursionZoneEffectiveScopeResponse
	GetBody() *UpdateRecursionZoneEffectiveScopeResponseBody
}

type UpdateRecursionZoneEffectiveScopeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionZoneEffectiveScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionZoneEffectiveScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneEffectiveScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneEffectiveScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionZoneEffectiveScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionZoneEffectiveScopeResponse) GetBody() *UpdateRecursionZoneEffectiveScopeResponseBody {
	return s.Body
}

func (s *UpdateRecursionZoneEffectiveScopeResponse) SetHeaders(v map[string]*string) *UpdateRecursionZoneEffectiveScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeResponse) SetStatusCode(v int32) *UpdateRecursionZoneEffectiveScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeResponse) SetBody(v *UpdateRecursionZoneEffectiveScopeResponseBody) *UpdateRecursionZoneEffectiveScopeResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeResponse) Validate() error {
	return dara.Validate(s)
}
