// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetABMetricGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetABMetricGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetABMetricGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetABMetricGroupResponseBody) *GetABMetricGroupResponse
	GetBody() *GetABMetricGroupResponseBody
}

type GetABMetricGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetABMetricGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *GetABMetricGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetABMetricGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetABMetricGroupResponse) GetBody() *GetABMetricGroupResponseBody {
	return s.Body
}

func (s *GetABMetricGroupResponse) SetHeaders(v map[string]*string) *GetABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *GetABMetricGroupResponse) SetStatusCode(v int32) *GetABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetABMetricGroupResponse) SetBody(v *GetABMetricGroupResponseBody) *GetABMetricGroupResponse {
	s.Body = v
	return s
}

func (s *GetABMetricGroupResponse) Validate() error {
	return dara.Validate(s)
}
