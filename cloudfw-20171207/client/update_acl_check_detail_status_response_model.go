// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclCheckDetailStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAclCheckDetailStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAclCheckDetailStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAclCheckDetailStatusResponseBody) *UpdateAclCheckDetailStatusResponse
	GetBody() *UpdateAclCheckDetailStatusResponseBody
}

type UpdateAclCheckDetailStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAclCheckDetailStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAclCheckDetailStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclCheckDetailStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclCheckDetailStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAclCheckDetailStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAclCheckDetailStatusResponse) GetBody() *UpdateAclCheckDetailStatusResponseBody {
	return s.Body
}

func (s *UpdateAclCheckDetailStatusResponse) SetHeaders(v map[string]*string) *UpdateAclCheckDetailStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponse) SetStatusCode(v int32) *UpdateAclCheckDetailStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAclCheckDetailStatusResponse) SetBody(v *UpdateAclCheckDetailStatusResponseBody) *UpdateAclCheckDetailStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAclCheckDetailStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
