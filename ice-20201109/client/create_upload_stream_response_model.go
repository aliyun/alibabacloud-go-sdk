// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadStreamResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadStreamResponseBody) *CreateUploadStreamResponse
	GetBody() *CreateUploadStreamResponseBody
}

type CreateUploadStreamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadStreamResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadStreamResponse) GetBody() *CreateUploadStreamResponseBody {
	return s.Body
}

func (s *CreateUploadStreamResponse) SetHeaders(v map[string]*string) *CreateUploadStreamResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadStreamResponse) SetStatusCode(v int32) *CreateUploadStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadStreamResponse) SetBody(v *CreateUploadStreamResponseBody) *CreateUploadStreamResponse {
	s.Body = v
	return s
}

func (s *CreateUploadStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
