// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyDownloadTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVerifyDownloadTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVerifyDownloadTaskResponse
	GetStatusCode() *int32
	SetBody(v *QueryVerifyDownloadTaskResponseBody) *QueryVerifyDownloadTaskResponse
	GetBody() *QueryVerifyDownloadTaskResponseBody
}

type QueryVerifyDownloadTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVerifyDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVerifyDownloadTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryVerifyDownloadTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVerifyDownloadTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVerifyDownloadTaskResponse) GetBody() *QueryVerifyDownloadTaskResponseBody {
	return s.Body
}

func (s *QueryVerifyDownloadTaskResponse) SetHeaders(v map[string]*string) *QueryVerifyDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryVerifyDownloadTaskResponse) SetStatusCode(v int32) *QueryVerifyDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVerifyDownloadTaskResponse) SetBody(v *QueryVerifyDownloadTaskResponseBody) *QueryVerifyDownloadTaskResponse {
	s.Body = v
	return s
}

func (s *QueryVerifyDownloadTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
