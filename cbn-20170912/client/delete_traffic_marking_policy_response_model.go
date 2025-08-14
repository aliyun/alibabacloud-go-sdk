// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMarkingPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficMarkingPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficMarkingPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficMarkingPolicyResponseBody) *DeleteTrafficMarkingPolicyResponse
	GetBody() *DeleteTrafficMarkingPolicyResponseBody
}

type DeleteTrafficMarkingPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficMarkingPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficMarkingPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMarkingPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMarkingPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficMarkingPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficMarkingPolicyResponse) GetBody() *DeleteTrafficMarkingPolicyResponseBody {
	return s.Body
}

func (s *DeleteTrafficMarkingPolicyResponse) SetHeaders(v map[string]*string) *DeleteTrafficMarkingPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficMarkingPolicyResponse) SetStatusCode(v int32) *DeleteTrafficMarkingPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficMarkingPolicyResponse) SetBody(v *DeleteTrafficMarkingPolicyResponseBody) *DeleteTrafficMarkingPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficMarkingPolicyResponse) Validate() error {
	return dara.Validate(s)
}
