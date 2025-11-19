// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshUploadVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshUploadVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshUploadVideoResponse
	GetStatusCode() *int32
	SetBody(v *RefreshUploadVideoResponseBody) *RefreshUploadVideoResponse
	GetBody() *RefreshUploadVideoResponseBody
}

type RefreshUploadVideoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshUploadVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshUploadVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshUploadVideoResponse) GoString() string {
	return s.String()
}

func (s *RefreshUploadVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshUploadVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshUploadVideoResponse) GetBody() *RefreshUploadVideoResponseBody {
	return s.Body
}

func (s *RefreshUploadVideoResponse) SetHeaders(v map[string]*string) *RefreshUploadVideoResponse {
	s.Headers = v
	return s
}

func (s *RefreshUploadVideoResponse) SetStatusCode(v int32) *RefreshUploadVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshUploadVideoResponse) SetBody(v *RefreshUploadVideoResponseBody) *RefreshUploadVideoResponse {
	s.Body = v
	return s
}

func (s *RefreshUploadVideoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
