// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstancePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstancePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstancePolicyResponseBody) *DeleteInstancePolicyResponse
	GetBody() *DeleteInstancePolicyResponseBody
}

type DeleteInstancePolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstancePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstancePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstancePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstancePolicyResponse) GetBody() *DeleteInstancePolicyResponseBody {
	return s.Body
}

func (s *DeleteInstancePolicyResponse) SetHeaders(v map[string]*string) *DeleteInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstancePolicyResponse) SetStatusCode(v int32) *DeleteInstancePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstancePolicyResponse) SetBody(v *DeleteInstancePolicyResponseBody) *DeleteInstancePolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteInstancePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
