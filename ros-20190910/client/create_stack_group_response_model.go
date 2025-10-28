// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStackGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStackGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateStackGroupResponseBody) *CreateStackGroupResponse
	GetBody() *CreateStackGroupResponseBody
}

type CreateStackGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStackGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateStackGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStackGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStackGroupResponse) GetBody() *CreateStackGroupResponseBody {
	return s.Body
}

func (s *CreateStackGroupResponse) SetHeaders(v map[string]*string) *CreateStackGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateStackGroupResponse) SetStatusCode(v int32) *CreateStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStackGroupResponse) SetBody(v *CreateStackGroupResponseBody) *CreateStackGroupResponse {
	s.Body = v
	return s
}

func (s *CreateStackGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
