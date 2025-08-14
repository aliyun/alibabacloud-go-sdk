// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrlUploadInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUrlUploadInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUrlUploadInfosResponse
	GetStatusCode() *int32
	SetBody(v *GetUrlUploadInfosResponseBody) *GetUrlUploadInfosResponse
	GetBody() *GetUrlUploadInfosResponseBody
}

type GetUrlUploadInfosResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUrlUploadInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUrlUploadInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUrlUploadInfosResponse) GoString() string {
	return s.String()
}

func (s *GetUrlUploadInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUrlUploadInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUrlUploadInfosResponse) GetBody() *GetUrlUploadInfosResponseBody {
	return s.Body
}

func (s *GetUrlUploadInfosResponse) SetHeaders(v map[string]*string) *GetUrlUploadInfosResponse {
	s.Headers = v
	return s
}

func (s *GetUrlUploadInfosResponse) SetStatusCode(v int32) *GetUrlUploadInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUrlUploadInfosResponse) SetBody(v *GetUrlUploadInfosResponseBody) *GetUrlUploadInfosResponse {
	s.Body = v
	return s
}

func (s *GetUrlUploadInfosResponse) Validate() error {
	return dara.Validate(s)
}
