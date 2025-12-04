// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDatasetVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewDatasetVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewDatasetVersionResponse
	GetStatusCode() *int32
	SetBody(v *PreviewDatasetVersionResponseBody) *PreviewDatasetVersionResponse
	GetBody() *PreviewDatasetVersionResponseBody
}

type PreviewDatasetVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewDatasetVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *PreviewDatasetVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewDatasetVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewDatasetVersionResponse) GetBody() *PreviewDatasetVersionResponseBody {
	return s.Body
}

func (s *PreviewDatasetVersionResponse) SetHeaders(v map[string]*string) *PreviewDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *PreviewDatasetVersionResponse) SetStatusCode(v int32) *PreviewDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewDatasetVersionResponse) SetBody(v *PreviewDatasetVersionResponseBody) *PreviewDatasetVersionResponse {
	s.Body = v
	return s
}

func (s *PreviewDatasetVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
