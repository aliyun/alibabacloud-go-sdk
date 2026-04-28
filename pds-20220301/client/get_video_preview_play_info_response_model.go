// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPreviewPlayInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoPreviewPlayInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoPreviewPlayInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoPreviewPlayInfoResponseBody) *GetVideoPreviewPlayInfoResponse
	GetBody() *GetVideoPreviewPlayInfoResponseBody
}

type GetVideoPreviewPlayInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoPreviewPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoPreviewPlayInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPreviewPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoPreviewPlayInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoPreviewPlayInfoResponse) GetBody() *GetVideoPreviewPlayInfoResponseBody {
	return s.Body
}

func (s *GetVideoPreviewPlayInfoResponse) SetHeaders(v map[string]*string) *GetVideoPreviewPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoPreviewPlayInfoResponse) SetStatusCode(v int32) *GetVideoPreviewPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponse) SetBody(v *GetVideoPreviewPlayInfoResponseBody) *GetVideoPreviewPlayInfoResponse {
	s.Body = v
	return s
}

func (s *GetVideoPreviewPlayInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
