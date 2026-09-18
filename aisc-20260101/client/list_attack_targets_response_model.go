// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttackTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAttackTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAttackTargetsResponse
	GetStatusCode() *int32
	SetBody(v *ListAttackTargetsResponseBody) *ListAttackTargetsResponse
	GetBody() *ListAttackTargetsResponseBody
}

type ListAttackTargetsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAttackTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAttackTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAttackTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListAttackTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAttackTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAttackTargetsResponse) GetBody() *ListAttackTargetsResponseBody {
	return s.Body
}

func (s *ListAttackTargetsResponse) SetHeaders(v map[string]*string) *ListAttackTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListAttackTargetsResponse) SetStatusCode(v int32) *ListAttackTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAttackTargetsResponse) SetBody(v *ListAttackTargetsResponseBody) *ListAttackTargetsResponse {
	s.Body = v
	return s
}

func (s *ListAttackTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
