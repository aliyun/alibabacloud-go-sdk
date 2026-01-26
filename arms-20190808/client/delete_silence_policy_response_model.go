// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSilencePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSilencePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSilencePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSilencePolicyResponseBody) *DeleteSilencePolicyResponse
	GetBody() *DeleteSilencePolicyResponseBody
}

type DeleteSilencePolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSilencePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSilencePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSilencePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSilencePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSilencePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSilencePolicyResponse) GetBody() *DeleteSilencePolicyResponseBody {
	return s.Body
}

func (s *DeleteSilencePolicyResponse) SetHeaders(v map[string]*string) *DeleteSilencePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSilencePolicyResponse) SetStatusCode(v int32) *DeleteSilencePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSilencePolicyResponse) SetBody(v *DeleteSilencePolicyResponseBody) *DeleteSilencePolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteSilencePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
