// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGadInstanceDbMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachGadInstanceDbMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachGadInstanceDbMemberResponse
	GetStatusCode() *int32
	SetBody(v *DetachGadInstanceDbMemberResponseBody) *DetachGadInstanceDbMemberResponse
	GetBody() *DetachGadInstanceDbMemberResponseBody
}

type DetachGadInstanceDbMemberResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachGadInstanceDbMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachGadInstanceDbMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachGadInstanceDbMemberResponse) GoString() string {
	return s.String()
}

func (s *DetachGadInstanceDbMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachGadInstanceDbMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachGadInstanceDbMemberResponse) GetBody() *DetachGadInstanceDbMemberResponseBody {
	return s.Body
}

func (s *DetachGadInstanceDbMemberResponse) SetHeaders(v map[string]*string) *DetachGadInstanceDbMemberResponse {
	s.Headers = v
	return s
}

func (s *DetachGadInstanceDbMemberResponse) SetStatusCode(v int32) *DetachGadInstanceDbMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachGadInstanceDbMemberResponse) SetBody(v *DetachGadInstanceDbMemberResponseBody) *DetachGadInstanceDbMemberResponse {
	s.Body = v
	return s
}

func (s *DetachGadInstanceDbMemberResponse) Validate() error {
	return dara.Validate(s)
}
