// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDocResponse
	GetStatusCode() *int32
	SetBody(v *UploadDocResponseBody) *UploadDocResponse
	GetBody() *UploadDocResponseBody
}

type UploadDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDocResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadDocResponse) GoString() string {
	return s.String()
}

func (s *UploadDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDocResponse) GetBody() *UploadDocResponseBody {
	return s.Body
}

func (s *UploadDocResponse) SetHeaders(v map[string]*string) *UploadDocResponse {
	s.Headers = v
	return s
}

func (s *UploadDocResponse) SetStatusCode(v int32) *UploadDocResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDocResponse) SetBody(v *UploadDocResponseBody) *UploadDocResponse {
	s.Body = v
	return s
}

func (s *UploadDocResponse) Validate() error {
	return dara.Validate(s)
}
