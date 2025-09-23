// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceShareInvitationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceShareInvitationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceShareInvitationsResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceShareInvitationsResponseBody) *ListResourceShareInvitationsResponse
	GetBody() *ListResourceShareInvitationsResponseBody
}

type ListResourceShareInvitationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceShareInvitationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceShareInvitationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceShareInvitationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceShareInvitationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceShareInvitationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceShareInvitationsResponse) GetBody() *ListResourceShareInvitationsResponseBody {
	return s.Body
}

func (s *ListResourceShareInvitationsResponse) SetHeaders(v map[string]*string) *ListResourceShareInvitationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceShareInvitationsResponse) SetStatusCode(v int32) *ListResourceShareInvitationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceShareInvitationsResponse) SetBody(v *ListResourceShareInvitationsResponseBody) *ListResourceShareInvitationsResponse {
	s.Body = v
	return s
}

func (s *ListResourceShareInvitationsResponse) Validate() error {
	return dara.Validate(s)
}
