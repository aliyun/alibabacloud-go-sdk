// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitUploadTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitUploadTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitUploadTaskResponse
	GetStatusCode() *int32
	SetBody(v *SubmitUploadTaskResponseBody) *SubmitUploadTaskResponse
	GetBody() *SubmitUploadTaskResponseBody
}

type SubmitUploadTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitUploadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitUploadTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitUploadTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitUploadTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitUploadTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitUploadTaskResponse) GetBody() *SubmitUploadTaskResponseBody {
	return s.Body
}

func (s *SubmitUploadTaskResponse) SetHeaders(v map[string]*string) *SubmitUploadTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitUploadTaskResponse) SetStatusCode(v int32) *SubmitUploadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitUploadTaskResponse) SetBody(v *SubmitUploadTaskResponseBody) *SubmitUploadTaskResponse {
	s.Body = v
	return s
}

func (s *SubmitUploadTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
