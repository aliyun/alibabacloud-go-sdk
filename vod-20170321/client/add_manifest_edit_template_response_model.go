// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddManifestEditTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddManifestEditTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddManifestEditTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddManifestEditTemplateResponseBody) *AddManifestEditTemplateResponse
	GetBody() *AddManifestEditTemplateResponseBody
}

type AddManifestEditTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddManifestEditTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddManifestEditTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddManifestEditTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddManifestEditTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddManifestEditTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddManifestEditTemplateResponse) GetBody() *AddManifestEditTemplateResponseBody {
	return s.Body
}

func (s *AddManifestEditTemplateResponse) SetHeaders(v map[string]*string) *AddManifestEditTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddManifestEditTemplateResponse) SetStatusCode(v int32) *AddManifestEditTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddManifestEditTemplateResponse) SetBody(v *AddManifestEditTemplateResponseBody) *AddManifestEditTemplateResponse {
	s.Body = v
	return s
}

func (s *AddManifestEditTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
