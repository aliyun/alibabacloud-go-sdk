// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBindAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBindAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBindAccountResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBindAccountResponseBody) *ModifyBindAccountResponse
	GetBody() *ModifyBindAccountResponseBody
}

type ModifyBindAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBindAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBindAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBindAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBindAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBindAccountResponse) GetBody() *ModifyBindAccountResponseBody {
	return s.Body
}

func (s *ModifyBindAccountResponse) SetHeaders(v map[string]*string) *ModifyBindAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyBindAccountResponse) SetStatusCode(v int32) *ModifyBindAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBindAccountResponse) SetBody(v *ModifyBindAccountResponseBody) *ModifyBindAccountResponse {
	s.Body = v
	return s
}

func (s *ModifyBindAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
