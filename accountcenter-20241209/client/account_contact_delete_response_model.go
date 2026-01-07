// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccountContactDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccountContactDeleteResponse
	GetStatusCode() *int32
	SetBody(v *AccountContactDeleteResponseBody) *AccountContactDeleteResponse
	GetBody() *AccountContactDeleteResponseBody
}

type AccountContactDeleteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountContactDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountContactDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s AccountContactDeleteResponse) GoString() string {
	return s.String()
}

func (s *AccountContactDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccountContactDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccountContactDeleteResponse) GetBody() *AccountContactDeleteResponseBody {
	return s.Body
}

func (s *AccountContactDeleteResponse) SetHeaders(v map[string]*string) *AccountContactDeleteResponse {
	s.Headers = v
	return s
}

func (s *AccountContactDeleteResponse) SetStatusCode(v int32) *AccountContactDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountContactDeleteResponse) SetBody(v *AccountContactDeleteResponseBody) *AccountContactDeleteResponse {
	s.Body = v
	return s
}

func (s *AccountContactDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
