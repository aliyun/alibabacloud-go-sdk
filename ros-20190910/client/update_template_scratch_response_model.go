// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateScratchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTemplateScratchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTemplateScratchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTemplateScratchResponseBody) *UpdateTemplateScratchResponse
	GetBody() *UpdateTemplateScratchResponseBody
}

type UpdateTemplateScratchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateScratchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTemplateScratchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTemplateScratchResponse) GetBody() *UpdateTemplateScratchResponseBody {
	return s.Body
}

func (s *UpdateTemplateScratchResponse) SetHeaders(v map[string]*string) *UpdateTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateScratchResponse) SetStatusCode(v int32) *UpdateTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateScratchResponse) SetBody(v *UpdateTemplateScratchResponseBody) *UpdateTemplateScratchResponse {
	s.Body = v
	return s
}

func (s *UpdateTemplateScratchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
