// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppInstanceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppInstanceGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppInstanceGroupResponseBody) *CreateAppInstanceGroupResponse
	GetBody() *CreateAppInstanceGroupResponseBody
}

type CreateAppInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppInstanceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppInstanceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppInstanceGroupResponse) GetBody() *CreateAppInstanceGroupResponseBody {
	return s.Body
}

func (s *CreateAppInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetStatusCode(v int32) *CreateAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetBody(v *CreateAppInstanceGroupResponseBody) *CreateAppInstanceGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAppInstanceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
