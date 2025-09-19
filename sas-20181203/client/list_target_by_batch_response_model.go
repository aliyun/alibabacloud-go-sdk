// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetByBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTargetByBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTargetByBatchResponse
	GetStatusCode() *int32
	SetBody(v *ListTargetByBatchResponseBody) *ListTargetByBatchResponse
	GetBody() *ListTargetByBatchResponseBody
}

type ListTargetByBatchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetByBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetByBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTargetByBatchResponse) GoString() string {
	return s.String()
}

func (s *ListTargetByBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTargetByBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTargetByBatchResponse) GetBody() *ListTargetByBatchResponseBody {
	return s.Body
}

func (s *ListTargetByBatchResponse) SetHeaders(v map[string]*string) *ListTargetByBatchResponse {
	s.Headers = v
	return s
}

func (s *ListTargetByBatchResponse) SetStatusCode(v int32) *ListTargetByBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetByBatchResponse) SetBody(v *ListTargetByBatchResponseBody) *ListTargetByBatchResponse {
	s.Body = v
	return s
}

func (s *ListTargetByBatchResponse) Validate() error {
	return dara.Validate(s)
}
