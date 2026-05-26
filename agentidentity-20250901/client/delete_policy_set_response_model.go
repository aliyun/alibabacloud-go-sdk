// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicySetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicySetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicySetResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicySetResponseBody) *DeletePolicySetResponse
	GetBody() *DeletePolicySetResponseBody
}

type DeletePolicySetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicySetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicySetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicySetResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicySetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicySetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicySetResponse) GetBody() *DeletePolicySetResponseBody {
	return s.Body
}

func (s *DeletePolicySetResponse) SetHeaders(v map[string]*string) *DeletePolicySetResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicySetResponse) SetStatusCode(v int32) *DeletePolicySetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicySetResponse) SetBody(v *DeletePolicySetResponseBody) *DeletePolicySetResponse {
	s.Body = v
	return s
}

func (s *DeletePolicySetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
