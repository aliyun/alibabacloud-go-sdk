// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicyInstanceResponseBody) *DeletePolicyInstanceResponse
	GetBody() *DeletePolicyInstanceResponseBody
}

type DeletePolicyInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyInstanceResponse) GetBody() *DeletePolicyInstanceResponseBody {
	return s.Body
}

func (s *DeletePolicyInstanceResponse) SetHeaders(v map[string]*string) *DeletePolicyInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyInstanceResponse) SetStatusCode(v int32) *DeletePolicyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyInstanceResponse) SetBody(v *DeletePolicyInstanceResponseBody) *DeletePolicyInstanceResponse {
	s.Body = v
	return s
}

func (s *DeletePolicyInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
