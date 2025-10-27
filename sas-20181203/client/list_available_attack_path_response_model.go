// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableAttackPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableAttackPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableAttackPathResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableAttackPathResponseBody) *ListAvailableAttackPathResponse
	GetBody() *ListAvailableAttackPathResponseBody
}

type ListAvailableAttackPathResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableAttackPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableAttackPathResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableAttackPathResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableAttackPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableAttackPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableAttackPathResponse) GetBody() *ListAvailableAttackPathResponseBody {
	return s.Body
}

func (s *ListAvailableAttackPathResponse) SetHeaders(v map[string]*string) *ListAvailableAttackPathResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableAttackPathResponse) SetStatusCode(v int32) *ListAvailableAttackPathResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableAttackPathResponse) SetBody(v *ListAvailableAttackPathResponseBody) *ListAvailableAttackPathResponse {
	s.Body = v
	return s
}

func (s *ListAvailableAttackPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
