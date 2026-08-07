// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarOSSAuthorizedAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarOSSAuthorizedAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarOSSAuthorizedAccountResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarOSSAuthorizedAccountResponseBody) *DeletePolarOSSAuthorizedAccountResponse
	GetBody() *DeletePolarOSSAuthorizedAccountResponseBody
}

type DeletePolarOSSAuthorizedAccountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarOSSAuthorizedAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarOSSAuthorizedAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarOSSAuthorizedAccountResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarOSSAuthorizedAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarOSSAuthorizedAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarOSSAuthorizedAccountResponse) GetBody() *DeletePolarOSSAuthorizedAccountResponseBody {
	return s.Body
}

func (s *DeletePolarOSSAuthorizedAccountResponse) SetHeaders(v map[string]*string) *DeletePolarOSSAuthorizedAccountResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountResponse) SetStatusCode(v int32) *DeletePolarOSSAuthorizedAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountResponse) SetBody(v *DeletePolarOSSAuthorizedAccountResponseBody) *DeletePolarOSSAuthorizedAccountResponse {
	s.Body = v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
