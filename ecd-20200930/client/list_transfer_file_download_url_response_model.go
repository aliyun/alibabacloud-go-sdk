// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransferFileDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransferFileDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransferFileDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *ListTransferFileDownloadUrlResponseBody) *ListTransferFileDownloadUrlResponse
	GetBody() *ListTransferFileDownloadUrlResponseBody
}

type ListTransferFileDownloadUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransferFileDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransferFileDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransferFileDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *ListTransferFileDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransferFileDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransferFileDownloadUrlResponse) GetBody() *ListTransferFileDownloadUrlResponseBody {
	return s.Body
}

func (s *ListTransferFileDownloadUrlResponse) SetHeaders(v map[string]*string) *ListTransferFileDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *ListTransferFileDownloadUrlResponse) SetStatusCode(v int32) *ListTransferFileDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransferFileDownloadUrlResponse) SetBody(v *ListTransferFileDownloadUrlResponseBody) *ListTransferFileDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *ListTransferFileDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
