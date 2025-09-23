// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSampleApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadSampleApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadSampleApiResponse
	GetStatusCode() *int32
	SetBody(v *UploadSampleApiResponseBody) *UploadSampleApiResponse
	GetBody() *UploadSampleApiResponseBody
}

type UploadSampleApiResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadSampleApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadSampleApiResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleApiResponse) GoString() string {
	return s.String()
}

func (s *UploadSampleApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadSampleApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadSampleApiResponse) GetBody() *UploadSampleApiResponseBody {
	return s.Body
}

func (s *UploadSampleApiResponse) SetHeaders(v map[string]*string) *UploadSampleApiResponse {
	s.Headers = v
	return s
}

func (s *UploadSampleApiResponse) SetStatusCode(v int32) *UploadSampleApiResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadSampleApiResponse) SetBody(v *UploadSampleApiResponseBody) *UploadSampleApiResponse {
	s.Body = v
	return s
}

func (s *UploadSampleApiResponse) Validate() error {
	return dara.Validate(s)
}
