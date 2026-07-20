// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceLinkedRoleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceLinkedRoleStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceLinkedRoleStatusResponseBody) *GetServiceLinkedRoleStatusResponse
	GetBody() *GetServiceLinkedRoleStatusResponseBody
}

type GetServiceLinkedRoleStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceLinkedRoleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceLinkedRoleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceLinkedRoleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceLinkedRoleStatusResponse) GetBody() *GetServiceLinkedRoleStatusResponseBody {
	return s.Body
}

func (s *GetServiceLinkedRoleStatusResponse) SetHeaders(v map[string]*string) *GetServiceLinkedRoleStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceLinkedRoleStatusResponse) SetStatusCode(v int32) *GetServiceLinkedRoleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceLinkedRoleStatusResponse) SetBody(v *GetServiceLinkedRoleStatusResponseBody) *GetServiceLinkedRoleStatusResponse {
	s.Body = v
	return s
}

func (s *GetServiceLinkedRoleStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
