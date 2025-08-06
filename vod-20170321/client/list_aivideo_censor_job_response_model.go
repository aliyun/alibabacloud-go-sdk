// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCensorJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoCensorJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoCensorJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoCensorJobResponseBody) *ListAIVideoCensorJobResponse
	GetBody() *ListAIVideoCensorJobResponseBody
}

type ListAIVideoCensorJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoCensorJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoCensorJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCensorJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoCensorJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoCensorJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoCensorJobResponse) GetBody() *ListAIVideoCensorJobResponseBody {
	return s.Body
}

func (s *ListAIVideoCensorJobResponse) SetHeaders(v map[string]*string) *ListAIVideoCensorJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoCensorJobResponse) SetStatusCode(v int32) *ListAIVideoCensorJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoCensorJobResponse) SetBody(v *ListAIVideoCensorJobResponseBody) *ListAIVideoCensorJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoCensorJobResponse) Validate() error {
	return dara.Validate(s)
}
