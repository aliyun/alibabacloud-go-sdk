// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateContactGroupResponseBody) *CreateContactGroupResponse
	GetBody() *CreateContactGroupResponseBody
}

type CreateContactGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContactGroupResponse) GetBody() *CreateContactGroupResponseBody {
	return s.Body
}

func (s *CreateContactGroupResponse) SetHeaders(v map[string]*string) *CreateContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateContactGroupResponse) SetStatusCode(v int32) *CreateContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContactGroupResponse) SetBody(v *CreateContactGroupResponseBody) *CreateContactGroupResponse {
	s.Body = v
	return s
}

func (s *CreateContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
