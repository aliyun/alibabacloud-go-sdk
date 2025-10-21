// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWordGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWordGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWordGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateWordGroupResponseBody) *CreateWordGroupResponse
	GetBody() *CreateWordGroupResponseBody
}

type CreateWordGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWordGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWordGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWordGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateWordGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWordGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWordGroupResponse) GetBody() *CreateWordGroupResponseBody {
	return s.Body
}

func (s *CreateWordGroupResponse) SetHeaders(v map[string]*string) *CreateWordGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateWordGroupResponse) SetStatusCode(v int32) *CreateWordGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWordGroupResponse) SetBody(v *CreateWordGroupResponseBody) *CreateWordGroupResponse {
	s.Body = v
	return s
}

func (s *CreateWordGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
