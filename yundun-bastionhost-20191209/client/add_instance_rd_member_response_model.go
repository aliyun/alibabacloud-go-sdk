// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceRdMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddInstanceRdMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddInstanceRdMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddInstanceRdMemberResponseBody) *AddInstanceRdMemberResponse
	GetBody() *AddInstanceRdMemberResponseBody
}

type AddInstanceRdMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInstanceRdMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInstanceRdMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceRdMemberResponse) GoString() string {
	return s.String()
}

func (s *AddInstanceRdMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddInstanceRdMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddInstanceRdMemberResponse) GetBody() *AddInstanceRdMemberResponseBody {
	return s.Body
}

func (s *AddInstanceRdMemberResponse) SetHeaders(v map[string]*string) *AddInstanceRdMemberResponse {
	s.Headers = v
	return s
}

func (s *AddInstanceRdMemberResponse) SetStatusCode(v int32) *AddInstanceRdMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInstanceRdMemberResponse) SetBody(v *AddInstanceRdMemberResponseBody) *AddInstanceRdMemberResponse {
	s.Body = v
	return s
}

func (s *AddInstanceRdMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
