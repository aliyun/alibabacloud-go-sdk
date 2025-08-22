// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateDisableJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateDisableJobsResponse
	GetStatusCode() *int32
	SetBody(v *OperateDisableJobsResponseBody) *OperateDisableJobsResponse
	GetBody() *OperateDisableJobsResponseBody
}

type OperateDisableJobsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateDisableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateDisableJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableJobsResponse) GoString() string {
	return s.String()
}

func (s *OperateDisableJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateDisableJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateDisableJobsResponse) GetBody() *OperateDisableJobsResponseBody {
	return s.Body
}

func (s *OperateDisableJobsResponse) SetHeaders(v map[string]*string) *OperateDisableJobsResponse {
	s.Headers = v
	return s
}

func (s *OperateDisableJobsResponse) SetStatusCode(v int32) *OperateDisableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateDisableJobsResponse) SetBody(v *OperateDisableJobsResponseBody) *OperateDisableJobsResponse {
	s.Body = v
	return s
}

func (s *OperateDisableJobsResponse) Validate() error {
	return dara.Validate(s)
}
