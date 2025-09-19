// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublishBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublishBatchResponse
	GetStatusCode() *int32
	SetBody(v *ListPublishBatchResponseBody) *ListPublishBatchResponse
	GetBody() *ListPublishBatchResponseBody
}

type ListPublishBatchResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublishBatchResponse) GoString() string {
	return s.String()
}

func (s *ListPublishBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublishBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublishBatchResponse) GetBody() *ListPublishBatchResponseBody {
	return s.Body
}

func (s *ListPublishBatchResponse) SetHeaders(v map[string]*string) *ListPublishBatchResponse {
	s.Headers = v
	return s
}

func (s *ListPublishBatchResponse) SetStatusCode(v int32) *ListPublishBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishBatchResponse) SetBody(v *ListPublishBatchResponseBody) *ListPublishBatchResponse {
	s.Body = v
	return s
}

func (s *ListPublishBatchResponse) Validate() error {
	return dara.Validate(s)
}
