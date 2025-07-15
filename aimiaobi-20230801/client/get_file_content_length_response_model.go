// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileContentLengthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileContentLengthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileContentLengthResponse
	GetStatusCode() *int32
	SetBody(v *GetFileContentLengthResponseBody) *GetFileContentLengthResponse
	GetBody() *GetFileContentLengthResponseBody
}

type GetFileContentLengthResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileContentLengthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileContentLengthResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileContentLengthResponse) GoString() string {
	return s.String()
}

func (s *GetFileContentLengthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileContentLengthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileContentLengthResponse) GetBody() *GetFileContentLengthResponseBody {
	return s.Body
}

func (s *GetFileContentLengthResponse) SetHeaders(v map[string]*string) *GetFileContentLengthResponse {
	s.Headers = v
	return s
}

func (s *GetFileContentLengthResponse) SetStatusCode(v int32) *GetFileContentLengthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileContentLengthResponse) SetBody(v *GetFileContentLengthResponseBody) *GetFileContentLengthResponse {
	s.Body = v
	return s
}

func (s *GetFileContentLengthResponse) Validate() error {
	return dara.Validate(s)
}
