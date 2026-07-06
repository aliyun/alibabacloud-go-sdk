// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByAgentStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcInfoByAgentStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcInfoByAgentStorageResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcInfoByAgentStorageResponseBody) *ListVpcInfoByAgentStorageResponse
	GetBody() *ListVpcInfoByAgentStorageResponseBody
}

type ListVpcInfoByAgentStorageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcInfoByAgentStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcInfoByAgentStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByAgentStorageResponse) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByAgentStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcInfoByAgentStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcInfoByAgentStorageResponse) GetBody() *ListVpcInfoByAgentStorageResponseBody {
	return s.Body
}

func (s *ListVpcInfoByAgentStorageResponse) SetHeaders(v map[string]*string) *ListVpcInfoByAgentStorageResponse {
	s.Headers = v
	return s
}

func (s *ListVpcInfoByAgentStorageResponse) SetStatusCode(v int32) *ListVpcInfoByAgentStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcInfoByAgentStorageResponse) SetBody(v *ListVpcInfoByAgentStorageResponseBody) *ListVpcInfoByAgentStorageResponse {
	s.Body = v
	return s
}

func (s *ListVpcInfoByAgentStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
