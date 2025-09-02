// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTablePartitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTablePartitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTablePartitionResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTablePartitionResponseBody) *GetMetaTablePartitionResponse
	GetBody() *GetMetaTablePartitionResponseBody
}

type GetMetaTablePartitionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTablePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTablePartitionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTablePartitionResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTablePartitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTablePartitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTablePartitionResponse) GetBody() *GetMetaTablePartitionResponseBody {
	return s.Body
}

func (s *GetMetaTablePartitionResponse) SetHeaders(v map[string]*string) *GetMetaTablePartitionResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTablePartitionResponse) SetStatusCode(v int32) *GetMetaTablePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTablePartitionResponse) SetBody(v *GetMetaTablePartitionResponseBody) *GetMetaTablePartitionResponse {
	s.Body = v
	return s
}

func (s *GetMetaTablePartitionResponse) Validate() error {
	return dara.Validate(s)
}
