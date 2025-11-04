// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshUploadMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshUploadMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshUploadMediaResponse
	GetStatusCode() *int32
	SetBody(v *RefreshUploadMediaResponseBody) *RefreshUploadMediaResponse
	GetBody() *RefreshUploadMediaResponseBody
}

type RefreshUploadMediaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshUploadMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshUploadMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshUploadMediaResponse) GoString() string {
	return s.String()
}

func (s *RefreshUploadMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshUploadMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshUploadMediaResponse) GetBody() *RefreshUploadMediaResponseBody {
	return s.Body
}

func (s *RefreshUploadMediaResponse) SetHeaders(v map[string]*string) *RefreshUploadMediaResponse {
	s.Headers = v
	return s
}

func (s *RefreshUploadMediaResponse) SetStatusCode(v int32) *RefreshUploadMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshUploadMediaResponse) SetBody(v *RefreshUploadMediaResponseBody) *RefreshUploadMediaResponse {
	s.Body = v
	return s
}

func (s *RefreshUploadMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
