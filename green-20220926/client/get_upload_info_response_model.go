// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadInfoResponseBody) *GetUploadInfoResponse
	GetBody() *GetUploadInfoResponseBody
}

type GetUploadInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUploadInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadInfoResponse) GetBody() *GetUploadInfoResponseBody {
	return s.Body
}

func (s *GetUploadInfoResponse) SetHeaders(v map[string]*string) *GetUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUploadInfoResponse) SetStatusCode(v int32) *GetUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadInfoResponse) SetBody(v *GetUploadInfoResponseBody) *GetUploadInfoResponse {
	s.Body = v
	return s
}

func (s *GetUploadInfoResponse) Validate() error {
	return dara.Validate(s)
}
