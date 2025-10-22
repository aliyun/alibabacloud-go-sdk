// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionByFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakePartitionByFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakePartitionByFilterResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakePartitionByFilterResponseBody) *ListDataLakePartitionByFilterResponse
	GetBody() *ListDataLakePartitionByFilterResponseBody
}

type ListDataLakePartitionByFilterResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakePartitionByFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakePartitionByFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionByFilterResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionByFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakePartitionByFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakePartitionByFilterResponse) GetBody() *ListDataLakePartitionByFilterResponseBody {
	return s.Body
}

func (s *ListDataLakePartitionByFilterResponse) SetHeaders(v map[string]*string) *ListDataLakePartitionByFilterResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakePartitionByFilterResponse) SetStatusCode(v int32) *ListDataLakePartitionByFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakePartitionByFilterResponse) SetBody(v *ListDataLakePartitionByFilterResponseBody) *ListDataLakePartitionByFilterResponse {
	s.Body = v
	return s
}

func (s *ListDataLakePartitionByFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
