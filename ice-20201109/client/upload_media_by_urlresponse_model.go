// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaByURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadMediaByURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadMediaByURLResponse
	GetStatusCode() *int32
	SetBody(v *UploadMediaByURLResponseBody) *UploadMediaByURLResponse
	GetBody() *UploadMediaByURLResponseBody
}

type UploadMediaByURLResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadMediaByURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadMediaByURLResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaByURLResponse) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadMediaByURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadMediaByURLResponse) GetBody() *UploadMediaByURLResponseBody {
	return s.Body
}

func (s *UploadMediaByURLResponse) SetHeaders(v map[string]*string) *UploadMediaByURLResponse {
	s.Headers = v
	return s
}

func (s *UploadMediaByURLResponse) SetStatusCode(v int32) *UploadMediaByURLResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMediaByURLResponse) SetBody(v *UploadMediaByURLResponseBody) *UploadMediaByURLResponse {
	s.Body = v
	return s
}

func (s *UploadMediaByURLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
