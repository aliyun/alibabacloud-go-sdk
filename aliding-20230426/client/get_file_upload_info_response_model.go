// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileUploadInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileUploadInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetFileUploadInfoResponseBody) *GetFileUploadInfoResponse
	GetBody() *GetFileUploadInfoResponseBody
}

type GetFileUploadInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileUploadInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileUploadInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileUploadInfoResponse) GetBody() *GetFileUploadInfoResponseBody {
	return s.Body
}

func (s *GetFileUploadInfoResponse) SetHeaders(v map[string]*string) *GetFileUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileUploadInfoResponse) SetStatusCode(v int32) *GetFileUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileUploadInfoResponse) SetBody(v *GetFileUploadInfoResponseBody) *GetFileUploadInfoResponse {
	s.Body = v
	return s
}

func (s *GetFileUploadInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
