// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainGroupResponseBody) *UpdateDomainGroupResponse
	GetBody() *UpdateDomainGroupResponseBody
}

type UpdateDomainGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainGroupResponse) GetBody() *UpdateDomainGroupResponseBody {
	return s.Body
}

func (s *UpdateDomainGroupResponse) SetHeaders(v map[string]*string) *UpdateDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainGroupResponse) SetStatusCode(v int32) *UpdateDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainGroupResponse) SetBody(v *UpdateDomainGroupResponseBody) *UpdateDomainGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
