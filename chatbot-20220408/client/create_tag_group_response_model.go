// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateTagGroupResponseBody) *CreateTagGroupResponse
	GetBody() *CreateTagGroupResponseBody
}

type CreateTagGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTagGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTagGroupResponse) GetBody() *CreateTagGroupResponseBody {
	return s.Body
}

func (s *CreateTagGroupResponse) SetHeaders(v map[string]*string) *CreateTagGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateTagGroupResponse) SetStatusCode(v int32) *CreateTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagGroupResponse) SetBody(v *CreateTagGroupResponseBody) *CreateTagGroupResponse {
	s.Body = v
	return s
}

func (s *CreateTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
