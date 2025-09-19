// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackPathEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAttackPathEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAttackPathEventResponse
	GetStatusCode() *int32
	SetBody(v *ListAttackPathEventResponseBody) *ListAttackPathEventResponse
	GetBody() *ListAttackPathEventResponseBody
}

type ListAttackPathEventResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAttackPathEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAttackPathEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAttackPathEventResponse) GoString() string {
	return s.String()
}

func (s *ListAttackPathEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAttackPathEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAttackPathEventResponse) GetBody() *ListAttackPathEventResponseBody {
	return s.Body
}

func (s *ListAttackPathEventResponse) SetHeaders(v map[string]*string) *ListAttackPathEventResponse {
	s.Headers = v
	return s
}

func (s *ListAttackPathEventResponse) SetStatusCode(v int32) *ListAttackPathEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAttackPathEventResponse) SetBody(v *ListAttackPathEventResponseBody) *ListAttackPathEventResponse {
	s.Body = v
	return s
}

func (s *ListAttackPathEventResponse) Validate() error {
	return dara.Validate(s)
}
