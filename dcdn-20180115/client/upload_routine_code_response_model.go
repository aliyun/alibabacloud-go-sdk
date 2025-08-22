// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRoutineCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadRoutineCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadRoutineCodeResponse
	GetStatusCode() *int32
	SetBody(v *UploadRoutineCodeResponseBody) *UploadRoutineCodeResponse
	GetBody() *UploadRoutineCodeResponseBody
}

type UploadRoutineCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRoutineCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadRoutineCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadRoutineCodeResponse) GoString() string {
	return s.String()
}

func (s *UploadRoutineCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadRoutineCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadRoutineCodeResponse) GetBody() *UploadRoutineCodeResponseBody {
	return s.Body
}

func (s *UploadRoutineCodeResponse) SetHeaders(v map[string]*string) *UploadRoutineCodeResponse {
	s.Headers = v
	return s
}

func (s *UploadRoutineCodeResponse) SetStatusCode(v int32) *UploadRoutineCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRoutineCodeResponse) SetBody(v *UploadRoutineCodeResponseBody) *UploadRoutineCodeResponse {
	s.Body = v
	return s
}

func (s *UploadRoutineCodeResponse) Validate() error {
	return dara.Validate(s)
}
