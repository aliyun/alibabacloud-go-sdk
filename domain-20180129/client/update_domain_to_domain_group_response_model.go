// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainToDomainGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainToDomainGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainToDomainGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainToDomainGroupResponseBody) *UpdateDomainToDomainGroupResponse
	GetBody() *UpdateDomainToDomainGroupResponseBody
}

type UpdateDomainToDomainGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainToDomainGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainToDomainGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainToDomainGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainToDomainGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainToDomainGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainToDomainGroupResponse) GetBody() *UpdateDomainToDomainGroupResponseBody {
	return s.Body
}

func (s *UpdateDomainToDomainGroupResponse) SetHeaders(v map[string]*string) *UpdateDomainToDomainGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainToDomainGroupResponse) SetStatusCode(v int32) *UpdateDomainToDomainGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainToDomainGroupResponse) SetBody(v *UpdateDomainToDomainGroupResponseBody) *UpdateDomainToDomainGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainToDomainGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
