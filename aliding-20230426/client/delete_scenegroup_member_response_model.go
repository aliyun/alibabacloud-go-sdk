// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenegroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScenegroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScenegroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScenegroupMemberResponseBody) *DeleteScenegroupMemberResponse
	GetBody() *DeleteScenegroupMemberResponseBody
}

type DeleteScenegroupMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScenegroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScenegroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenegroupMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteScenegroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScenegroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScenegroupMemberResponse) GetBody() *DeleteScenegroupMemberResponseBody {
	return s.Body
}

func (s *DeleteScenegroupMemberResponse) SetHeaders(v map[string]*string) *DeleteScenegroupMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteScenegroupMemberResponse) SetStatusCode(v int32) *DeleteScenegroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScenegroupMemberResponse) SetBody(v *DeleteScenegroupMemberResponseBody) *DeleteScenegroupMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteScenegroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
