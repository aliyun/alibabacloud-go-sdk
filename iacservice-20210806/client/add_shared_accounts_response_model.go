// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSharedAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSharedAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSharedAccountsResponse
	GetStatusCode() *int32
	SetBody(v *AddSharedAccountsResponseBody) *AddSharedAccountsResponse
	GetBody() *AddSharedAccountsResponseBody
}

type AddSharedAccountsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSharedAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *AddSharedAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSharedAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSharedAccountsResponse) GetBody() *AddSharedAccountsResponseBody {
	return s.Body
}

func (s *AddSharedAccountsResponse) SetHeaders(v map[string]*string) *AddSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *AddSharedAccountsResponse) SetStatusCode(v int32) *AddSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSharedAccountsResponse) SetBody(v *AddSharedAccountsResponseBody) *AddSharedAccountsResponse {
	s.Body = v
	return s
}

func (s *AddSharedAccountsResponse) Validate() error {
	return dara.Validate(s)
}
