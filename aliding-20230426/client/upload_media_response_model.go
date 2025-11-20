// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadMediaResponse
	GetStatusCode() *int32
	SetBody(v *UploadMediaResponseBody) *UploadMediaResponse
	GetBody() *UploadMediaResponseBody
}

type UploadMediaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaResponse) GoString() string {
	return s.String()
}

func (s *UploadMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadMediaResponse) GetBody() *UploadMediaResponseBody {
	return s.Body
}

func (s *UploadMediaResponse) SetHeaders(v map[string]*string) *UploadMediaResponse {
	s.Headers = v
	return s
}

func (s *UploadMediaResponse) SetStatusCode(v int32) *UploadMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMediaResponse) SetBody(v *UploadMediaResponseBody) *UploadMediaResponse {
	s.Body = v
	return s
}

func (s *UploadMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
