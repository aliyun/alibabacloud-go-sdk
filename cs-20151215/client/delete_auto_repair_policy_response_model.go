// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoRepairPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoRepairPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoRepairPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoRepairPolicyResponseBody) *DeleteAutoRepairPolicyResponse
	GetBody() *DeleteAutoRepairPolicyResponseBody
}

type DeleteAutoRepairPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoRepairPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoRepairPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoRepairPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoRepairPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoRepairPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoRepairPolicyResponse) GetBody() *DeleteAutoRepairPolicyResponseBody {
	return s.Body
}

func (s *DeleteAutoRepairPolicyResponse) SetHeaders(v map[string]*string) *DeleteAutoRepairPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoRepairPolicyResponse) SetStatusCode(v int32) *DeleteAutoRepairPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoRepairPolicyResponse) SetBody(v *DeleteAutoRepairPolicyResponseBody) *DeleteAutoRepairPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoRepairPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
