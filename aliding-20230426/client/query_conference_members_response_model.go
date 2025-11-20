// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConferenceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConferenceMembersResponse
	GetStatusCode() *int32
	SetBody(v *QueryConferenceMembersResponseBody) *QueryConferenceMembersResponse
	GetBody() *QueryConferenceMembersResponseBody
}

type QueryConferenceMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConferenceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConferenceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConferenceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConferenceMembersResponse) GetBody() *QueryConferenceMembersResponseBody {
	return s.Body
}

func (s *QueryConferenceMembersResponse) SetHeaders(v map[string]*string) *QueryConferenceMembersResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceMembersResponse) SetStatusCode(v int32) *QueryConferenceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConferenceMembersResponse) SetBody(v *QueryConferenceMembersResponseBody) *QueryConferenceMembersResponse {
	s.Body = v
	return s
}

func (s *QueryConferenceMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
