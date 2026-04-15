// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScencegroupFileDownloadurlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScencegroupFileDownloadurlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScencegroupFileDownloadurlResponse
	GetStatusCode() *int32
	SetBody(v *GetScencegroupFileDownloadurlResponseBody) *GetScencegroupFileDownloadurlResponse
	GetBody() *GetScencegroupFileDownloadurlResponseBody
}

type GetScencegroupFileDownloadurlResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScencegroupFileDownloadurlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScencegroupFileDownloadurlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScencegroupFileDownloadurlResponse) GoString() string {
	return s.String()
}

func (s *GetScencegroupFileDownloadurlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScencegroupFileDownloadurlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScencegroupFileDownloadurlResponse) GetBody() *GetScencegroupFileDownloadurlResponseBody {
	return s.Body
}

func (s *GetScencegroupFileDownloadurlResponse) SetHeaders(v map[string]*string) *GetScencegroupFileDownloadurlResponse {
	s.Headers = v
	return s
}

func (s *GetScencegroupFileDownloadurlResponse) SetStatusCode(v int32) *GetScencegroupFileDownloadurlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScencegroupFileDownloadurlResponse) SetBody(v *GetScencegroupFileDownloadurlResponseBody) *GetScencegroupFileDownloadurlResponse {
	s.Body = v
	return s
}

func (s *GetScencegroupFileDownloadurlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
