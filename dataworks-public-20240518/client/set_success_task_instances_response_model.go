// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuccessTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSuccessTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSuccessTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *SetSuccessTaskInstancesResponseBody) *SetSuccessTaskInstancesResponse
	GetBody() *SetSuccessTaskInstancesResponseBody
}

type SetSuccessTaskInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSuccessTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSuccessTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSuccessTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *SetSuccessTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSuccessTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSuccessTaskInstancesResponse) GetBody() *SetSuccessTaskInstancesResponseBody {
	return s.Body
}

func (s *SetSuccessTaskInstancesResponse) SetHeaders(v map[string]*string) *SetSuccessTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *SetSuccessTaskInstancesResponse) SetStatusCode(v int32) *SetSuccessTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSuccessTaskInstancesResponse) SetBody(v *SetSuccessTaskInstancesResponseBody) *SetSuccessTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *SetSuccessTaskInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
