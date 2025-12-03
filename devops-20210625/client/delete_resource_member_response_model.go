// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceMemberResponseBody) *DeleteResourceMemberResponse
	GetBody() *DeleteResourceMemberResponseBody
}

type DeleteResourceMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceMemberResponse) GetBody() *DeleteResourceMemberResponseBody {
	return s.Body
}

func (s *DeleteResourceMemberResponse) SetHeaders(v map[string]*string) *DeleteResourceMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceMemberResponse) SetStatusCode(v int32) *DeleteResourceMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceMemberResponse) SetBody(v *DeleteResourceMemberResponseBody) *DeleteResourceMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
