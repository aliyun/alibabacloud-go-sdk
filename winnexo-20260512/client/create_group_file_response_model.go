// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupFileResponseBody) *CreateGroupFileResponse
	GetBody() *CreateGroupFileResponseBody
}

type CreateGroupFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFileResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupFileResponse) GetBody() *CreateGroupFileResponseBody {
	return s.Body
}

func (s *CreateGroupFileResponse) SetHeaders(v map[string]*string) *CreateGroupFileResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupFileResponse) SetStatusCode(v int32) *CreateGroupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupFileResponse) SetBody(v *CreateGroupFileResponseBody) *CreateGroupFileResponse {
	s.Body = v
	return s
}

func (s *CreateGroupFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
