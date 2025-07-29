// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourcesDeleteProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourcesDeleteProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourcesDeleteProtectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourcesDeleteProtectionResponseBody) *UpdateResourcesDeleteProtectionResponse
	GetBody() *UpdateResourcesDeleteProtectionResponseBody
}

type UpdateResourcesDeleteProtectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourcesDeleteProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourcesDeleteProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourcesDeleteProtectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourcesDeleteProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourcesDeleteProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourcesDeleteProtectionResponse) GetBody() *UpdateResourcesDeleteProtectionResponseBody {
	return s.Body
}

func (s *UpdateResourcesDeleteProtectionResponse) SetHeaders(v map[string]*string) *UpdateResourcesDeleteProtectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponse) SetStatusCode(v int32) *UpdateResourcesDeleteProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponse) SetBody(v *UpdateResourcesDeleteProtectionResponseBody) *UpdateResourcesDeleteProtectionResponse {
	s.Body = v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponse) Validate() error {
	return dara.Validate(s)
}
