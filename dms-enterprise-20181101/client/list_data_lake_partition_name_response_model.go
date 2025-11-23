// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakePartitionNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakePartitionNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakePartitionNameResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakePartitionNameResponseBody) *ListDataLakePartitionNameResponse
	GetBody() *ListDataLakePartitionNameResponseBody
}

type ListDataLakePartitionNameResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakePartitionNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakePartitionNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakePartitionNameResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakePartitionNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakePartitionNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakePartitionNameResponse) GetBody() *ListDataLakePartitionNameResponseBody {
	return s.Body
}

func (s *ListDataLakePartitionNameResponse) SetHeaders(v map[string]*string) *ListDataLakePartitionNameResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakePartitionNameResponse) SetStatusCode(v int32) *ListDataLakePartitionNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakePartitionNameResponse) SetBody(v *ListDataLakePartitionNameResponseBody) *ListDataLakePartitionNameResponse {
	s.Body = v
	return s
}

func (s *ListDataLakePartitionNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
