// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncCreateAgAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAsyncCreateAgAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAsyncCreateAgAccountResponse
	GetStatusCode() *int32
	SetBody(v *CancelAsyncCreateAgAccountResponseBody) *CancelAsyncCreateAgAccountResponse
	GetBody() *CancelAsyncCreateAgAccountResponseBody
}

type CancelAsyncCreateAgAccountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAsyncCreateAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAsyncCreateAgAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncCreateAgAccountResponse) GoString() string {
	return s.String()
}

func (s *CancelAsyncCreateAgAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAsyncCreateAgAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAsyncCreateAgAccountResponse) GetBody() *CancelAsyncCreateAgAccountResponseBody {
	return s.Body
}

func (s *CancelAsyncCreateAgAccountResponse) SetHeaders(v map[string]*string) *CancelAsyncCreateAgAccountResponse {
	s.Headers = v
	return s
}

func (s *CancelAsyncCreateAgAccountResponse) SetStatusCode(v int32) *CancelAsyncCreateAgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAsyncCreateAgAccountResponse) SetBody(v *CancelAsyncCreateAgAccountResponseBody) *CancelAsyncCreateAgAccountResponse {
	s.Body = v
	return s
}

func (s *CancelAsyncCreateAgAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
