// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOSSMultimodalDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadOSSMultimodalDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadOSSMultimodalDatasetResponse
	GetStatusCode() *int32
	SetBody(v *UploadOSSMultimodalDatasetResponseBody) *UploadOSSMultimodalDatasetResponse
	GetBody() *UploadOSSMultimodalDatasetResponseBody
}

type UploadOSSMultimodalDatasetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadOSSMultimodalDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadOSSMultimodalDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadOSSMultimodalDatasetResponse) GoString() string {
	return s.String()
}

func (s *UploadOSSMultimodalDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadOSSMultimodalDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadOSSMultimodalDatasetResponse) GetBody() *UploadOSSMultimodalDatasetResponseBody {
	return s.Body
}

func (s *UploadOSSMultimodalDatasetResponse) SetHeaders(v map[string]*string) *UploadOSSMultimodalDatasetResponse {
	s.Headers = v
	return s
}

func (s *UploadOSSMultimodalDatasetResponse) SetStatusCode(v int32) *UploadOSSMultimodalDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadOSSMultimodalDatasetResponse) SetBody(v *UploadOSSMultimodalDatasetResponseBody) *UploadOSSMultimodalDatasetResponse {
	s.Body = v
	return s
}

func (s *UploadOSSMultimodalDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
