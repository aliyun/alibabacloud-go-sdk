// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateJobGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateJobGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateJobGroupResponseBody) *CreateJobGroupResponse
	GetBody() *CreateJobGroupResponseBody
}

type CreateJobGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateJobGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateJobGroupResponse) GetBody() *CreateJobGroupResponseBody {
	return s.Body
}

func (s *CreateJobGroupResponse) SetHeaders(v map[string]*string) *CreateJobGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateJobGroupResponse) SetStatusCode(v int32) *CreateJobGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobGroupResponse) SetBody(v *CreateJobGroupResponseBody) *CreateJobGroupResponse {
	s.Body = v
	return s
}

func (s *CreateJobGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
