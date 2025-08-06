// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoCoverJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoCoverJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoCoverJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoCoverJobResponseBody) *ListAIVideoCoverJobResponse
	GetBody() *ListAIVideoCoverJobResponseBody
}

type ListAIVideoCoverJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoCoverJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoCoverJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoCoverJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoCoverJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoCoverJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoCoverJobResponse) GetBody() *ListAIVideoCoverJobResponseBody {
	return s.Body
}

func (s *ListAIVideoCoverJobResponse) SetHeaders(v map[string]*string) *ListAIVideoCoverJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoCoverJobResponse) SetStatusCode(v int32) *ListAIVideoCoverJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoCoverJobResponse) SetBody(v *ListAIVideoCoverJobResponseBody) *ListAIVideoCoverJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoCoverJobResponse) Validate() error {
	return dara.Validate(s)
}
