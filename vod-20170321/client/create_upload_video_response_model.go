// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadVideoResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadVideoResponseBody) *CreateUploadVideoResponse
	GetBody() *CreateUploadVideoResponseBody
}

type CreateUploadVideoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadVideoResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadVideoResponse) GetBody() *CreateUploadVideoResponseBody {
	return s.Body
}

func (s *CreateUploadVideoResponse) SetHeaders(v map[string]*string) *CreateUploadVideoResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadVideoResponse) SetStatusCode(v int32) *CreateUploadVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadVideoResponse) SetBody(v *CreateUploadVideoResponseBody) *CreateUploadVideoResponse {
	s.Body = v
	return s
}

func (s *CreateUploadVideoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
