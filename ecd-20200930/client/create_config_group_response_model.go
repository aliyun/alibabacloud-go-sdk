// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigGroupResponseBody) *CreateConfigGroupResponse
	GetBody() *CreateConfigGroupResponseBody
}

type CreateConfigGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigGroupResponse) GetBody() *CreateConfigGroupResponseBody {
	return s.Body
}

func (s *CreateConfigGroupResponse) SetHeaders(v map[string]*string) *CreateConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigGroupResponse) SetStatusCode(v int32) *CreateConfigGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigGroupResponse) SetBody(v *CreateConfigGroupResponseBody) *CreateConfigGroupResponse {
	s.Body = v
	return s
}

func (s *CreateConfigGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
