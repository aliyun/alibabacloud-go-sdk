// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipartFileUploadInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultipartFileUploadInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultipartFileUploadInfosResponse
	GetStatusCode() *int32
	SetBody(v *GetMultipartFileUploadInfosResponseBody) *GetMultipartFileUploadInfosResponse
	GetBody() *GetMultipartFileUploadInfosResponseBody
}

type GetMultipartFileUploadInfosResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultipartFileUploadInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultipartFileUploadInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosResponse) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultipartFileUploadInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultipartFileUploadInfosResponse) GetBody() *GetMultipartFileUploadInfosResponseBody {
	return s.Body
}

func (s *GetMultipartFileUploadInfosResponse) SetHeaders(v map[string]*string) *GetMultipartFileUploadInfosResponse {
	s.Headers = v
	return s
}

func (s *GetMultipartFileUploadInfosResponse) SetStatusCode(v int32) *GetMultipartFileUploadInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultipartFileUploadInfosResponse) SetBody(v *GetMultipartFileUploadInfosResponseBody) *GetMultipartFileUploadInfosResponse {
	s.Body = v
	return s
}

func (s *GetMultipartFileUploadInfosResponse) Validate() error {
	return dara.Validate(s)
}
