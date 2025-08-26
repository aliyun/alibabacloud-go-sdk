// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakePartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataLakePartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataLakePartitionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataLakePartitionResponseBody) *DeleteDataLakePartitionResponse
	GetBody() *DeleteDataLakePartitionResponseBody
}

type DeleteDataLakePartitionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataLakePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataLakePartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakePartitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataLakePartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataLakePartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataLakePartitionResponse) GetBody() *DeleteDataLakePartitionResponseBody {
	return s.Body
}

func (s *DeleteDataLakePartitionResponse) SetHeaders(v map[string]*string) *DeleteDataLakePartitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataLakePartitionResponse) SetStatusCode(v int32) *DeleteDataLakePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataLakePartitionResponse) SetBody(v *DeleteDataLakePartitionResponseBody) *DeleteDataLakePartitionResponse {
	s.Body = v
	return s
}

func (s *DeleteDataLakePartitionResponse) Validate() error {
	return dara.Validate(s)
}
