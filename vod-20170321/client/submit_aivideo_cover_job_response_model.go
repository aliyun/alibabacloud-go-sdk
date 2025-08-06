// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCoverJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoCoverJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoCoverJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoCoverJobResponseBody) *SubmitAIVideoCoverJobResponse
	GetBody() *SubmitAIVideoCoverJobResponseBody
}

type SubmitAIVideoCoverJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoCoverJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoCoverJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCoverJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCoverJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoCoverJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoCoverJobResponse) GetBody() *SubmitAIVideoCoverJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoCoverJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoCoverJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoCoverJobResponse) SetStatusCode(v int32) *SubmitAIVideoCoverJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoCoverJobResponse) SetBody(v *SubmitAIVideoCoverJobResponseBody) *SubmitAIVideoCoverJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoCoverJobResponse) Validate() error {
	return dara.Validate(s)
}
