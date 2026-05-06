// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSampleFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadSampleFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadSampleFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadSampleFileResponseBody) *UploadSampleFileResponse
	GetBody() *UploadSampleFileResponseBody
}

type UploadSampleFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadSampleFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadSampleFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleFileResponse) GoString() string {
	return s.String()
}

func (s *UploadSampleFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadSampleFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadSampleFileResponse) GetBody() *UploadSampleFileResponseBody {
	return s.Body
}

func (s *UploadSampleFileResponse) SetHeaders(v map[string]*string) *UploadSampleFileResponse {
	s.Headers = v
	return s
}

func (s *UploadSampleFileResponse) SetStatusCode(v int32) *UploadSampleFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadSampleFileResponse) SetBody(v *UploadSampleFileResponseBody) *UploadSampleFileResponse {
	s.Body = v
	return s
}

func (s *UploadSampleFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
