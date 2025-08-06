// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCensorJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoCensorJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoCensorJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoCensorJobResponseBody) *SubmitAIVideoCensorJobResponse
	GetBody() *SubmitAIVideoCensorJobResponseBody
}

type SubmitAIVideoCensorJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoCensorJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoCensorJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCensorJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCensorJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoCensorJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoCensorJobResponse) GetBody() *SubmitAIVideoCensorJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoCensorJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoCensorJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoCensorJobResponse) SetStatusCode(v int32) *SubmitAIVideoCensorJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoCensorJobResponse) SetBody(v *SubmitAIVideoCensorJobResponseBody) *SubmitAIVideoCensorJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoCensorJobResponse) Validate() error {
	return dara.Validate(s)
}
