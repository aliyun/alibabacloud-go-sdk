// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABMetricGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateABMetricGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateABMetricGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateABMetricGroupResponseBody) *UpdateABMetricGroupResponse
	GetBody() *UpdateABMetricGroupResponseBody
}

type UpdateABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABMetricGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateABMetricGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateABMetricGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateABMetricGroupResponse) GetBody() *UpdateABMetricGroupResponseBody {
	return s.Body
}

func (s *UpdateABMetricGroupResponse) SetHeaders(v map[string]*string) *UpdateABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateABMetricGroupResponse) SetStatusCode(v int32) *UpdateABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABMetricGroupResponse) SetBody(v *UpdateABMetricGroupResponseBody) *UpdateABMetricGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateABMetricGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
