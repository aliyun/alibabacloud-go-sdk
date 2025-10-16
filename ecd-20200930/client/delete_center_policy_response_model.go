// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenterPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenterPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenterPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenterPolicyResponseBody) *DeleteCenterPolicyResponse
	GetBody() *DeleteCenterPolicyResponseBody
}

type DeleteCenterPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenterPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenterPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenterPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenterPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenterPolicyResponse) GetBody() *DeleteCenterPolicyResponseBody {
	return s.Body
}

func (s *DeleteCenterPolicyResponse) SetHeaders(v map[string]*string) *DeleteCenterPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenterPolicyResponse) SetStatusCode(v int32) *DeleteCenterPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenterPolicyResponse) SetBody(v *DeleteCenterPolicyResponseBody) *DeleteCenterPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteCenterPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
