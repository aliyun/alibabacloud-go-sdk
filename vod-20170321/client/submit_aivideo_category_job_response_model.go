// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIVideoCategoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAIVideoCategoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAIVideoCategoryJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAIVideoCategoryJobResponseBody) *SubmitAIVideoCategoryJobResponse
	GetBody() *SubmitAIVideoCategoryJobResponseBody
}

type SubmitAIVideoCategoryJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAIVideoCategoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAIVideoCategoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIVideoCategoryJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIVideoCategoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAIVideoCategoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAIVideoCategoryJobResponse) GetBody() *SubmitAIVideoCategoryJobResponseBody {
	return s.Body
}

func (s *SubmitAIVideoCategoryJobResponse) SetHeaders(v map[string]*string) *SubmitAIVideoCategoryJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIVideoCategoryJobResponse) SetStatusCode(v int32) *SubmitAIVideoCategoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAIVideoCategoryJobResponse) SetBody(v *SubmitAIVideoCategoryJobResponseBody) *SubmitAIVideoCategoryJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAIVideoCategoryJobResponse) Validate() error {
	return dara.Validate(s)
}
