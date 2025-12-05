// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostAccountResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostAccountResponseBody) *ModifyHostAccountResponse
	GetBody() *ModifyHostAccountResponseBody
}

type ModifyHostAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostAccountResponse) GetBody() *ModifyHostAccountResponseBody {
	return s.Body
}

func (s *ModifyHostAccountResponse) SetHeaders(v map[string]*string) *ModifyHostAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostAccountResponse) SetStatusCode(v int32) *ModifyHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostAccountResponse) SetBody(v *ModifyHostAccountResponseBody) *ModifyHostAccountResponse {
	s.Body = v
	return s
}

func (s *ModifyHostAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
