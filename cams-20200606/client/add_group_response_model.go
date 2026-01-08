// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddGroupResponseBody) *AddGroupResponse
	GetBody() *AddGroupResponseBody
}

type AddGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGroupResponse) GoString() string {
	return s.String()
}

func (s *AddGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGroupResponse) GetBody() *AddGroupResponseBody {
	return s.Body
}

func (s *AddGroupResponse) SetHeaders(v map[string]*string) *AddGroupResponse {
	s.Headers = v
	return s
}

func (s *AddGroupResponse) SetStatusCode(v int32) *AddGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGroupResponse) SetBody(v *AddGroupResponseBody) *AddGroupResponse {
	s.Body = v
	return s
}

func (s *AddGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
