// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomGroupResponseBody) *CreateCustomGroupResponse
	GetBody() *CreateCustomGroupResponseBody
}

type CreateCustomGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomGroupResponse) GetBody() *CreateCustomGroupResponseBody {
	return s.Body
}

func (s *CreateCustomGroupResponse) SetHeaders(v map[string]*string) *CreateCustomGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomGroupResponse) SetStatusCode(v int32) *CreateCustomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomGroupResponse) SetBody(v *CreateCustomGroupResponseBody) *CreateCustomGroupResponse {
	s.Body = v
	return s
}

func (s *CreateCustomGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
