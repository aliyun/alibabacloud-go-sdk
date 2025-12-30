// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadMultimodalSearchTaskResultMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadMultimodalSearchTaskResultMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadMultimodalSearchTaskResultMetadataResponse
	GetStatusCode() *int32
	SetBody(v *DownloadMultimodalSearchTaskResultMetadataResponseBody) *DownloadMultimodalSearchTaskResultMetadataResponse
	GetBody() *DownloadMultimodalSearchTaskResultMetadataResponseBody
}

type DownloadMultimodalSearchTaskResultMetadataResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadMultimodalSearchTaskResultMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadMultimodalSearchTaskResultMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadMultimodalSearchTaskResultMetadataResponse) GoString() string {
	return s.String()
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponse) GetBody() *DownloadMultimodalSearchTaskResultMetadataResponseBody {
	return s.Body
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponse) SetHeaders(v map[string]*string) *DownloadMultimodalSearchTaskResultMetadataResponse {
	s.Headers = v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponse) SetStatusCode(v int32) *DownloadMultimodalSearchTaskResultMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponse) SetBody(v *DownloadMultimodalSearchTaskResultMetadataResponseBody) *DownloadMultimodalSearchTaskResultMetadataResponse {
	s.Body = v
	return s
}

func (s *DownloadMultimodalSearchTaskResultMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
