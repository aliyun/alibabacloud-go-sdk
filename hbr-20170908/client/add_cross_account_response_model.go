// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCrossAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCrossAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCrossAccountResponse
	GetStatusCode() *int32
	SetBody(v *AddCrossAccountResponseBody) *AddCrossAccountResponse
	GetBody() *AddCrossAccountResponseBody
}

type AddCrossAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCrossAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCrossAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCrossAccountResponse) GoString() string {
	return s.String()
}

func (s *AddCrossAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCrossAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCrossAccountResponse) GetBody() *AddCrossAccountResponseBody {
	return s.Body
}

func (s *AddCrossAccountResponse) SetHeaders(v map[string]*string) *AddCrossAccountResponse {
	s.Headers = v
	return s
}

func (s *AddCrossAccountResponse) SetStatusCode(v int32) *AddCrossAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCrossAccountResponse) SetBody(v *AddCrossAccountResponseBody) *AddCrossAccountResponse {
	s.Body = v
	return s
}

func (s *AddCrossAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
