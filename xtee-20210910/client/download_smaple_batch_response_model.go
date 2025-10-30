// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadSmapleBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadSmapleBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadSmapleBatchResponse
	GetStatusCode() *int32
	SetBody(v *DownloadSmapleBatchResponseBody) *DownloadSmapleBatchResponse
	GetBody() *DownloadSmapleBatchResponseBody
}

type DownloadSmapleBatchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadSmapleBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadSmapleBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadSmapleBatchResponse) GoString() string {
	return s.String()
}

func (s *DownloadSmapleBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadSmapleBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadSmapleBatchResponse) GetBody() *DownloadSmapleBatchResponseBody {
	return s.Body
}

func (s *DownloadSmapleBatchResponse) SetHeaders(v map[string]*string) *DownloadSmapleBatchResponse {
	s.Headers = v
	return s
}

func (s *DownloadSmapleBatchResponse) SetStatusCode(v int32) *DownloadSmapleBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadSmapleBatchResponse) SetBody(v *DownloadSmapleBatchResponseBody) *DownloadSmapleBatchResponse {
	s.Body = v
	return s
}

func (s *DownloadSmapleBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
