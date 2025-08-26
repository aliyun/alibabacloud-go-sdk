// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakePartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakePartitionResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakePartitionResponseBody) *ListDataLakePartitionResponse
	GetBody() *ListDataLakePartitionResponseBody
}

type ListDataLakePartitionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakePartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakePartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakePartitionResponse) GetBody() *ListDataLakePartitionResponseBody {
	return s.Body
}

func (s *ListDataLakePartitionResponse) SetHeaders(v map[string]*string) *ListDataLakePartitionResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakePartitionResponse) SetStatusCode(v int32) *ListDataLakePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakePartitionResponse) SetBody(v *ListDataLakePartitionResponseBody) *ListDataLakePartitionResponse {
	s.Body = v
	return s
}

func (s *ListDataLakePartitionResponse) Validate() error {
	return dara.Validate(s)
}
