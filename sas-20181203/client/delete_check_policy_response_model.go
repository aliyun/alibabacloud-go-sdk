// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCheckPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCheckPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCheckPolicyResponseBody) *DeleteCheckPolicyResponse
	GetBody() *DeleteCheckPolicyResponseBody
}

type DeleteCheckPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCheckPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCheckPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCheckPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCheckPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCheckPolicyResponse) GetBody() *DeleteCheckPolicyResponseBody {
	return s.Body
}

func (s *DeleteCheckPolicyResponse) SetHeaders(v map[string]*string) *DeleteCheckPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCheckPolicyResponse) SetStatusCode(v int32) *DeleteCheckPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCheckPolicyResponse) SetBody(v *DeleteCheckPolicyResponseBody) *DeleteCheckPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteCheckPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
