// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManifestEditTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManifestEditTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManifestEditTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListManifestEditTemplateResponseBody) *ListManifestEditTemplateResponse
	GetBody() *ListManifestEditTemplateResponseBody
}

type ListManifestEditTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManifestEditTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManifestEditTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManifestEditTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListManifestEditTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManifestEditTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManifestEditTemplateResponse) GetBody() *ListManifestEditTemplateResponseBody {
	return s.Body
}

func (s *ListManifestEditTemplateResponse) SetHeaders(v map[string]*string) *ListManifestEditTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListManifestEditTemplateResponse) SetStatusCode(v int32) *ListManifestEditTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManifestEditTemplateResponse) SetBody(v *ListManifestEditTemplateResponseBody) *ListManifestEditTemplateResponse {
	s.Body = v
	return s
}

func (s *ListManifestEditTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
