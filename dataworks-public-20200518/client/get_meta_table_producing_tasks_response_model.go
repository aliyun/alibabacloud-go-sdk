// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableProducingTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableProducingTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableProducingTasksResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableProducingTasksResponseBody) *GetMetaTableProducingTasksResponse
	GetBody() *GetMetaTableProducingTasksResponseBody
}

type GetMetaTableProducingTasksResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableProducingTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableProducingTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableProducingTasksResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableProducingTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableProducingTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableProducingTasksResponse) GetBody() *GetMetaTableProducingTasksResponseBody {
	return s.Body
}

func (s *GetMetaTableProducingTasksResponse) SetHeaders(v map[string]*string) *GetMetaTableProducingTasksResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableProducingTasksResponse) SetStatusCode(v int32) *GetMetaTableProducingTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableProducingTasksResponse) SetBody(v *GetMetaTableProducingTasksResponseBody) *GetMetaTableProducingTasksResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableProducingTasksResponse) Validate() error {
	return dara.Validate(s)
}
