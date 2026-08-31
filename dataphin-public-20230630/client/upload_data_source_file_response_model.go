// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSourceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDataSourceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDataSourceFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadDataSourceFileResponseBody) *UploadDataSourceFileResponse
	GetBody() *UploadDataSourceFileResponseBody
}

type UploadDataSourceFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataSourceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataSourceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSourceFileResponse) GoString() string {
	return s.String()
}

func (s *UploadDataSourceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDataSourceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDataSourceFileResponse) GetBody() *UploadDataSourceFileResponseBody {
	return s.Body
}

func (s *UploadDataSourceFileResponse) SetHeaders(v map[string]*string) *UploadDataSourceFileResponse {
	s.Headers = v
	return s
}

func (s *UploadDataSourceFileResponse) SetStatusCode(v int32) *UploadDataSourceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataSourceFileResponse) SetBody(v *UploadDataSourceFileResponseBody) *UploadDataSourceFileResponse {
	s.Body = v
	return s
}

func (s *UploadDataSourceFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
