// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateAgAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AsyncCreateAgAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AsyncCreateAgAccountResponse
	GetStatusCode() *int32
	SetBody(v *AsyncCreateAgAccountResponseBody) *AsyncCreateAgAccountResponse
	GetBody() *AsyncCreateAgAccountResponseBody
}

type AsyncCreateAgAccountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AsyncCreateAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AsyncCreateAgAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateAgAccountResponse) GoString() string {
	return s.String()
}

func (s *AsyncCreateAgAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AsyncCreateAgAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AsyncCreateAgAccountResponse) GetBody() *AsyncCreateAgAccountResponseBody {
	return s.Body
}

func (s *AsyncCreateAgAccountResponse) SetHeaders(v map[string]*string) *AsyncCreateAgAccountResponse {
	s.Headers = v
	return s
}

func (s *AsyncCreateAgAccountResponse) SetStatusCode(v int32) *AsyncCreateAgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *AsyncCreateAgAccountResponse) SetBody(v *AsyncCreateAgAccountResponseBody) *AsyncCreateAgAccountResponse {
	s.Body = v
	return s
}

func (s *AsyncCreateAgAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
