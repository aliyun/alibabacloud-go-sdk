// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadBiddingDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadBiddingDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadBiddingDocResponse
	GetStatusCode() *int32
	SetBody(v *DownloadBiddingDocResponseBody) *DownloadBiddingDocResponse
	GetBody() *DownloadBiddingDocResponseBody
}

type DownloadBiddingDocResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadBiddingDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadBiddingDocResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadBiddingDocResponse) GoString() string {
	return s.String()
}

func (s *DownloadBiddingDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadBiddingDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadBiddingDocResponse) GetBody() *DownloadBiddingDocResponseBody {
	return s.Body
}

func (s *DownloadBiddingDocResponse) SetHeaders(v map[string]*string) *DownloadBiddingDocResponse {
	s.Headers = v
	return s
}

func (s *DownloadBiddingDocResponse) SetStatusCode(v int32) *DownloadBiddingDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadBiddingDocResponse) SetBody(v *DownloadBiddingDocResponseBody) *DownloadBiddingDocResponse {
	s.Body = v
	return s
}

func (s *DownloadBiddingDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
