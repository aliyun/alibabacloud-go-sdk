// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOuterAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOuterAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOuterAccountResponse
	GetStatusCode() *int32
	SetBody(v *AddOuterAccountResponseBody) *AddOuterAccountResponse
	GetBody() *AddOuterAccountResponseBody
}

type AddOuterAccountResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOuterAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOuterAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOuterAccountResponse) GoString() string {
	return s.String()
}

func (s *AddOuterAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOuterAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOuterAccountResponse) GetBody() *AddOuterAccountResponseBody {
	return s.Body
}

func (s *AddOuterAccountResponse) SetHeaders(v map[string]*string) *AddOuterAccountResponse {
	s.Headers = v
	return s
}

func (s *AddOuterAccountResponse) SetStatusCode(v int32) *AddOuterAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOuterAccountResponse) SetBody(v *AddOuterAccountResponseBody) *AddOuterAccountResponse {
	s.Body = v
	return s
}

func (s *AddOuterAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
