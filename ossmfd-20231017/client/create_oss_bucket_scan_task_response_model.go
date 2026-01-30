// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssBucketScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOssBucketScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOssBucketScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateOssBucketScanTaskResponseBody) *CreateOssBucketScanTaskResponse
	GetBody() *CreateOssBucketScanTaskResponseBody
}

type CreateOssBucketScanTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOssBucketScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOssBucketScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOssBucketScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateOssBucketScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOssBucketScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOssBucketScanTaskResponse) GetBody() *CreateOssBucketScanTaskResponseBody {
	return s.Body
}

func (s *CreateOssBucketScanTaskResponse) SetHeaders(v map[string]*string) *CreateOssBucketScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateOssBucketScanTaskResponse) SetStatusCode(v int32) *CreateOssBucketScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOssBucketScanTaskResponse) SetBody(v *CreateOssBucketScanTaskResponseBody) *CreateOssBucketScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateOssBucketScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
