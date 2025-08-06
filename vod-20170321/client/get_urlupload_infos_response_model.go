// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetURLUploadInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetURLUploadInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetURLUploadInfosResponse
	GetStatusCode() *int32
	SetBody(v *GetURLUploadInfosResponseBody) *GetURLUploadInfosResponse
	GetBody() *GetURLUploadInfosResponseBody
}

type GetURLUploadInfosResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetURLUploadInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetURLUploadInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s GetURLUploadInfosResponse) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetURLUploadInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetURLUploadInfosResponse) GetBody() *GetURLUploadInfosResponseBody {
	return s.Body
}

func (s *GetURLUploadInfosResponse) SetHeaders(v map[string]*string) *GetURLUploadInfosResponse {
	s.Headers = v
	return s
}

func (s *GetURLUploadInfosResponse) SetStatusCode(v int32) *GetURLUploadInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetURLUploadInfosResponse) SetBody(v *GetURLUploadInfosResponseBody) *GetURLUploadInfosResponse {
	s.Body = v
	return s
}

func (s *GetURLUploadInfosResponse) Validate() error {
	return dara.Validate(s)
}
