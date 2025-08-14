// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCognitionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitVideoCognitionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitVideoCognitionJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitVideoCognitionJobResponseBody) *SubmitVideoCognitionJobResponse
	GetBody() *SubmitVideoCognitionJobResponseBody
}

type SubmitVideoCognitionJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitVideoCognitionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitVideoCognitionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCognitionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitVideoCognitionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitVideoCognitionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitVideoCognitionJobResponse) GetBody() *SubmitVideoCognitionJobResponseBody {
	return s.Body
}

func (s *SubmitVideoCognitionJobResponse) SetHeaders(v map[string]*string) *SubmitVideoCognitionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitVideoCognitionJobResponse) SetStatusCode(v int32) *SubmitVideoCognitionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitVideoCognitionJobResponse) SetBody(v *SubmitVideoCognitionJobResponseBody) *SubmitVideoCognitionJobResponse {
	s.Body = v
	return s
}

func (s *SubmitVideoCognitionJobResponse) Validate() error {
	return dara.Validate(s)
}
