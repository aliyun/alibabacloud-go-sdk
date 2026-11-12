// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomImageShareAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomImageShareAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomImageShareAccountResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomImageShareAccountResponseBody) *AddCustomImageShareAccountResponse
	GetBody() *AddCustomImageShareAccountResponseBody
}

type AddCustomImageShareAccountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomImageShareAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomImageShareAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomImageShareAccountResponse) GoString() string {
	return s.String()
}

func (s *AddCustomImageShareAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomImageShareAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomImageShareAccountResponse) GetBody() *AddCustomImageShareAccountResponseBody {
	return s.Body
}

func (s *AddCustomImageShareAccountResponse) SetHeaders(v map[string]*string) *AddCustomImageShareAccountResponse {
	s.Headers = v
	return s
}

func (s *AddCustomImageShareAccountResponse) SetStatusCode(v int32) *AddCustomImageShareAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomImageShareAccountResponse) SetBody(v *AddCustomImageShareAccountResponseBody) *AddCustomImageShareAccountResponse {
	s.Body = v
	return s
}

func (s *AddCustomImageShareAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
