// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolicyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolicyVersionResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolicyVersionResponseBody) *DeletePolicyVersionResponse
	GetBody() *DeletePolicyVersionResponseBody
}

type DeletePolicyVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolicyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolicyVersionResponse) GetBody() *DeletePolicyVersionResponseBody {
	return s.Body
}

func (s *DeletePolicyVersionResponse) SetHeaders(v map[string]*string) *DeletePolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyVersionResponse) SetStatusCode(v int32) *DeletePolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyVersionResponse) SetBody(v *DeletePolicyVersionResponseBody) *DeletePolicyVersionResponse {
	s.Body = v
	return s
}

func (s *DeletePolicyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
