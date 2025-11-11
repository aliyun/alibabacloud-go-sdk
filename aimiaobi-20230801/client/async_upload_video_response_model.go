// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncUploadVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncUploadVideoResponse
	GetStatusCode() *int32
	SetBody(v *AsyncUploadVideoResponseBody) *AsyncUploadVideoResponse
	GetBody() *AsyncUploadVideoResponseBody
}

type AsyncUploadVideoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncUploadVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncUploadVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoResponse) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncUploadVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncUploadVideoResponse) GetBody() *AsyncUploadVideoResponseBody {
	return s.Body
}

func (s *AsyncUploadVideoResponse) SetHeaders(v map[string]*string) *AsyncUploadVideoResponse {
	s.Headers = v
	return s
}

func (s *AsyncUploadVideoResponse) SetStatusCode(v int32) *AsyncUploadVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncUploadVideoResponse) SetBody(v *AsyncUploadVideoResponseBody) *AsyncUploadVideoResponse {
	s.Body = v
	return s
}

func (s *AsyncUploadVideoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
