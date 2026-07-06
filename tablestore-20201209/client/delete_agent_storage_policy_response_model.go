// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentStoragePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentStoragePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentStoragePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentStoragePolicyResponseBody) *DeleteAgentStoragePolicyResponse
	GetBody() *DeleteAgentStoragePolicyResponseBody
}

type DeleteAgentStoragePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentStoragePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentStoragePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentStoragePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentStoragePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentStoragePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentStoragePolicyResponse) GetBody() *DeleteAgentStoragePolicyResponseBody {
	return s.Body
}

func (s *DeleteAgentStoragePolicyResponse) SetHeaders(v map[string]*string) *DeleteAgentStoragePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentStoragePolicyResponse) SetStatusCode(v int32) *DeleteAgentStoragePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentStoragePolicyResponse) SetBody(v *DeleteAgentStoragePolicyResponseBody) *DeleteAgentStoragePolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentStoragePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
