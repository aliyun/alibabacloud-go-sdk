// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackEventInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAttackEventInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAttackEventInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListAttackEventInfoResponseBody) *ListAttackEventInfoResponse
	GetBody() *ListAttackEventInfoResponseBody
}

type ListAttackEventInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAttackEventInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAttackEventInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAttackEventInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAttackEventInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAttackEventInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAttackEventInfoResponse) GetBody() *ListAttackEventInfoResponseBody {
	return s.Body
}

func (s *ListAttackEventInfoResponse) SetHeaders(v map[string]*string) *ListAttackEventInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAttackEventInfoResponse) SetStatusCode(v int32) *ListAttackEventInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAttackEventInfoResponse) SetBody(v *ListAttackEventInfoResponseBody) *ListAttackEventInfoResponse {
	s.Body = v
	return s
}

func (s *ListAttackEventInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
