// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadInstanceMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGadInstanceMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGadInstanceMemberResponse
	GetStatusCode() *int32
	SetBody(v *CreateGadInstanceMemberResponseBody) *CreateGadInstanceMemberResponse
	GetBody() *CreateGadInstanceMemberResponseBody
}

type CreateGadInstanceMemberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGadInstanceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGadInstanceMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGadInstanceMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGadInstanceMemberResponse) GetBody() *CreateGadInstanceMemberResponseBody {
	return s.Body
}

func (s *CreateGadInstanceMemberResponse) SetHeaders(v map[string]*string) *CreateGadInstanceMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateGadInstanceMemberResponse) SetStatusCode(v int32) *CreateGadInstanceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGadInstanceMemberResponse) SetBody(v *CreateGadInstanceMemberResponseBody) *CreateGadInstanceMemberResponse {
	s.Body = v
	return s
}

func (s *CreateGadInstanceMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
