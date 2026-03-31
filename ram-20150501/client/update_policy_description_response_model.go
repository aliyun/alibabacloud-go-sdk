// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolicyDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolicyDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolicyDescriptionResponseBody) *UpdatePolicyDescriptionResponse
	GetBody() *UpdatePolicyDescriptionResponseBody
}

type UpdatePolicyDescriptionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolicyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolicyDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolicyDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolicyDescriptionResponse) GetBody() *UpdatePolicyDescriptionResponseBody {
	return s.Body
}

func (s *UpdatePolicyDescriptionResponse) SetHeaders(v map[string]*string) *UpdatePolicyDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolicyDescriptionResponse) SetStatusCode(v int32) *UpdatePolicyDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolicyDescriptionResponse) SetBody(v *UpdatePolicyDescriptionResponseBody) *UpdatePolicyDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdatePolicyDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
