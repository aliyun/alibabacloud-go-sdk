// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTenantMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTenantMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTenantMemberResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTenantMemberResponseBody) *RemoveTenantMemberResponse
	GetBody() *RemoveTenantMemberResponseBody
}

type RemoveTenantMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTenantMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTenantMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTenantMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTenantMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTenantMemberResponse) GetBody() *RemoveTenantMemberResponseBody {
	return s.Body
}

func (s *RemoveTenantMemberResponse) SetHeaders(v map[string]*string) *RemoveTenantMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveTenantMemberResponse) SetStatusCode(v int32) *RemoveTenantMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTenantMemberResponse) SetBody(v *RemoveTenantMemberResponseBody) *RemoveTenantMemberResponse {
	s.Body = v
	return s
}

func (s *RemoveTenantMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
