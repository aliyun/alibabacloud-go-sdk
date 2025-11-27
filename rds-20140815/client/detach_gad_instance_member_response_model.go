// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGadInstanceMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachGadInstanceMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachGadInstanceMemberResponse
	GetStatusCode() *int32
	SetBody(v *DetachGadInstanceMemberResponseBody) *DetachGadInstanceMemberResponse
	GetBody() *DetachGadInstanceMemberResponseBody
}

type DetachGadInstanceMemberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachGadInstanceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachGadInstanceMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachGadInstanceMemberResponse) GoString() string {
	return s.String()
}

func (s *DetachGadInstanceMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachGadInstanceMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachGadInstanceMemberResponse) GetBody() *DetachGadInstanceMemberResponseBody {
	return s.Body
}

func (s *DetachGadInstanceMemberResponse) SetHeaders(v map[string]*string) *DetachGadInstanceMemberResponse {
	s.Headers = v
	return s
}

func (s *DetachGadInstanceMemberResponse) SetStatusCode(v int32) *DetachGadInstanceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachGadInstanceMemberResponse) SetBody(v *DetachGadInstanceMemberResponseBody) *DetachGadInstanceMemberResponse {
	s.Body = v
	return s
}

func (s *DetachGadInstanceMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
