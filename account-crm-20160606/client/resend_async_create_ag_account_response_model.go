// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendAsyncCreateAgAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResendAsyncCreateAgAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResendAsyncCreateAgAccountResponse
	GetStatusCode() *int32
	SetBody(v *ResendAsyncCreateAgAccountResponseBody) *ResendAsyncCreateAgAccountResponse
	GetBody() *ResendAsyncCreateAgAccountResponseBody
}

type ResendAsyncCreateAgAccountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendAsyncCreateAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendAsyncCreateAgAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ResendAsyncCreateAgAccountResponse) GoString() string {
	return s.String()
}

func (s *ResendAsyncCreateAgAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResendAsyncCreateAgAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResendAsyncCreateAgAccountResponse) GetBody() *ResendAsyncCreateAgAccountResponseBody {
	return s.Body
}

func (s *ResendAsyncCreateAgAccountResponse) SetHeaders(v map[string]*string) *ResendAsyncCreateAgAccountResponse {
	s.Headers = v
	return s
}

func (s *ResendAsyncCreateAgAccountResponse) SetStatusCode(v int32) *ResendAsyncCreateAgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendAsyncCreateAgAccountResponse) SetBody(v *ResendAsyncCreateAgAccountResponseBody) *ResendAsyncCreateAgAccountResponse {
	s.Body = v
	return s
}

func (s *ResendAsyncCreateAgAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
