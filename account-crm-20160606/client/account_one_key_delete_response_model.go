// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountOneKeyDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccountOneKeyDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccountOneKeyDeleteResponse
	GetStatusCode() *int32
	SetBody(v *AccountOneKeyDeleteResponseBody) *AccountOneKeyDeleteResponse
	GetBody() *AccountOneKeyDeleteResponseBody
}

type AccountOneKeyDeleteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountOneKeyDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountOneKeyDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountOneKeyDeleteResponse) GoString() string {
	return s.String()
}

func (s *AccountOneKeyDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccountOneKeyDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccountOneKeyDeleteResponse) GetBody() *AccountOneKeyDeleteResponseBody {
	return s.Body
}

func (s *AccountOneKeyDeleteResponse) SetHeaders(v map[string]*string) *AccountOneKeyDeleteResponse {
	s.Headers = v
	return s
}

func (s *AccountOneKeyDeleteResponse) SetStatusCode(v int32) *AccountOneKeyDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountOneKeyDeleteResponse) SetBody(v *AccountOneKeyDeleteResponseBody) *AccountOneKeyDeleteResponse {
	s.Body = v
	return s
}

func (s *AccountOneKeyDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
