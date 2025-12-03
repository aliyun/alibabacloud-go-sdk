// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateApiGroupResponseBody) *CreateApiGroupResponse
	GetBody() *CreateApiGroupResponseBody
}

type CreateApiGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApiGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApiGroupResponse) GetBody() *CreateApiGroupResponseBody {
	return s.Body
}

func (s *CreateApiGroupResponse) SetHeaders(v map[string]*string) *CreateApiGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateApiGroupResponse) SetStatusCode(v int32) *CreateApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiGroupResponse) SetBody(v *CreateApiGroupResponseBody) *CreateApiGroupResponse {
	s.Body = v
	return s
}

func (s *CreateApiGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
