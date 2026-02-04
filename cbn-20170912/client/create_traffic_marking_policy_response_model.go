// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMarkingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrafficMarkingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrafficMarkingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrafficMarkingPolicyResponseBody) *CreateTrafficMarkingPolicyResponse
	GetBody() *CreateTrafficMarkingPolicyResponseBody
}

type CreateTrafficMarkingPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficMarkingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrafficMarkingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrafficMarkingPolicyResponse) GetBody() *CreateTrafficMarkingPolicyResponseBody {
	return s.Body
}

func (s *CreateTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *CreateTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficMarkingPolicyResponse) SetStatusCode(v int32) *CreateTrafficMarkingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficMarkingPolicyResponse) SetBody(v *CreateTrafficMarkingPolicyResponseBody) *CreateTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateTrafficMarkingPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
