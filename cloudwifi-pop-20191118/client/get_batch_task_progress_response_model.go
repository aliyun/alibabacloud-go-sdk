// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTaskProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTaskProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTaskProgressResponseBody) *GetBatchTaskProgressResponse
	GetBody() *GetBatchTaskProgressResponseBody
}

type GetBatchTaskProgressResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTaskProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTaskProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTaskProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTaskProgressResponse) GetBody() *GetBatchTaskProgressResponseBody {
	return s.Body
}

func (s *GetBatchTaskProgressResponse) SetHeaders(v map[string]*string) *GetBatchTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTaskProgressResponse) SetStatusCode(v int32) *GetBatchTaskProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTaskProgressResponse) SetBody(v *GetBatchTaskProgressResponseBody) *GetBatchTaskProgressResponse {
	s.Body = v
	return s
}

func (s *GetBatchTaskProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
