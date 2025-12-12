// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDataResponse
	GetStatusCode() *int32
	SetBody(v *UploadDataResponseBody) *UploadDataResponse
	GetBody() *UploadDataResponseBody
}

type UploadDataResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDataResponse) GoString() string {
	return s.String()
}

func (s *UploadDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDataResponse) GetBody() *UploadDataResponseBody {
	return s.Body
}

func (s *UploadDataResponse) SetHeaders(v map[string]*string) *UploadDataResponse {
	s.Headers = v
	return s
}

func (s *UploadDataResponse) SetStatusCode(v int32) *UploadDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDataResponse) SetBody(v *UploadDataResponseBody) *UploadDataResponse {
	s.Body = v
	return s
}

func (s *UploadDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
