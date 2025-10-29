// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RerunTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RerunTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *RerunTaskInstancesResponseBody) *RerunTaskInstancesResponse
	GetBody() *RerunTaskInstancesResponseBody
}

type RerunTaskInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RerunTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RerunTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s RerunTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *RerunTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RerunTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RerunTaskInstancesResponse) GetBody() *RerunTaskInstancesResponseBody {
	return s.Body
}

func (s *RerunTaskInstancesResponse) SetHeaders(v map[string]*string) *RerunTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *RerunTaskInstancesResponse) SetStatusCode(v int32) *RerunTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunTaskInstancesResponse) SetBody(v *RerunTaskInstancesResponseBody) *RerunTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *RerunTaskInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
