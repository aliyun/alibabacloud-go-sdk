// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLCollectorPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySQLCollectorPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySQLCollectorPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifySQLCollectorPolicyResponseBody) *ModifySQLCollectorPolicyResponse
	GetBody() *ModifySQLCollectorPolicyResponseBody
}

type ModifySQLCollectorPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySQLCollectorPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySQLCollectorPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySQLCollectorPolicyResponse) GetBody() *ModifySQLCollectorPolicyResponseBody {
	return s.Body
}

func (s *ModifySQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *ModifySQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifySQLCollectorPolicyResponse) SetStatusCode(v int32) *ModifySQLCollectorPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySQLCollectorPolicyResponse) SetBody(v *ModifySQLCollectorPolicyResponseBody) *ModifySQLCollectorPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifySQLCollectorPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
