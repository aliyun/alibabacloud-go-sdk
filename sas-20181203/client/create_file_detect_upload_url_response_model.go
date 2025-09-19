// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDetectUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileDetectUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileDetectUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileDetectUploadUrlResponseBody) *CreateFileDetectUploadUrlResponse
	GetBody() *CreateFileDetectUploadUrlResponseBody
}

type CreateFileDetectUploadUrlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileDetectUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileDetectUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateFileDetectUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileDetectUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileDetectUploadUrlResponse) GetBody() *CreateFileDetectUploadUrlResponseBody {
	return s.Body
}

func (s *CreateFileDetectUploadUrlResponse) SetHeaders(v map[string]*string) *CreateFileDetectUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateFileDetectUploadUrlResponse) SetStatusCode(v int32) *CreateFileDetectUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponse) SetBody(v *CreateFileDetectUploadUrlResponseBody) *CreateFileDetectUploadUrlResponse {
	s.Body = v
	return s
}

func (s *CreateFileDetectUploadUrlResponse) Validate() error {
	return dara.Validate(s)
}
