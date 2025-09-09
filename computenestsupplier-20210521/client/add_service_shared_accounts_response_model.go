// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServiceSharedAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddServiceSharedAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddServiceSharedAccountsResponse
	GetStatusCode() *int32
	SetBody(v *AddServiceSharedAccountsResponseBody) *AddServiceSharedAccountsResponse
	GetBody() *AddServiceSharedAccountsResponseBody
}

type AddServiceSharedAccountsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddServiceSharedAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddServiceSharedAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddServiceSharedAccountsResponse) GoString() string {
	return s.String()
}

func (s *AddServiceSharedAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddServiceSharedAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddServiceSharedAccountsResponse) GetBody() *AddServiceSharedAccountsResponseBody {
	return s.Body
}

func (s *AddServiceSharedAccountsResponse) SetHeaders(v map[string]*string) *AddServiceSharedAccountsResponse {
	s.Headers = v
	return s
}

func (s *AddServiceSharedAccountsResponse) SetStatusCode(v int32) *AddServiceSharedAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddServiceSharedAccountsResponse) SetBody(v *AddServiceSharedAccountsResponseBody) *AddServiceSharedAccountsResponse {
	s.Body = v
	return s
}

func (s *AddServiceSharedAccountsResponse) Validate() error {
	return dara.Validate(s)
}
