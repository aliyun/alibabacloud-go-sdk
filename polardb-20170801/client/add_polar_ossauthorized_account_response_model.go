// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarOSSAuthorizedAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPolarOSSAuthorizedAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPolarOSSAuthorizedAccountResponse
	GetStatusCode() *int32
	SetBody(v *AddPolarOSSAuthorizedAccountResponseBody) *AddPolarOSSAuthorizedAccountResponse
	GetBody() *AddPolarOSSAuthorizedAccountResponseBody
}

type AddPolarOSSAuthorizedAccountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPolarOSSAuthorizedAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPolarOSSAuthorizedAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPolarOSSAuthorizedAccountResponse) GoString() string {
	return s.String()
}

func (s *AddPolarOSSAuthorizedAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPolarOSSAuthorizedAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPolarOSSAuthorizedAccountResponse) GetBody() *AddPolarOSSAuthorizedAccountResponseBody {
	return s.Body
}

func (s *AddPolarOSSAuthorizedAccountResponse) SetHeaders(v map[string]*string) *AddPolarOSSAuthorizedAccountResponse {
	s.Headers = v
	return s
}

func (s *AddPolarOSSAuthorizedAccountResponse) SetStatusCode(v int32) *AddPolarOSSAuthorizedAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountResponse) SetBody(v *AddPolarOSSAuthorizedAccountResponseBody) *AddPolarOSSAuthorizedAccountResponse {
	s.Body = v
	return s
}

func (s *AddPolarOSSAuthorizedAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
