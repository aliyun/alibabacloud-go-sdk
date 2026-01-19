// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeResourceServerScopesFromGroupResponse
	GetStatusCode() *int32
	SetBody(v *RevokeResourceServerScopesFromGroupResponseBody) *RevokeResourceServerScopesFromGroupResponse
	GetBody() *RevokeResourceServerScopesFromGroupResponseBody
}

type RevokeResourceServerScopesFromGroupResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResourceServerScopesFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResourceServerScopesFromGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeResourceServerScopesFromGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeResourceServerScopesFromGroupResponse) GetBody() *RevokeResourceServerScopesFromGroupResponseBody {
	return s.Body
}

func (s *RevokeResourceServerScopesFromGroupResponse) SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RevokeResourceServerScopesFromGroupResponse) SetStatusCode(v int32) *RevokeResourceServerScopesFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResourceServerScopesFromGroupResponse) SetBody(v *RevokeResourceServerScopesFromGroupResponseBody) *RevokeResourceServerScopesFromGroupResponse {
	s.Body = v
	return s
}

func (s *RevokeResourceServerScopesFromGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
