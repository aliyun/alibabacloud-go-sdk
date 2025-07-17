// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakePartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataLakePartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataLakePartitionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataLakePartitionResponseBody) *CreateDataLakePartitionResponse
	GetBody() *CreateDataLakePartitionResponseBody
}

type CreateDataLakePartitionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataLakePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataLakePartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakePartitionResponse) GoString() string {
	return s.String()
}

func (s *CreateDataLakePartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataLakePartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataLakePartitionResponse) GetBody() *CreateDataLakePartitionResponseBody {
	return s.Body
}

func (s *CreateDataLakePartitionResponse) SetHeaders(v map[string]*string) *CreateDataLakePartitionResponse {
	s.Headers = v
	return s
}

func (s *CreateDataLakePartitionResponse) SetStatusCode(v int32) *CreateDataLakePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataLakePartitionResponse) SetBody(v *CreateDataLakePartitionResponseBody) *CreateDataLakePartitionResponse {
	s.Body = v
	return s
}

func (s *CreateDataLakePartitionResponse) Validate() error {
	return dara.Validate(s)
}
