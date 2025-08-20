// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateScratchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTemplateScratchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTemplateScratchResponse
	GetStatusCode() *int32
	SetBody(v *CreateTemplateScratchResponseBody) *CreateTemplateScratchResponse
	GetBody() *CreateTemplateScratchResponseBody
}

type CreateTemplateScratchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateScratchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTemplateScratchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTemplateScratchResponse) GetBody() *CreateTemplateScratchResponseBody {
	return s.Body
}

func (s *CreateTemplateScratchResponse) SetHeaders(v map[string]*string) *CreateTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateScratchResponse) SetStatusCode(v int32) *CreateTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateScratchResponse) SetBody(v *CreateTemplateScratchResponseBody) *CreateTemplateScratchResponse {
	s.Body = v
	return s
}

func (s *CreateTemplateScratchResponse) Validate() error {
	return dara.Validate(s)
}
