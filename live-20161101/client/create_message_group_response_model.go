// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateMessageGroupResponseBody) *CreateMessageGroupResponse
	GetBody() *CreateMessageGroupResponseBody
}

type CreateMessageGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMessageGroupResponse) GetBody() *CreateMessageGroupResponseBody {
	return s.Body
}

func (s *CreateMessageGroupResponse) SetHeaders(v map[string]*string) *CreateMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageGroupResponse) SetStatusCode(v int32) *CreateMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageGroupResponse) SetBody(v *CreateMessageGroupResponseBody) *CreateMessageGroupResponse {
	s.Body = v
	return s
}

func (s *CreateMessageGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
