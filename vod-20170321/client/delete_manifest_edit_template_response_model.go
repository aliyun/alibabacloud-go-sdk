// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteManifestEditTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteManifestEditTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteManifestEditTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteManifestEditTemplateResponseBody) *DeleteManifestEditTemplateResponse
	GetBody() *DeleteManifestEditTemplateResponseBody
}

type DeleteManifestEditTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteManifestEditTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteManifestEditTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteManifestEditTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteManifestEditTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteManifestEditTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteManifestEditTemplateResponse) GetBody() *DeleteManifestEditTemplateResponseBody {
	return s.Body
}

func (s *DeleteManifestEditTemplateResponse) SetHeaders(v map[string]*string) *DeleteManifestEditTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteManifestEditTemplateResponse) SetStatusCode(v int32) *DeleteManifestEditTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteManifestEditTemplateResponse) SetBody(v *DeleteManifestEditTemplateResponseBody) *DeleteManifestEditTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteManifestEditTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
