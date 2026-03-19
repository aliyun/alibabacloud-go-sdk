// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptArtifactExportResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPptArtifactExportResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPptArtifactExportResultResponse
	GetStatusCode() *int32
	SetBody(v *GetPptArtifactExportResultResponseBody) *GetPptArtifactExportResultResponse
	GetBody() *GetPptArtifactExportResultResponseBody
}

type GetPptArtifactExportResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPptArtifactExportResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPptArtifactExportResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactExportResultResponse) GoString() string {
	return s.String()
}

func (s *GetPptArtifactExportResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPptArtifactExportResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPptArtifactExportResultResponse) GetBody() *GetPptArtifactExportResultResponseBody {
	return s.Body
}

func (s *GetPptArtifactExportResultResponse) SetHeaders(v map[string]*string) *GetPptArtifactExportResultResponse {
	s.Headers = v
	return s
}

func (s *GetPptArtifactExportResultResponse) SetStatusCode(v int32) *GetPptArtifactExportResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPptArtifactExportResultResponse) SetBody(v *GetPptArtifactExportResultResponseBody) *GetPptArtifactExportResultResponse {
	s.Body = v
	return s
}

func (s *GetPptArtifactExportResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
