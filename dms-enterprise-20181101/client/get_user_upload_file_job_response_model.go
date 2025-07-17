// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserUploadFileJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserUploadFileJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserUploadFileJobResponse
	GetStatusCode() *int32
	SetBody(v *GetUserUploadFileJobResponseBody) *GetUserUploadFileJobResponse
	GetBody() *GetUserUploadFileJobResponseBody
}

type GetUserUploadFileJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserUploadFileJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserUploadFileJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserUploadFileJobResponse) GoString() string {
	return s.String()
}

func (s *GetUserUploadFileJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserUploadFileJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserUploadFileJobResponse) GetBody() *GetUserUploadFileJobResponseBody {
	return s.Body
}

func (s *GetUserUploadFileJobResponse) SetHeaders(v map[string]*string) *GetUserUploadFileJobResponse {
	s.Headers = v
	return s
}

func (s *GetUserUploadFileJobResponse) SetStatusCode(v int32) *GetUserUploadFileJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserUploadFileJobResponse) SetBody(v *GetUserUploadFileJobResponseBody) *GetUserUploadFileJobResponse {
	s.Body = v
	return s
}

func (s *GetUserUploadFileJobResponse) Validate() error {
	return dara.Validate(s)
}
