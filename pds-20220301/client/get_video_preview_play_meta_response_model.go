// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPreviewPlayMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoPreviewPlayMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoPreviewPlayMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoPreviewPlayMetaResponseBody) *GetVideoPreviewPlayMetaResponse
	GetBody() *GetVideoPreviewPlayMetaResponseBody
}

type GetVideoPreviewPlayMetaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoPreviewPlayMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoPreviewPlayMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPreviewPlayMetaResponse) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoPreviewPlayMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoPreviewPlayMetaResponse) GetBody() *GetVideoPreviewPlayMetaResponseBody {
	return s.Body
}

func (s *GetVideoPreviewPlayMetaResponse) SetHeaders(v map[string]*string) *GetVideoPreviewPlayMetaResponse {
	s.Headers = v
	return s
}

func (s *GetVideoPreviewPlayMetaResponse) SetStatusCode(v int32) *GetVideoPreviewPlayMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponse) SetBody(v *GetVideoPreviewPlayMetaResponseBody) *GetVideoPreviewPlayMetaResponse {
	s.Body = v
	return s
}

func (s *GetVideoPreviewPlayMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
