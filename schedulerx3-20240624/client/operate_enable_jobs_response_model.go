// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateEnableJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateEnableJobsResponse
	GetStatusCode() *int32
	SetBody(v *OperateEnableJobsResponseBody) *OperateEnableJobsResponse
	GetBody() *OperateEnableJobsResponseBody
}

type OperateEnableJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateEnableJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateEnableJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableJobsResponse) GoString() string {
	return s.String()
}

func (s *OperateEnableJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateEnableJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateEnableJobsResponse) GetBody() *OperateEnableJobsResponseBody {
	return s.Body
}

func (s *OperateEnableJobsResponse) SetHeaders(v map[string]*string) *OperateEnableJobsResponse {
	s.Headers = v
	return s
}

func (s *OperateEnableJobsResponse) SetStatusCode(v int32) *OperateEnableJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateEnableJobsResponse) SetBody(v *OperateEnableJobsResponseBody) *OperateEnableJobsResponse {
	s.Body = v
	return s
}

func (s *OperateEnableJobsResponse) Validate() error {
	return dara.Validate(s)
}
