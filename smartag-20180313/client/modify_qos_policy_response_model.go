// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQosPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQosPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQosPolicyResponseBody) *ModifyQosPolicyResponse
	GetBody() *ModifyQosPolicyResponseBody
}

type ModifyQosPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQosPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQosPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyQosPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQosPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQosPolicyResponse) GetBody() *ModifyQosPolicyResponseBody {
	return s.Body
}

func (s *ModifyQosPolicyResponse) SetHeaders(v map[string]*string) *ModifyQosPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyQosPolicyResponse) SetStatusCode(v int32) *ModifyQosPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQosPolicyResponse) SetBody(v *ModifyQosPolicyResponseBody) *ModifyQosPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyQosPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
