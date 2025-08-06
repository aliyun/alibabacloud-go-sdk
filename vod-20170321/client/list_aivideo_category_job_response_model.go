// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCategoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoCategoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoCategoryJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoCategoryJobResponseBody) *ListAIVideoCategoryJobResponse
	GetBody() *ListAIVideoCategoryJobResponseBody
}

type ListAIVideoCategoryJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoCategoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoCategoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCategoryJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoCategoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoCategoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoCategoryJobResponse) GetBody() *ListAIVideoCategoryJobResponseBody {
	return s.Body
}

func (s *ListAIVideoCategoryJobResponse) SetHeaders(v map[string]*string) *ListAIVideoCategoryJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoCategoryJobResponse) SetStatusCode(v int32) *ListAIVideoCategoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoCategoryJobResponse) SetBody(v *ListAIVideoCategoryJobResponseBody) *ListAIVideoCategoryJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoCategoryJobResponse) Validate() error {
	return dara.Validate(s)
}
