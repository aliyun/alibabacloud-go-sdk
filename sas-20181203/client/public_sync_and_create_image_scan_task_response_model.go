// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicSyncAndCreateImageScanTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublicSyncAndCreateImageScanTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublicSyncAndCreateImageScanTaskResponse
	GetStatusCode() *int32
	SetBody(v *PublicSyncAndCreateImageScanTaskResponseBody) *PublicSyncAndCreateImageScanTaskResponse
	GetBody() *PublicSyncAndCreateImageScanTaskResponseBody
}

type PublicSyncAndCreateImageScanTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublicSyncAndCreateImageScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublicSyncAndCreateImageScanTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PublicSyncAndCreateImageScanTaskResponse) GoString() string {
	return s.String()
}

func (s *PublicSyncAndCreateImageScanTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublicSyncAndCreateImageScanTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublicSyncAndCreateImageScanTaskResponse) GetBody() *PublicSyncAndCreateImageScanTaskResponseBody {
	return s.Body
}

func (s *PublicSyncAndCreateImageScanTaskResponse) SetHeaders(v map[string]*string) *PublicSyncAndCreateImageScanTaskResponse {
	s.Headers = v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponse) SetStatusCode(v int32) *PublicSyncAndCreateImageScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponse) SetBody(v *PublicSyncAndCreateImageScanTaskResponseBody) *PublicSyncAndCreateImageScanTaskResponse {
	s.Body = v
	return s
}

func (s *PublicSyncAndCreateImageScanTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
