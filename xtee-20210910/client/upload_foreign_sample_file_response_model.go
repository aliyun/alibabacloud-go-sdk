// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadForeignSampleFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadForeignSampleFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadForeignSampleFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadForeignSampleFileResponseBody) *UploadForeignSampleFileResponse
	GetBody() *UploadForeignSampleFileResponseBody
}

type UploadForeignSampleFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadForeignSampleFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadForeignSampleFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadForeignSampleFileResponse) GoString() string {
	return s.String()
}

func (s *UploadForeignSampleFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadForeignSampleFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadForeignSampleFileResponse) GetBody() *UploadForeignSampleFileResponseBody {
	return s.Body
}

func (s *UploadForeignSampleFileResponse) SetHeaders(v map[string]*string) *UploadForeignSampleFileResponse {
	s.Headers = v
	return s
}

func (s *UploadForeignSampleFileResponse) SetStatusCode(v int32) *UploadForeignSampleFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadForeignSampleFileResponse) SetBody(v *UploadForeignSampleFileResponseBody) *UploadForeignSampleFileResponse {
	s.Body = v
	return s
}

func (s *UploadForeignSampleFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
