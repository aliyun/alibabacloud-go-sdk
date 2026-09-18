// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupTextResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupTextResponseBody) *CreateGroupTextResponse
	GetBody() *CreateGroupTextResponseBody
}

type CreateGroupTextResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupTextResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupTextResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupTextResponse) GetBody() *CreateGroupTextResponseBody {
	return s.Body
}

func (s *CreateGroupTextResponse) SetHeaders(v map[string]*string) *CreateGroupTextResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupTextResponse) SetStatusCode(v int32) *CreateGroupTextResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupTextResponse) SetBody(v *CreateGroupTextResponseBody) *CreateGroupTextResponse {
	s.Body = v
	return s
}

func (s *CreateGroupTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
