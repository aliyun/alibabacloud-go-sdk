// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserTagValueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserTagValueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserTagValueResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserTagValueResponseBody) *UpdateUserTagValueResponse
	GetBody() *UpdateUserTagValueResponseBody
}

type UpdateUserTagValueResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserTagValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserTagValueResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserTagValueResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserTagValueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserTagValueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserTagValueResponse) GetBody() *UpdateUserTagValueResponseBody {
	return s.Body
}

func (s *UpdateUserTagValueResponse) SetHeaders(v map[string]*string) *UpdateUserTagValueResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserTagValueResponse) SetStatusCode(v int32) *UpdateUserTagValueResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserTagValueResponse) SetBody(v *UpdateUserTagValueResponseBody) *UpdateUserTagValueResponse {
	s.Body = v
	return s
}

func (s *UpdateUserTagValueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
