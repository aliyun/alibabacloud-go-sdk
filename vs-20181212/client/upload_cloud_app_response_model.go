// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCloudAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadCloudAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadCloudAppResponse
	GetStatusCode() *int32
	SetBody(v *UploadCloudAppResponseBody) *UploadCloudAppResponse
	GetBody() *UploadCloudAppResponseBody
}

type UploadCloudAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadCloudAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadCloudAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadCloudAppResponse) GoString() string {
	return s.String()
}

func (s *UploadCloudAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadCloudAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadCloudAppResponse) GetBody() *UploadCloudAppResponseBody {
	return s.Body
}

func (s *UploadCloudAppResponse) SetHeaders(v map[string]*string) *UploadCloudAppResponse {
	s.Headers = v
	return s
}

func (s *UploadCloudAppResponse) SetStatusCode(v int32) *UploadCloudAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCloudAppResponse) SetBody(v *UploadCloudAppResponseBody) *UploadCloudAppResponse {
	s.Body = v
	return s
}

func (s *UploadCloudAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
