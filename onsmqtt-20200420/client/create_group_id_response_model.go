// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupIdResponseBody) *CreateGroupIdResponse
	GetBody() *CreateGroupIdResponseBody
}

type CreateGroupIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupIdResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupIdResponse) GetBody() *CreateGroupIdResponseBody {
	return s.Body
}

func (s *CreateGroupIdResponse) SetHeaders(v map[string]*string) *CreateGroupIdResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupIdResponse) SetStatusCode(v int32) *CreateGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupIdResponse) SetBody(v *CreateGroupIdResponseBody) *CreateGroupIdResponse {
	s.Body = v
	return s
}

func (s *CreateGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
