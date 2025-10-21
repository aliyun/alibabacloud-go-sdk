// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemberResponseBody) *DeleteMemberResponse
	GetBody() *DeleteMemberResponseBody
}

type DeleteMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemberResponse) GetBody() *DeleteMemberResponseBody {
	return s.Body
}

func (s *DeleteMemberResponse) SetHeaders(v map[string]*string) *DeleteMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemberResponse) SetStatusCode(v int32) *DeleteMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemberResponse) SetBody(v *DeleteMemberResponseBody) *DeleteMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
