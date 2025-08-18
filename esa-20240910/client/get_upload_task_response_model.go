// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadTaskResponseBody) *GetUploadTaskResponse
	GetBody() *GetUploadTaskResponseBody
}

type GetUploadTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadTaskResponse) GoString() string {
	return s.String()
}

func (s *GetUploadTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadTaskResponse) GetBody() *GetUploadTaskResponseBody {
	return s.Body
}

func (s *GetUploadTaskResponse) SetHeaders(v map[string]*string) *GetUploadTaskResponse {
	s.Headers = v
	return s
}

func (s *GetUploadTaskResponse) SetStatusCode(v int32) *GetUploadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadTaskResponse) SetBody(v *GetUploadTaskResponseBody) *GetUploadTaskResponse {
	s.Body = v
	return s
}

func (s *GetUploadTaskResponse) Validate() error {
	return dara.Validate(s)
}
