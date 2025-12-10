// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveAccountResponse
	GetStatusCode() *int32
	SetBody(v *MoveAccountResponseBody) *MoveAccountResponse
	GetBody() *MoveAccountResponseBody
}

type MoveAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveAccountResponse) GoString() string {
	return s.String()
}

func (s *MoveAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveAccountResponse) GetBody() *MoveAccountResponseBody {
	return s.Body
}

func (s *MoveAccountResponse) SetHeaders(v map[string]*string) *MoveAccountResponse {
	s.Headers = v
	return s
}

func (s *MoveAccountResponse) SetStatusCode(v int32) *MoveAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveAccountResponse) SetBody(v *MoveAccountResponseBody) *MoveAccountResponse {
	s.Body = v
	return s
}

func (s *MoveAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
