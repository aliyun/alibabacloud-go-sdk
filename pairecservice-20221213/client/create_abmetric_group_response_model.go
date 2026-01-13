// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABMetricGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateABMetricGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateABMetricGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateABMetricGroupResponseBody) *CreateABMetricGroupResponse
	GetBody() *CreateABMetricGroupResponseBody
}

type CreateABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateABMetricGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateABMetricGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateABMetricGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateABMetricGroupResponse) GetBody() *CreateABMetricGroupResponseBody {
	return s.Body
}

func (s *CreateABMetricGroupResponse) SetHeaders(v map[string]*string) *CreateABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateABMetricGroupResponse) SetStatusCode(v int32) *CreateABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABMetricGroupResponse) SetBody(v *CreateABMetricGroupResponseBody) *CreateABMetricGroupResponse {
	s.Body = v
	return s
}

func (s *CreateABMetricGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
