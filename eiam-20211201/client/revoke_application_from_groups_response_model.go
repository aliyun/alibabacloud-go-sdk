// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeApplicationFromGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeApplicationFromGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeApplicationFromGroupsResponse
	GetStatusCode() *int32
	SetBody(v *RevokeApplicationFromGroupsResponseBody) *RevokeApplicationFromGroupsResponse
	GetBody() *RevokeApplicationFromGroupsResponseBody
}

type RevokeApplicationFromGroupsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeApplicationFromGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeApplicationFromGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeApplicationFromGroupsResponse) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeApplicationFromGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeApplicationFromGroupsResponse) GetBody() *RevokeApplicationFromGroupsResponseBody {
	return s.Body
}

func (s *RevokeApplicationFromGroupsResponse) SetHeaders(v map[string]*string) *RevokeApplicationFromGroupsResponse {
	s.Headers = v
	return s
}

func (s *RevokeApplicationFromGroupsResponse) SetStatusCode(v int32) *RevokeApplicationFromGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeApplicationFromGroupsResponse) SetBody(v *RevokeApplicationFromGroupsResponseBody) *RevokeApplicationFromGroupsResponse {
	s.Body = v
	return s
}

func (s *RevokeApplicationFromGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
