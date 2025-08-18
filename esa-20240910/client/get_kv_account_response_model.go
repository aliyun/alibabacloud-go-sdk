// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKvAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKvAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetKvAccountResponseBody) *GetKvAccountResponse
	GetBody() *GetKvAccountResponseBody
}

type GetKvAccountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKvAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKvAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKvAccountResponse) GoString() string {
	return s.String()
}

func (s *GetKvAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKvAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKvAccountResponse) GetBody() *GetKvAccountResponseBody {
	return s.Body
}

func (s *GetKvAccountResponse) SetHeaders(v map[string]*string) *GetKvAccountResponse {
	s.Headers = v
	return s
}

func (s *GetKvAccountResponse) SetStatusCode(v int32) *GetKvAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKvAccountResponse) SetBody(v *GetKvAccountResponseBody) *GetKvAccountResponse {
	s.Body = v
	return s
}

func (s *GetKvAccountResponse) Validate() error {
	return dara.Validate(s)
}
