// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadStreamByURLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadStreamByURLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadStreamByURLResponse
	GetStatusCode() *int32
	SetBody(v *UploadStreamByURLResponseBody) *UploadStreamByURLResponse
	GetBody() *UploadStreamByURLResponseBody
}

type UploadStreamByURLResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadStreamByURLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadStreamByURLResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadStreamByURLResponse) GoString() string {
	return s.String()
}

func (s *UploadStreamByURLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadStreamByURLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadStreamByURLResponse) GetBody() *UploadStreamByURLResponseBody {
	return s.Body
}

func (s *UploadStreamByURLResponse) SetHeaders(v map[string]*string) *UploadStreamByURLResponse {
	s.Headers = v
	return s
}

func (s *UploadStreamByURLResponse) SetStatusCode(v int32) *UploadStreamByURLResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadStreamByURLResponse) SetBody(v *UploadStreamByURLResponseBody) *UploadStreamByURLResponse {
	s.Body = v
	return s
}

func (s *UploadStreamByURLResponse) Validate() error {
	return dara.Validate(s)
}
