// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakePartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataLakePartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataLakePartitionResponse
	GetStatusCode() *int32
	SetBody(v *GetDataLakePartitionResponseBody) *GetDataLakePartitionResponse
	GetBody() *GetDataLakePartitionResponseBody
}

type GetDataLakePartitionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataLakePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataLakePartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakePartitionResponse) GoString() string {
	return s.String()
}

func (s *GetDataLakePartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataLakePartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataLakePartitionResponse) GetBody() *GetDataLakePartitionResponseBody {
	return s.Body
}

func (s *GetDataLakePartitionResponse) SetHeaders(v map[string]*string) *GetDataLakePartitionResponse {
	s.Headers = v
	return s
}

func (s *GetDataLakePartitionResponse) SetStatusCode(v int32) *GetDataLakePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataLakePartitionResponse) SetBody(v *GetDataLakePartitionResponseBody) *GetDataLakePartitionResponse {
	s.Body = v
	return s
}

func (s *GetDataLakePartitionResponse) Validate() error {
	return dara.Validate(s)
}
