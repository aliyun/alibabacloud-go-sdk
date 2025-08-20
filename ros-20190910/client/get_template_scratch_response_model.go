// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateScratchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemplateScratchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemplateScratchResponse
	GetStatusCode() *int32
	SetBody(v *GetTemplateScratchResponseBody) *GetTemplateScratchResponse
	GetBody() *GetTemplateScratchResponseBody
}

type GetTemplateScratchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateScratchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemplateScratchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemplateScratchResponse) GetBody() *GetTemplateScratchResponseBody {
	return s.Body
}

func (s *GetTemplateScratchResponse) SetHeaders(v map[string]*string) *GetTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateScratchResponse) SetStatusCode(v int32) *GetTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateScratchResponse) SetBody(v *GetTemplateScratchResponseBody) *GetTemplateScratchResponse {
	s.Body = v
	return s
}

func (s *GetTemplateScratchResponse) Validate() error {
	return dara.Validate(s)
}
