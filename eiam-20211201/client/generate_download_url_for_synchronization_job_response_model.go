// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDownloadUrlForSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDownloadUrlForSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDownloadUrlForSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDownloadUrlForSynchronizationJobResponseBody) *GenerateDownloadUrlForSynchronizationJobResponse
	GetBody() *GenerateDownloadUrlForSynchronizationJobResponseBody
}

type GenerateDownloadUrlForSynchronizationJobResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDownloadUrlForSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDownloadUrlForSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDownloadUrlForSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *GenerateDownloadUrlForSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDownloadUrlForSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDownloadUrlForSynchronizationJobResponse) GetBody() *GenerateDownloadUrlForSynchronizationJobResponseBody {
	return s.Body
}

func (s *GenerateDownloadUrlForSynchronizationJobResponse) SetHeaders(v map[string]*string) *GenerateDownloadUrlForSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *GenerateDownloadUrlForSynchronizationJobResponse) SetStatusCode(v int32) *GenerateDownloadUrlForSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDownloadUrlForSynchronizationJobResponse) SetBody(v *GenerateDownloadUrlForSynchronizationJobResponseBody) *GenerateDownloadUrlForSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *GenerateDownloadUrlForSynchronizationJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
