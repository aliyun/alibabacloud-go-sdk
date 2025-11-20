// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyBindingResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicyBindingResponseBody) *DeletePolicyBindingResponse
	GetBody() *DeletePolicyBindingResponseBody
}

type DeletePolicyBindingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyBindingResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyBindingResponse) GetBody() *DeletePolicyBindingResponseBody {
	return s.Body
}

func (s *DeletePolicyBindingResponse) SetHeaders(v map[string]*string) *DeletePolicyBindingResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyBindingResponse) SetStatusCode(v int32) *DeletePolicyBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyBindingResponse) SetBody(v *DeletePolicyBindingResponseBody) *DeletePolicyBindingResponse {
	s.Body = v
	return s
}

func (s *DeletePolicyBindingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
