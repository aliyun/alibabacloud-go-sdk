// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceMemberResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceMemberResponseBody) *CreateResourceMemberResponse
	GetBody() *CreateResourceMemberResponseBody
}

type CreateResourceMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceMemberResponse) GetBody() *CreateResourceMemberResponseBody {
	return s.Body
}

func (s *CreateResourceMemberResponse) SetHeaders(v map[string]*string) *CreateResourceMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceMemberResponse) SetStatusCode(v int32) *CreateResourceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceMemberResponse) SetBody(v *CreateResourceMemberResponseBody) *CreateResourceMemberResponse {
	s.Body = v
	return s
}

func (s *CreateResourceMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
