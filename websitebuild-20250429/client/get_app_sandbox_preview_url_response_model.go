// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSandboxPreviewUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSandboxPreviewUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSandboxPreviewUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSandboxPreviewUrlResponseBody) *GetAppSandboxPreviewUrlResponse
	GetBody() *GetAppSandboxPreviewUrlResponseBody
}

type GetAppSandboxPreviewUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSandboxPreviewUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSandboxPreviewUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSandboxPreviewUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAppSandboxPreviewUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSandboxPreviewUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSandboxPreviewUrlResponse) GetBody() *GetAppSandboxPreviewUrlResponseBody {
	return s.Body
}

func (s *GetAppSandboxPreviewUrlResponse) SetHeaders(v map[string]*string) *GetAppSandboxPreviewUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAppSandboxPreviewUrlResponse) SetStatusCode(v int32) *GetAppSandboxPreviewUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSandboxPreviewUrlResponse) SetBody(v *GetAppSandboxPreviewUrlResponseBody) *GetAppSandboxPreviewUrlResponse {
	s.Body = v
	return s
}

func (s *GetAppSandboxPreviewUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
