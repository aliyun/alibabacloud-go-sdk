// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadFileJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadFileJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadFileJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadFileJobResponseBody) *CreateUploadFileJobResponse
	GetBody() *CreateUploadFileJobResponseBody
}

type CreateUploadFileJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadFileJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadFileJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadFileJobResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadFileJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadFileJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadFileJobResponse) GetBody() *CreateUploadFileJobResponseBody {
	return s.Body
}

func (s *CreateUploadFileJobResponse) SetHeaders(v map[string]*string) *CreateUploadFileJobResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadFileJobResponse) SetStatusCode(v int32) *CreateUploadFileJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadFileJobResponse) SetBody(v *CreateUploadFileJobResponseBody) *CreateUploadFileJobResponse {
	s.Body = v
	return s
}

func (s *CreateUploadFileJobResponse) Validate() error {
	return dara.Validate(s)
}
