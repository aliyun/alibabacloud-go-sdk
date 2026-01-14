// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAclAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAclAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAclAttributeResponseBody) *UpdateAclAttributeResponse
	GetBody() *UpdateAclAttributeResponseBody
}

type UpdateAclAttributeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAclAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAclAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAclAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAclAttributeResponse) GetBody() *UpdateAclAttributeResponseBody {
	return s.Body
}

func (s *UpdateAclAttributeResponse) SetHeaders(v map[string]*string) *UpdateAclAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclAttributeResponse) SetStatusCode(v int32) *UpdateAclAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAclAttributeResponse) SetBody(v *UpdateAclAttributeResponseBody) *UpdateAclAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateAclAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
