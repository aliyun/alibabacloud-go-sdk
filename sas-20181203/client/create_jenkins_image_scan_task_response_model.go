// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJenkinsImageScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJenkinsImageScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJenkinsImageScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateJenkinsImageScanTaskResponseBody) *CreateJenkinsImageScanTaskResponse
	GetBody() *CreateJenkinsImageScanTaskResponseBody
}

type CreateJenkinsImageScanTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJenkinsImageScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJenkinsImageScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJenkinsImageScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateJenkinsImageScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJenkinsImageScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJenkinsImageScanTaskResponse) GetBody() *CreateJenkinsImageScanTaskResponseBody {
	return s.Body
}

func (s *CreateJenkinsImageScanTaskResponse) SetHeaders(v map[string]*string) *CreateJenkinsImageScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateJenkinsImageScanTaskResponse) SetStatusCode(v int32) *CreateJenkinsImageScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJenkinsImageScanTaskResponse) SetBody(v *CreateJenkinsImageScanTaskResponseBody) *CreateJenkinsImageScanTaskResponse {
	s.Body = v
	return s
}

func (s *CreateJenkinsImageScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
