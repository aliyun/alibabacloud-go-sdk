// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTemplatePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetTemplatePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetTemplatePermissionResponse
	GetStatusCode() *int32
	SetBody(v *SetTemplatePermissionResponseBody) *SetTemplatePermissionResponse
	GetBody() *SetTemplatePermissionResponseBody
}

type SetTemplatePermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTemplatePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTemplatePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetTemplatePermissionResponse) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetTemplatePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetTemplatePermissionResponse) GetBody() *SetTemplatePermissionResponseBody {
	return s.Body
}

func (s *SetTemplatePermissionResponse) SetHeaders(v map[string]*string) *SetTemplatePermissionResponse {
	s.Headers = v
	return s
}

func (s *SetTemplatePermissionResponse) SetStatusCode(v int32) *SetTemplatePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTemplatePermissionResponse) SetBody(v *SetTemplatePermissionResponseBody) *SetTemplatePermissionResponse {
	s.Body = v
	return s
}

func (s *SetTemplatePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
