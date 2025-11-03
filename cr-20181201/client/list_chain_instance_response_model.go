// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChainInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChainInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChainInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListChainInstanceResponseBody) *ListChainInstanceResponse
	GetBody() *ListChainInstanceResponseBody
}

type ListChainInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChainInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChainInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChainInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChainInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChainInstanceResponse) GetBody() *ListChainInstanceResponseBody {
	return s.Body
}

func (s *ListChainInstanceResponse) SetHeaders(v map[string]*string) *ListChainInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListChainInstanceResponse) SetStatusCode(v int32) *ListChainInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChainInstanceResponse) SetBody(v *ListChainInstanceResponseBody) *ListChainInstanceResponse {
	s.Body = v
	return s
}

func (s *ListChainInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
