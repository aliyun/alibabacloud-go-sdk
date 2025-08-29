// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABMetricGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteABMetricGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteABMetricGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteABMetricGroupResponseBody) *DeleteABMetricGroupResponse
	GetBody() *DeleteABMetricGroupResponseBody
}

type DeleteABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteABMetricGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteABMetricGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteABMetricGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteABMetricGroupResponse) GetBody() *DeleteABMetricGroupResponseBody {
	return s.Body
}

func (s *DeleteABMetricGroupResponse) SetHeaders(v map[string]*string) *DeleteABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteABMetricGroupResponse) SetStatusCode(v int32) *DeleteABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABMetricGroupResponse) SetBody(v *DeleteABMetricGroupResponseBody) *DeleteABMetricGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteABMetricGroupResponse) Validate() error {
	return dara.Validate(s)
}
