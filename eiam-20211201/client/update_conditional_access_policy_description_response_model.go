// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConditionalAccessPolicyDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConditionalAccessPolicyDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConditionalAccessPolicyDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConditionalAccessPolicyDescriptionResponseBody) *UpdateConditionalAccessPolicyDescriptionResponse
	GetBody() *UpdateConditionalAccessPolicyDescriptionResponseBody
}

type UpdateConditionalAccessPolicyDescriptionResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConditionalAccessPolicyDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConditionalAccessPolicyDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConditionalAccessPolicyDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConditionalAccessPolicyDescriptionResponse) GetBody() *UpdateConditionalAccessPolicyDescriptionResponseBody {
	return s.Body
}

func (s *UpdateConditionalAccessPolicyDescriptionResponse) SetHeaders(v map[string]*string) *UpdateConditionalAccessPolicyDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionResponse) SetStatusCode(v int32) *UpdateConditionalAccessPolicyDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionResponse) SetBody(v *UpdateConditionalAccessPolicyDescriptionResponseBody) *UpdateConditionalAccessPolicyDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
